# Tree Utility

A simple command-line utility written in **Go** to display the directory structure in a tree-like format.

## Features

- Display files and directories in a structured, easy-to-read tree format.
- Optionally display hidden files and directories using the `-a` flag.
- Exclude specific files or directories by name using the `-exclude` flag.

---

## Installation

1. Install [Go](https://golang.org/) (if not already installed).
2. Clone this repository or copy the `tree.go` file.
   ```bash
   git clone https://github.com/vovakirdan/tree-view.git
   cd tree-view
   ```
3. Build the utility:
   ```bash
   go build tree.go
   ```

---

## Usage

Run the utility from the command line:

```bash
./tree [options] <directory>
```

### Options:
- `-a`  
   Show hidden files and directories (those starting with a dot `.`).
- `-exclude=<names>`  
   Exclude files and directories by name. Pass a comma-separated list.

### Examples

1. **Display the current directory:**
   ```bash
   ./tree .
   ```

2. **Show hidden files and directories:**
   ```bash
   ./tree -a .
   ```

3. **Exclude specific directories (e.g., `node_modules` and `build`):**
   ```bash
   ./tree -exclude=node_modules,build .
   ```

4. **Combine options:**
   Show hidden files and exclude specific directories:
   ```bash
   ./tree -a -exclude=.git,node_modules .
   ```

---

## Output Example

Given a directory structure:

```
GoCall/
├── go.mod
├── main.go
├── server/
│   └── main.go
├── static/
│   ├── app.js
│   └── index.html
└── .git/
    ├── HEAD
    └── config
```

### Without hidden files:
```bash
./tree GoCall
```

**Output:**
```
GoCall/
├── go.mod
├── main.go
├── server/
│   └── main.go
└── static/
    ├── app.js
    └── index.html
```

### With hidden files:
```bash
./tree -a GoCall
```

**Output:**
```
GoCall/
├── .git/
│   ├── HEAD
│   └── config
├── go.mod
├── main.go
├── server/
│   └── main.go
└── static/
    ├── app.js
    └── index.html
```

---

## Notes
- Hidden files and directories (starting with `.`) are ignored by default.
- The program works recursively and handles any depth of directories.

---

## License

This project is licensed under the MIT License.
