package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error whiile creating file:", err)
		return
	}
	defer file.Close()

	content := "hello world"
	byte, error := io.WriteString(file, content+"\n")
	fmt.Println("Number of bytes written:", byte)
	if error != nil {
		fmt.Println("Error while writing to file:", error)
		return
	}

	fmt.Println("File created successfully:")

	*/
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}
		fmt.Println(string (buffer[:n]))
	}





}

// this code creates a file named "test.txt" and writes "hello world" to it and returns the number of bytes written.
// it then opens the file and reads its content, printing it to the console.
