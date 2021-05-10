Skeleton is a CLI utility that takes a directory tree with go template files and creates an output 
directory tree executing the templates based on user supplied values

# Table of Contents

- [Overview](#overview)
- [Usage](#Usage)
- [Installing](#installing)
- [License](#license)

# Overview

Skeleton is a CLI tool which creates a directory tree structure based on a template and dynamically
updates template files based on user supplied values

Skeleton provides:
* Easy project bootstrapping: no more tedious copying files
* Dynamic value interpolation: no more tedious updating names of things
* The flexibility to define your own directory and files structure.

## Usage

```bash
Usage:
  skeleton create [flags] TEMPLATES_DIR OUTPUT_DIR

Flags:
  -h, --help            help for create
  -v, --values string   path to values.yaml file
```

Lets look at an example with a templates directory that looks like this:

```
templates
├── LICENSE.tpl
├── Makefile.tpl
├── README.md.tpl
├── cmd
│   └── root.go.tpl
├── go.mod.tpl
└── main.go.tpl
```

The structure of the templates directory will match the structure of the output directory, which
will be created in the path the user supplies.

The above templates directory would result in the following: 

```
outputDir
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── root.go
├── go.mod
└── main.go
```

### Note: Template files must have a `.tpl` suffix or they won't be processed as a template.
This means that if there is a file you simply want to include in your skeleton project that doesn't need any
templating, like a LICENCE file, simply leave off the `.tpl` suffix and it will be copied over without any 
extra processing.


Skeleton uses Go templates for templating your output files [docs](https://golang.org/pkg/text/template/). 
It also includes the sprig template functions library to provide additional functionality [docs](http://masterminds.github.io/sprig/)

# Installing
Visit the current release

# License

Skeleton is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/shaneu/skeleton/blob/master/LICENSE)