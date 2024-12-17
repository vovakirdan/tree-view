package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sort"
)

const indentSpace = "    "  // spaces for indent

// printTree flush dir and files tree recursively
func printTree(path string, indent string, isLast bool, printRoot bool) {
	// skip hidden paths
	if strings.HasPrefix(filepath.Base(path), ".") {
		return
	}

	// define symbols for branch
	var branch, nextIndent string
	if printRoot { // for root
		branch = ""
		nextIndent = ""
	} else if isLast {
		branch = "└── "
		nextIndent = indent + "    "
	} else {
		branch = "├── "
		nextIndent = indent + "│   "
	}

	// mark current path
	name := filepath.Base(path)
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		name += "/" // add "/" for paths
	}

	// print current path
	fmt.Println(indent + branch + name)

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

	// filter hidden
	var filteredEntries []os.FileInfo
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	for i, entry := range filteredEntries {
		isLastEntry := i == len(filteredEntries)-1
		childPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			printTree(childPath, nextIndent, isLastEntry, false)
		} else {
			branchSymbol := "├── "
			if isLastEntry {
				branchSymbol = "└── "
			}
			fmt.Println(nextIndent + branchSymbol + entry.Name())
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

	printTree(root, "", true, true)
}