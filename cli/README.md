# cli

## Installation

Assuming you have `go` installed and properly configured on your system, installation is as simple as navigating to this directory, and running,

```go
go install
```

You should now have access to the `courses` command. You can test this by running,

```bash
courses --help

# Output

courses is a CLI which lets you convert course notes between formats,
	quickly set up new courses, and write beautiful notes.

Usage:
  courses [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     Convert course notes from .tex to .mdx
  help        Help about any command

Flags:
  -h, --help     help for courses
  -t, --toggle   Help message for toggle

Use "courses [command] --help" for more information about a command.
```

## Usage

> All `courses` commands should be run from the root directory.

The main usage of the cli will be in conversion of files from `.tex` to `.mdx` format. In order to convert a chapter,

```bash
courses convert <path-to-chapter>
```

### Continuous conversion

It is possible to continuously convert a course using the `--continuous` flag. For example,

```bash
courses convert -c notes/mlnn
```
