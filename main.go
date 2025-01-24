package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/tools/txtar"
)

func main() {
	var (
		flagStripPrefix = flag.String("strip", "", "string which remove from head of path")
		flagListOnly    = flag.Bool("list", false, "only list matching files without creating archive")
	)
	flag.Parse()

	var ar txtar.Archive
	dirs := flag.Args()

	if len(dirs) == 0 {
		// Read files from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			path := scanner.Text()
			data, err := ioutil.ReadFile(path)
			if err != nil {
				continue // Skip files we can't read
			}

			if !utf8.Valid(data) {
				// TODO: log skipped files
				continue
			}

			p := filepath.ToSlash(path)
			if *flagListOnly {
				fmt.Println(strings.TrimPrefix(p, *flagStripPrefix))
				continue
			}

			ar.Files = append(ar.Files, txtar.File{
				Name: strings.TrimPrefix(p, *flagStripPrefix),
				Data: data,
			})
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error reading stdin:", err)
			os.Exit(1)
		}
	} else {
		for _, dir := range dirs {
			err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				base := filepath.Base(path)

				if info.IsDir() {
					if len(base) > 1 && base[0] == '.' && base[1] != '.' && base[1] != '/' {
						return filepath.SkipDir
					}
					return nil
				}

				if len(base) > 0 && base[0] == '.' {
					return nil
				}

				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				if !utf8.Valid(data) {
					// TODO: log skipped files
					return nil
				}

				p := filepath.ToSlash(path)
				if *flagListOnly {
					fmt.Println(strings.TrimPrefix(p, *flagStripPrefix))
					return nil
				}

				ar.Files = append(ar.Files, txtar.File{
					Name: strings.TrimPrefix(p, *flagStripPrefix),
					Data: data,
				})

				return nil
			})
			if err != nil {
				panic(err)
			}
		}
	}

	if !*flagListOnly {
		if len(ar.Files) == 0 {
			fmt.Fprintln(os.Stderr, "no files found")
			os.Exit(1)
		}
		archived := string(txtar.Format(&ar))
		fmt.Println(archived)
	}
}
