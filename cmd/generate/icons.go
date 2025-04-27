package main

import (
	"fmt"
	"os"
)

const TablerIcons = "https://github.com/tabler/tabler-icons"

func generateIcons() (map[string]int, error) {
	if err := os.MkdirAll("./tabler", 0775); err != nil {
		return nil, fmt.Errorf("error while creating tabler icons folder: %w", err)
	}

	f, err := downloadRepo(TablerIcons)
	if err != nil {
		return nil, fmt.Errorf("error while downloading tabler icons: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("error while getting file info: %w", err)
	}

	stats, err := processZipFile(f, info.Size())
	if err != nil {
		return nil, fmt.Errorf("error while processing zip file: %w", err)
	}

	return stats, nil
}
