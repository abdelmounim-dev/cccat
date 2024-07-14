package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func isFlag(arg string) bool {
	return len(arg) > 0 && arg[0] == '-'
}

func getFlagsAndFiles(args []string) ([]string, []string) {
	flags := []string{}
	files := []string{}
	for _, arg := range args {
		if isFlag(arg) {
			flags = append(flags, arg)
		} else {
			files = append(files, arg)
		}
	}
	return flags, files
}

func numberType(flags []string) rune {
	for _, flag := range flags {
		if flag == "-n" {
			return 'n'
		}
		if flag == "-b" {
			return 'b'
		}
	}
	return 'g'
}

func printNormal(scanner bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}
}

func printNumbered(scanner bufio.Scanner) {
	n := 1
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(n, " ", text)
		n++
	}
}

func printNumberedNonBlank(scanner bufio.Scanner) {
	n := 1
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			fmt.Println(n, " ", text)
			n++
		} else {
			fmt.Println(text)
		}
	}
}

func readFromReader(r io.Reader, flags []string) {
	scanner := bufio.NewScanner(r)

	nt := numberType(flags)
	if nt == 'n' {
		printNumbered(*scanner)
	} else if nt == 'b' {
		printNumberedNonBlank(*scanner)
	} else {
		printNormal(*scanner)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func main() {
	flags, files := getFlagsAndFiles(os.Args[1:])
	if len(files) == 0 {
		readFromReader(os.Stdin, flags)
		return
	}

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		readFromReader(file, flags)
	}
}
