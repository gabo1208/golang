# Git CLI

This cli aims to make life easier when working with git in go

## Pre requisites

- [make](https://www.gnu.org/software/make/manual/make.html)
- [git](https://github.com/git-guides/install-git)
- [golang](https://go.dev/)

## Getting Started

### Installation

To install git-cli run:

```sh
./scripts/build.sh  
```

Now copy the `git-cli` binary to a directory on your `PATH` (i.e, `/usr/local/bin`) and make sure its filename is `git-cli`. i.e:
```sh
mv git-cli /usr/local/bin
git-cli -h
```

### Usage

run `git-cli -h` to see all the available commands

```sh
Simplify and automate git tasks with go.

Usage:
  git-cli [command]

Available Commands:
  completion      Generate the autocompletion script for the specified shell
  help            Help about any command
  test    test package
  test-other other test package

Flags:
  -h, --help   help for git-cli

Use "git-cli [command] --help" for more information about a command.

```

## Next Steps

- add more functionality
