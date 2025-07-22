package internal

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func readSQLFiles(dir string) ([]Migration, error) {
	var files []Migration

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}

		path := filepath.Join(dir, e.Name())
		hash, err := computeHash(path)
		if err != nil {
			return nil, err
		}

		files = append(files, Migration{
			Name: e.Name(),
			Path: path,
			Hash: hash,
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return files, nil
}
