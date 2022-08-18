// License: GPL-3.0-only
// (c) 2022 Dakota Walsh <kota@nilsu.org>
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func findWords() string {
	paths := []string{
		"/usr/share/thesaurus/words.txt",
		"/usr/local/share/thesaurus/words.txt",
		"/usr/share/thesaurus/thesaurus.txt",
		"/usr/local/share/thesaurus/thesaurus.txt",
	}
	for _, path := range paths {
		info, err := os.Stat(path)
		if err == nil && info.IsDir() == false {
			return path
		}
	}
	return "words.txt"
}

func main() {
	path := flag.String("w", "", "path to thesaurus file")
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "usage: thesaurus -w words.txt word")
		os.Exit(1)
	}
	if *path == "" {
		*path = findWords()
	}

	f, err := os.Open(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed opening thesaurus: %v\n", err)
		os.Exit(1)
	}

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	for {
		record, err := r.Read()
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "failed reading thesaurus: %v\n", err)
			os.Exit(1)
		}

		if len(record) < 1 {
			continue
		}

		if record[0] == flag.Arg(0) {
			for _, match := range record[1:] {
				fmt.Println(match)
			}
			os.Exit(0)
		}
	}
}
