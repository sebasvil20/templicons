package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	err := os.MkdirAll("./tabler", 0755)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	outlineDir := "./tmp/outline"
	if _, err := os.Stat(outlineDir); err == nil {
		processDirectory(outlineDir, false)
	} else {
		fmt.Printf("Directory %s not found or not accessible\n", outlineDir)
	}

	filledDir := "./tmp/filled"
	if _, err := os.Stat(filledDir); err == nil {
		processDirectory(filledDir, true)
	} else {
		fmt.Printf("Directory %s not found or not accessible\n", filledDir)
	}

	fmt.Println("Process completed successfully.")
}

func processDirectory(dirPath string, isFilled bool) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dirPath, err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".svg" {
			processFile(dirPath, file.Name(), isFilled)
		}
	}
}

func processFile(dirPath string, filename string, isFilled bool) {
	inputPath := filepath.Join(dirPath, filename)
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	baseName := strings.TrimSuffix(filename, ".svg")
	camelCaseName := toCamelCase(baseName)
	pascalCaseName := toPascalCase(baseName)

	if isFilled {
		camelCaseName += "Filled"
		pascalCaseName += "Filled"
	}

	svgContent := extractAndProcessSVGContent(string(content))

	outputPath := filepath.Join("./tabler", camelCaseName+".templ")

	// Usar el componente adecuado según el tipo de icono
	var iconComponent string
	if isFilled {
		iconComponent = "FilledIcon"
	} else {
		iconComponent = "OutlinedIcon"
	}

	templContent := fmt.Sprintf(`package tabler
import "github.com/sebasvil20/templicons/i"

templ %s(props ...i.Props) {
    @i.%s("0 0 24 24", props...) {
        %s
    }
}
`, pascalCaseName, iconComponent, svgContent)

	err = ioutil.WriteFile(outputPath, []byte(templContent), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputPath, err)
		return
	}

	fmt.Printf("File generated: %s\n", outputPath)
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
