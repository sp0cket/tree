# Tree

tree is a command line tool that produces a depth indented listing of files.

## Usage

```bash
$ tree --help
Usage of ./tree:
  -a	List all files
  -depth int
    	Max depth
  -o string
    	Output file, default to console
  -p string
    	Working path, default is current path (default ".")
```

## Example
```bash
$ tree
[.]
├─ LICENSE
├─ README.md
├─ cmd/
│  ├─ cmd.go
│  └─ node.go
├─ dir/
│  └─ visit.go
├─ go.mod
├─ go.sum
└─ main.go
```

