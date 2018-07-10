package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			log.Println(err)
		}
	}
}

func tree(root string, ident string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %v: %v", root, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir: %v: %v", root, err)
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] == '.' {
			continue
		}
		names = append(names, fi.Name())
	}

	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Print(ident + "└──")
			add = "   "
		} else {
			fmt.Print(ident + "├──")
		}

		err = tree(filepath.Join(root, name), ident+add)
		if err != nil {
			return err
		}
	}

	return nil
}
