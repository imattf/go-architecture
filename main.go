// Hands-on exercise #3 - Ninja Level 4
// Using fmt.Errorf(), handle all errors in a program that does the following:
// open file #1
// create file #2
// copy the contents of file #1 to file #2
// tag v1.31.0
// video: 92 Hands On 3

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	src := "file1.txt"
	dst := "file2.txt"
	err := copyFile(dst, src)
	if err != nil {
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
	//panic for the fun of it
	log.Panic("for the fun of it")
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
