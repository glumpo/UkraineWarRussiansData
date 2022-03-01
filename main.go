package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MichaelTJones/walk"
	"github.com/tealeg/xlsx"
)

func search(file string, patterns []string) error {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		return err
	}

	for _, sheet := range xlFile.Sheets {
		for rowN, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				if equal(text, patterns) {
					fmt.Printf("In %s:%s:%d: %s\n", file, sheet.Name, rowN, dumpRow(row))
				}
			}
		}
	}

	return nil
}

func equal(s string, patterns []string) bool {
	for _, p := range patterns {
		if !strings.Contains(s, p) {
			return false
		}
	}
	return true
}

func dumpRow(row *xlsx.Row) string {
	var b strings.Builder
	for _, cell := range row.Cells {
		b.WriteString(cell.String())
		b.WriteString(" ")
	}
	return b.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Usage: searcher pattern1 pattern2 ...")
	}

	patterns := os.Args[1:]
	err := walk.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".xlsx" {
			return nil
		}
		return search(path, patterns)
	})

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
