package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	iconStats, err := generateIcons()
	if err != nil {
		fmt.Printf("Error generating icons: %v\n", err)
		return
	}

	flagStats, err := generateFlags()
	if err != nil {
		fmt.Printf("Error generating flags: %v\n", err)
		return
	}

	fmt.Printf("Icon stats: %d Outlined, %d Filled\n", iconStats["outlined"], iconStats["filled"])
	fmt.Printf("Flag stats: %d flags\n", flagStats["flags"])

	fmt.Println("Process completed successfully.")
}

func downloadRepo(url string) (*os.File, error) {
	resp, err := http.Get(url + "/archive/refs/heads/main.zip")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f, err := os.CreateTemp("", "tabler-icons")
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		return nil, err
	}

	// Make sure the file is synced to disk since we need to read the filesize
	if err := f.Sync(); err != nil {
		return nil, err
	}

	return f, nil
}

func processZipFile(f io.ReaderAt, filesize int64) (map[string]int, error) {
	r, err := zip.NewReader(f, filesize)
	if err != nil {
		return nil, err
	}

	stats := map[string]int{
		"filled":   0,
		"outlined": 0,
		"flags":    0,
	}

	for _, file := range r.File {
		if file.Mode().IsDir() || filepath.Ext(file.Name) != ".svg" {
			continue
		}

		dir := filepath.Dir(file.Name)
		switch {
		case strings.HasSuffix(dir, "icons/filled"):
			processFile(file, "filled")
			stats["filled"]++
		case strings.HasSuffix(dir, "icons/outline"):
			processFile(file, "outlined")
			stats["outlined"]++
		case strings.HasSuffix(dir, "flags"):
			processFile(file, "flags")
			stats["flags"]++
		default:
			continue
		}
	}

	return stats, nil
}

func processFile(zf *zip.File, svgType string) {
	f, err := zf.Open()
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", zf.Name, err)
		return
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", zf.Name, err)
		return
	}

	filename := filepath.Base(zf.Name)
	fmt.Printf("File %s %s\n", zf.Name, svgType)

	baseName := strings.TrimSuffix(filename, ".svg")

	icon := Icon{
		baseName: baseName,
		svgType:  svgType,
	}

	svgContent := extractAndProcessSVGContent(string(content))

	templContent := fmt.Sprintf(`package %s
import "github.com/sebasvil20/templicons/i"

templ %s(props ...i.Props) {
    @i.%s("%s", props...) {
        %s
    }
}
`, icon.packageName(), icon.componentName(), icon.iconComponent(), icon.viewbox(), svgContent)

	err = os.WriteFile(icon.filename(), []byte(templContent), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", icon.filename(), err)
		return
	}

	fmt.Printf("File generated: %s\n", icon.filename())
}

func extractAndProcessSVGContent(svgStr string) string {
	svgRegex := regexp.MustCompile(`(?s)<svg[^>]*>(.*?)</svg>`)
	matches := svgRegex.FindStringSubmatch(svgStr)

	if len(matches) < 2 {
		return ""
	}

	innerContent := strings.TrimSpace(matches[1])

	groupRegex := regexp.MustCompile(`(?s)<g([^>]*?)>(.*?)</g>`)

	innerContent = groupRegex.ReplaceAllStringFunc(innerContent, func(groupMatch string) string {
		parts := groupRegex.FindStringSubmatch(groupMatch)
		if len(parts) < 3 {
			return groupMatch
		}

		groupAttrs := parts[1]
		groupContent := parts[2]

		processedGroupContent := processTagsInContent(groupContent)

		return fmt.Sprintf("<g%s>%s</g>", groupAttrs, processedGroupContent)
	})

	innerContent = processTagsInContent(innerContent)

	return innerContent
}

func processTagsInContent(content string) string {
	selfClosingRegex := regexp.MustCompile(`<([a-zA-Z][a-zA-Z0-9]*)((?:[^>]*?))/>`)
	content = selfClosingRegex.ReplaceAllString(content, "<$1$2></$1>")

	svgTags := []string{"path", "circle", "rect", "line", "polygon", "polyline", "ellipse", "text", "tspan"}

	for _, tag := range svgTags {
		doubleClosePattern := fmt.Sprintf("</%s></%s>", tag, tag)
		content = strings.ReplaceAll(content, doubleClosePattern, fmt.Sprintf("</%s>", tag))
	}

	return content
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "-")

	result := strings.ToLower(parts[0])

	for i := 1; i < len(parts); i++ {
		if parts[i] != "" {
			word := parts[i]
			if len(word) > 0 && unicode.IsLetter(rune(word[0])) {
				result += string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
			} else {
				result += word
			}
		}
	}

	return result
}

func toPascalCase(s string) string {
	camel := toCamelCase(s)
	if camel == "" {
		return ""
	}

	return strings.ToUpper(camel[:1]) + camel[1:]
}
