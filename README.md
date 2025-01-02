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

## Install in **PATH**
> If you want to use it from anywhere in your linux system!
0. [Compile](#installation). Note that you can specify output file name.
1. Copy (or move) binary to dir that included in PATH. As usual, it is a `/usr/local/bin` 
```bash
sudo cp tree /usr/local/bin/tree
```
---
### For MacOS
You can use also:
```bash
sudo mv tree /usr/local/bin/
```
But sometimes somethings goes wrong and you have no such dir. Then, of course, you can create it:
```bash
sudo mkdir /usr/local/bin
```
An then, if you don't have this path in your `$PATH` go:
1. Open configuration file for your sh. For mac os it is zsh (`~/.zshrc`) by default, for bash it is bash profile (`~/.bash_profile`)
```bash
vim ~/.zshrc
```
2. Add `export PATH=$PATH:/usr/local/bin` to the file if don't even exists
3. Save and run `source ~/.zshrc`
Then you can just move it to the dir.

### Via Brew
1. Move binary file to the homebrew path:
```bash
mkdir -p /usr/local/Cellar/tree/1.0/bin
mv tree /usr/local/Cellar/tree/1.0/bin/
```
2. Create symbol link to utility:
```bash
ln -s /usr/local/Cellar/tree/1.0/bin/tree /usr/local/bin/tree
```
### Finally:
   **Check access via `tree --help`. If everything is ok, you will see the help.**

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
- `-emoji`  
   Shows emoji near the file or path.

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
### With emoji flag
```
./tree -emoji
```

**Output**
```
📂 gocall/
├── 📜 GoCall
├── 📂 client/
│   ├── 💻 main.go
│   └── 💻 webrtc.go
├── 📜 go.mod
├── 📜 go.sum
├── 💻 main.go
├── 📂 server/
│   ├── 💻 main.go
│   ├── 💻 signaling.go
│   └── 💻 webrtc.go
└── 📂 static/
    ├── 💻 app.js
    └── 📜 index.html
```
---

## Notes
- Hidden files and directories (starting with `.`) are ignored by default.
- The program works recursively and handles any depth of directories.

---

## License

This project is licensed under the MIT License.
