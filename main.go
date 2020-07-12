// Hands-on exercise #5 - Ninja Level 4
// Using the code in the previous example, use errors.As() to access *PathError and print out some fields or run some methods attached to *PathError. Notes:
// https://godoc.org/os#Open
// https://godoc.org/os#Create
// https://godoc.org/os#PathError
// tag v1.31.2
// video: 94 Hands On 5

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
	var perr *os.PathError
	if errors.As(err, &perr) && errors.Is(err, os.ErrNotExist) {
		fmt.Println("you need to provide valid file name of:", src)
	} else if errors.As(err, &perr) {
		fmt.Printf("main exception: error in copyFile: %s - OPERATION: %s - Error: %s\n", perr.Path, perr.Op, err)
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
