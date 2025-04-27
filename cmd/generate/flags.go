package main

import (
	"fmt"
	"os"
)

const TablerFlags = "https://github.com/tabler/tabler-flags"

func generateFlags() (map[string]int, error) {
	if err := os.MkdirAll("./flags", 0755); err != nil {
		return nil, fmt.Errorf("error creating flags directory: %w", err)
	}

	f, err := downloadRepo(TablerFlags)
	if err != nil {
		return nil, fmt.Errorf("error downloading flags repository: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("error getting file info: %w", err)
	}

	stats, err := processZipFile(f, info.Size())
	if err != nil {
		return nil, fmt.Errorf("error processing zip file: %w", err)
	}

	return stats, nil
}
