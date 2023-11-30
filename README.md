# Subtext Occurrence Finder

This simple Go program allows you to find and print the positions of occurrences of a subtext within a given text. It prompts the user to enter the text to search and the subtext to find, then prints the positions of occurrences.

## Table of Contents

- [Installation of Go](#installation-of-go)
  - [Mac](#mac)
  - [Linux](#linux)
- [Compile and Run](#compile-and-run)
- [Usage](#usage)

## Installation of Go

### Mac

1. Open a terminal.
2. Install Homebrew (if not installed) by running the following command:
```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
3. Install Go using Homebrew:
```
brew install go
```

### Linux

1. Open a terminal.
2. Update the package list:
```sh
sudo apt update
```
3. Install Go using the package manager:
```sh
sudo apt install golang
```

### Compile and Run

1. Open a terminal.
2. Change to the directory containing the `main.go` file.
3. Compile the code using the following command:
```sh
go build main.go
```
4. Run the compiled executable:
```sh
./main
```
Alternatively, you can run the code directly using:
```sh
go run main.go
```

### Usage

1. Enter the text to search when prompted.
2. Enter the subtext to find when prompted.
3. The program will output the positions of occurrences of the subtext in the given text.

Note: The positions are printed in a comma-separated format. If there are no occurrences, no output will be displayed.
