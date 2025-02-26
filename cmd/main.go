package main

import (
    "fmt"
    "os"
)

const kind_deploymnent = "Deployment"
const kind_service = "Service"
const kind_ingress = "Ingress"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CreateFile(filename string) {
	fmt.Println("Creating file")
	file, err := os.Create(filename)

	check(err)

	fmt.Printf("File with %s name created\n", file.Name())
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var kind_test string = "Deployment" 

func main() {
	if !FileExists(kind_test + ".yaml") {
		CreateFile(kind_test + ".yaml")
	} else {
		fmt.Println("File already exists")
	}
}
