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
	// Crear directorio de salida si no existe
	err := os.MkdirAll("./tabler", 0755)
	if err != nil {
		fmt.Printf("Error al crear directorio de salida: %v\n", err)
		return
	}

	// Procesar carpeta de outline
	outlineDir := "./tmp/outline"
	if _, err := os.Stat(outlineDir); err == nil {
		processDirectory(outlineDir, false)
	} else {
		fmt.Printf("Directorio %s no encontrado o no accesible\n", outlineDir)
	}

	// Procesar carpeta de filled
	filledDir := "./tmp/filled"
	if _, err := os.Stat(filledDir); err == nil {
		processDirectory(filledDir, true)
	} else {
		fmt.Printf("Directorio %s no encontrado o no accesible\n", filledDir)
	}

	fmt.Println("Proceso completado con éxito.")
}

func processDirectory(dirPath string, isFilled bool) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Error al leer directorio %s: %v\n", dirPath, err)
		return
	}

	// Procesar cada archivo SVG
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".svg" {
			processFile(dirPath, file.Name(), isFilled)
		}
	}
}

func processFile(dirPath string, filename string, isFilled bool) {
	// Leer el archivo SVG
	inputPath := filepath.Join(dirPath, filename)
	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error al leer archivo %s: %v\n", filename, err)
		return
	}

	// Convertir nombres
	baseName := strings.TrimSuffix(filename, ".svg")
	camelCaseName := toCamelCase(baseName)
	pascalCaseName := toPascalCase(baseName)

	// Añadir sufijo "Filled" si es necesario
	if isFilled {
		camelCaseName += "Filled"
		pascalCaseName += "Filled"
	}

	// Extraer contenido del SVG y procesar las etiquetas
	svgContent := extractAndProcessSVGContent(string(content))

	// Generar archivo TEMPL
	outputPath := filepath.Join("./tabler", camelCaseName+".templ")

	// Crear contenido del archivo TEMPL
	templContent := fmt.Sprintf(`package tabler
import "github.com/sebasvil20/templicons/i"

templ %s(props ...Props) {
    @i.Icon("0 0 24 24", props...) {
        %s
    }
}
`, pascalCaseName, svgContent)

	// Escribir archivo TEMPL
	err = ioutil.WriteFile(outputPath, []byte(templContent), 0644)
	if err != nil {
		fmt.Printf("Error al escribir archivo %s: %v\n", outputPath, err)
		return
	}

	fmt.Printf("Archivo generado: %s\n", outputPath)
}

func extractAndProcessSVGContent(svgStr string) string {
	// Extraer contenido entre las etiquetas <svg> y </svg>
	svgRegex := regexp.MustCompile(`(?s)<svg[^>]*>(.*?)</svg>`)
	matches := svgRegex.FindStringSubmatch(svgStr)

	if len(matches) < 2 {
		return ""
	}

	innerContent := strings.TrimSpace(matches[1])

	// Paso 1: Procesar grupos primero (<g> y su contenido)
	// Preparamos una expresión más simple para extraer los grupos
	groupRegex := regexp.MustCompile(`(?s)<g([^>]*?)>(.*?)</g>`)

	// Reemplazamos cada grupo con su versión procesada
	innerContent = groupRegex.ReplaceAllStringFunc(innerContent, func(groupMatch string) string {
		parts := groupRegex.FindStringSubmatch(groupMatch)
		if len(parts) < 3 {
			return groupMatch
		}

		groupAttrs := parts[1]
		groupContent := parts[2]

		// Procesar contenido dentro del grupo
		processedGroupContent := processTagsInContent(groupContent)

		return fmt.Sprintf("<g%s>%s</g>", groupAttrs, processedGroupContent)
	})

	// Paso 2: Procesar las etiquetas fuera de los grupos
	innerContent = processTagsInContent(innerContent)

	return innerContent
}

func processTagsInContent(content string) string {
	// Este enfoque es más simple: usamos expresiones regulares específicas
	// para cada tipo de transformación

	// 1. Convertir etiquetas autocerrables (terminan con />)
	// Patrón: busca cualquier etiqueta que termine con />
	selfClosingRegex := regexp.MustCompile(`<([a-zA-Z][a-zA-Z0-9]*)((?:[^>]*?))/>`)
	content = selfClosingRegex.ReplaceAllString(content, "<$1$2></$1>")

	// 2. Hacemos un post-procesamiento para eliminar etiquetas de cierre duplicadas
	// Como Go no soporta referencias de captura en los reemplazos (\1),
	// tenemos que hacerlo de forma específica para cada etiqueta SVG común

	// Lista de etiquetas SVG comunes
	svgTags := []string{"path", "circle", "rect", "line", "polygon", "polyline", "ellipse", "text", "tspan"}

	// Para cada tipo de etiqueta, buscamos cierres duplicados
	for _, tag := range svgTags {
		doubleClosePattern := fmt.Sprintf("</%s></%s>", tag, tag)
		content = strings.ReplaceAll(content, doubleClosePattern, fmt.Sprintf("</%s>", tag))
	}

	return content
}

func toCamelCase(s string) string {
	// Dividir por guiones
	parts := strings.Split(s, "-")

	// Primera palabra en minúscula
	result := strings.ToLower(parts[0])

	// Resto de palabras con la primera letra en mayúscula
	for i := 1; i < len(parts); i++ {
		if parts[i] != "" {
			word := parts[i]
			// Verificar si el primer carácter es letra (para no convertir números iniciales)
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
	// Obtener camelCase y convertir primera letra a mayúscula
	camel := toCamelCase(s)
	if camel == "" {
		return ""
	}

	// Convertir primera letra a mayúscula
	return strings.ToUpper(camel[:1]) + camel[1:]
}
