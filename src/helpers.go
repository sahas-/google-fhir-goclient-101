package main

import (
	"io/fs"
	"sort"
	"strings"
)

func reorderFiles(fileInfos []fs.FileInfo, order []string) []fs.FileInfo {
	sort.Slice(fileInfos, func(i, j int) bool {
		indexI := findSubstringIndex(fileInfos[i].Name(), order)
		indexJ := findSubstringIndex(fileInfos[j].Name(), order)
		return indexI < indexJ
	})

	return fileInfos
}

func findSubstringIndex(s string, substrings []string) int {
	for i, substr := range substrings {
		if strings.Contains(s, substr) {
			return i
		}
	}
	return len(substrings)
}
