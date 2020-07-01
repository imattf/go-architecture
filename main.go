// Hands-on exercise #5
// In this exercise, copy the contents from one file to a new file.
// Create a file “file-01.txt” from which to copy.
// To complete this exercise, use
//   io.Copy
//   os.Create
//   os.Open

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Open a file
	f1, err := os.Open("file1.txt")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	// Create a new file
	f2, err := os.Create("file2.txt")
	if err != nil {
		panic(err)
	}

	// Copy contents from original file to new file
	n, err := io.Copy(f2, f1)
	if err != nil {
		panic(err)
	}
	fmt.Println("file2.txt size is:", n)

}
