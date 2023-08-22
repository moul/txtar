package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/txtar"
)

func main() {
	var (
		flagStripPrefix = flag.String("strip", "", "string which remove from head of path")
	)
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		fmt.Fprintln(os.Stderr, "target directory must be specified")
		os.Exit(1)
	}

	var ar txtar.Archive

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

		p := filepath.ToSlash(path)
		ar.Files = append(ar.Files, txtar.File{
			Name: strings.TrimPrefix(p, *flagStripPrefix),
			Data: data,
		})

		return nil
	})
	if err != nil {
		panic(err)
	}

	if len(ar.Files) == 0 {
		fmt.Fprintln(os.Stderr, "target directory is empty")
		os.Exit(1)
	}
	archived := string(txtar.Format(&ar))
	fmt.Println(archived)
}
