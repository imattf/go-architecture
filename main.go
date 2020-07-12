// Hands-on exercise #4 - Ninja Level 4
// Using the code in the previous example, use errors.Is() to give a special message if the file you are trying to open does not exist. Notes:
// https://godoc.org/os#pkg-variables
// tag 1.31.1
// video: 93 Hands On 4

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	src := "file1.txt"
	dst := "file2.txt"
	err := copyFile(dst, src)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("you need to provide valid file name of:", src)
	} else if err != nil {
		log.Panicln("main exception: copyFile returned an error - ", err)
	}

}

func copyFile(dst, src string) error {
	f1, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Can't open file in Copyfile: %w", err)
	}
	defer f1.Close()

	f2, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Can't create a file in Copyfile: %w", err)
	}
	defer f2.Close()

	n, err := io.Copy(f2, f1)
	if err != nil {
		return fmt.Errorf("Can't copy a file in Copyfile: %w", err)
	}
	defer f2.Close()

	fmt.Println("Copy successful, bytes written:", n)

	return nil

}
