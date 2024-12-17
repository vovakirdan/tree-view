package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const indentSpace = "    " // Number of spaces for each indentation

// printTree recursively displays the directory tree structure.
func printTree(path string, indent string, isLast bool, printRoot bool, showHidden bool, excludeList []string) {
	// Extract the file or folder name
	baseName := filepath.Base(path)

	// Skip hidden files and directories if showHidden is false
	if !showHidden && strings.HasPrefix(baseName, ".") {
		return
	}

	// Skip files and directories from the exclude list
	for _, exclude := range excludeList {
		if baseName == exclude {
			return
		}
	}

	// Determine the branch and indentation symbols
	var branch, nextIndent string
	if printRoot { // Root directory formatting
		branch = ""
		nextIndent = ""
	} else if isLast {
		branch = "└── "
		nextIndent = indent + "    "
	} else {
		branch = "├── "
		nextIndent = indent + "│   "
	}

	// Add "/" at the end for directories
	name := baseName
	if info, err := os.Stat(path); err == nil && info.IsDir() {
		name += "/"
	}

	// Print the current file or folder
	fmt.Println(indent + branch + name)

	// If not a directory, return
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		return
	}

	// Read the directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// no need to sort actually
	// sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	// Filter out hidden entries and excluded names
	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		entryName := entry.Name()
		if (!showHidden && strings.HasPrefix(entryName, ".")) || contains(excludeList, entryName) {
			continue
		}
		filteredEntries = append(filteredEntries, entry)
	}

	// Iterate through the directory contents
	for i, entry := range filteredEntries {
		isLastEntry := i == len(filteredEntries)-1
		childPath := filepath.Join(path, entry.Name())
		printTree(childPath, nextIndent, isLastEntry, false, showHidden, excludeList)
	}
}

// contains checks if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	// Command-line flags
	showHidden := flag.Bool("a", false, "Show hidden files and directories")
	exclude := flag.String("exclude", "", "Comma-separated list of files/directories to exclude")
	flag.Parse()

	// Get the directory path from the arguments
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: tree [options] <directory>")
		fmt.Println("Options:")
		flag.PrintDefaults()
		return
	}
	root := args[0]

	// Check if the provided path is valid
	info, err := os.Stat(root)
	if os.IsNotExist(err) {
		fmt.Println("Error: directory does not exist")
		return
	}
	if !info.IsDir() {
		fmt.Println("Error: provided path is not a directory")
		return
	}

	// Parse the exclude list
	excludeList := []string{}
	if *exclude != "" {
		excludeList = strings.Split(*exclude, ",")
	}

	// Print the directory tree
	printTree(root, "", true, true, *showHidden, excludeList)
}
