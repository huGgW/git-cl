# git-cl

Git Branch Cleaner - A tool for efficiently managing accumulated git branches

## Introduction

`git-cl` is a command-line tool designed to efficiently organize and manage Git repositories with numerous accumulated branches. It helps you easily identify and remove unnecessary branches to keep your repository clean during the development process.

## Installation

### Binary Installation (Recommended)
You can download Linux and Mac binaries from the GitHub releases page.


### Install from Source

To build and install directly from source, follow these steps:

#### Prerequisites
- Go 1.24 or higher

```bash
git clone https://github.com/huGgW/git-cl.git
cd git-cl
go install
```

## Usage

### Examples

```bash
# List all local branches
git-cl list

# Preview branches to be deleted
git-cl clean -d

# Clean branches (perform deletion)
git-cl clean

# Protect specific branches from deletion
git-cl clean -b branch1,branch2

# Update remote repository information before cleaning branches
git-cl clean -f
```

### Options for clean command

| Option | Alias | Description |
|--------|-------|-------------|
| `--dry` | `-d` | Preview branches to be deleted without actually deleting them |
| `--blacklist` | `-b` | Specify a list of branches to be excluded from deletion |
| `--fetch` | `-f` | Update remote repository information before cleaning branches |
