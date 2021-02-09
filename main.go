package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createFile()
	writeFile()
	readFile()
}

func createFile() {
	var file, err = os.Create("./files/file.txt")
	defer file.Close() //Close file created
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Created Successfully")
}

/*
	permission:
	os.O_RDONLY -> open the file read-only. (0)
	os.O_WRONLY -> open the file write-only. (1)
	os.O_RDWR   -> open the file read-write. (2)
*/

func writeFile() {
	var file, err = os.OpenFile("./files/file.txt", 2, 644)
	defer file.Close() //Close file after write
	if err != nil {
		log.Fatal(err)
	}

	var _, errC = file.WriteString("Insert text here")
	if errC != nil {
		log.Fatal(errC)
	}
	fmt.Println("File Updated Successfully")
}

func readFile() {
	var file, err = os.OpenFile("./files/file.txt", 2, 644)
	defer file.Close() //Close file after read
	if err != nil {
		log.Fatal(err)
	}

	var content, errC = ioutil.ReadAll(file)
	if errC != nil {
		log.Fatal(errC)
	}
	fmt.Println("Content:", string(content))
}
