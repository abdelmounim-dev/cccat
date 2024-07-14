package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getFlagsAndFiles(args []string) ([]string, []string) {
	flags := []string{}
	files := []string{}
	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
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

func printLines(scanner *bufio.Scanner, numberNonBlank bool) {
	n := 1
	for scanner.Scan() {
		text := scanner.Text()
		if numberNonBlank {
			if text != "" {
				fmt.Printf("%d %s\n", n, text)
				n++
			} else {
				fmt.Println(text)
			}
		} else {
			fmt.Printf("%d %s\n", n, text)
			n++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func printNormal(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func readFromReader(r io.Reader, flags []string) {
	scanner := bufio.NewScanner(r)
	switch numberType(flags) {
	case 'n':
		printLines(scanner, false)
	case 'b':
		printLines(scanner, true)
	default:
		printNormal(scanner)
	}
}

func main() {
	flags, files := getFlagsAndFiles(os.Args[1:])
	if len(files) == 0 {
		readFromReader(os.Stdin, flags)
		return
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			continue
		}
		readFromReader(f, flags)
		f.Close()
	}
}
