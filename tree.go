package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

const indentSpace = "    "  // spaces for indent

// printTree flush dir and files tree recursively
func printTree(path string, indent string, isLast bool) {
	// define symbols for branch
	var branch, nextIndent string
	if isLast {
		branch = "└── "
		nextIndent = indent + indentSpace
	} else {
		branch = "├── "
		nextIndent = indent + "│   "
	}

	// print current path
	fmt.Println(indent + branch + filepath.Base(path))

	// check if dir
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening:", err)
		return
	}
	defer file.Close()

	entries, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// sort by name
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	// recursively print dirs and files
	for i, entry := range entries {
		isLastEntry := i == len(entries) - 1
		childPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			printTree(childPath, nextIndent, isLastEntry)
		} else {
			fmt.Println(nextIndent + "├── " + entry.Name())
		}
	}
}

func main() {
	// get args
	if len(os.Args) < 2 {
		fmt.Println("Usage: tree <path>")
		return
	}
	root := os.Args[1]

	// check if dir exists
	info, err := os.Stat(root)
	if os.IsNotExist(err) {
		fmt.Println("Error: directory not found")
		return
	}
	if !info.IsDir() {
		fmt.Println("Error: not a directory")
		return
	}

	printTree(root, "", true)
}