package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

type GeneratedFile struct {
	Path    string
	Content string
	Mode    os.FileMode
}

func writeFiles(outputDir string, force bool, files []GeneratedFile) error {
	for _, file := range files {
		target := filepath.Join(outputDir, file.Path)
		info, err := os.Stat(target)
		if err == nil {
			if info.IsDir() {
				return fmt.Errorf("cannot write %s: path is a directory", target)
			}
			if !force {
				return fmt.Errorf("refusing to overwrite existing file %s; use --force to overwrite generated files", target)
			}
			continue
		}
		if !os.IsNotExist(err) {
			return fmt.Errorf("cannot inspect %s: %w", target, err)
		}
	}

	for _, file := range files {
		target := filepath.Join(outputDir, file.Path)
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return fmt.Errorf("create directory for %s: %w", target, err)
		}
		if err := os.WriteFile(target, []byte(file.Content), file.Mode); err != nil {
			return fmt.Errorf("write %s: %w", target, err)
		}
	}

	return nil
}
