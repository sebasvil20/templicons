package main

import "path/filepath"

type Icon struct {
	baseName string
	svgType  string
}

func (i Icon) packageName() string {
	if i.svgType == "flags" {
		return "flags"
	}

	return "tabler"
}

func (i Icon) componentName() string {
	name := toPascalCase(i.baseName)

	if i.svgType == "filled" {
		name += "Filled"
	}

	return name
}

func (i Icon) iconComponent() string {
	switch i.svgType {
	case "filled":
		return "FilledIcon"
	case "outlined":
		return "OutlinedIcon"
	case "flags":
		return "FlagIcon"
	}
	return ""
}

func (i Icon) viewbox() string {
	if i.svgType == "flags" {
		return "0 0 30 24"
	}

	return "0 0 24 24"
}

func (i Icon) filename() string {
	dir := "tabler"
	if i.svgType == "flags" {
		dir = "flags"
	}

	var filled string
	if i.svgType == "filled" {
		filled = "Filled"
	}

	return filepath.Join(dir, toCamelCase(i.baseName)+filled+".templ")
}
