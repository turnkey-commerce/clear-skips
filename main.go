package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Clears the test skips from an Exercism c# test.")
		fmt.Println("Usage: clear-skips filenameTest.cs")
		return
	}

	inFile := os.Args[1]
	info, err := os.Stat(inFile)
	if err != nil {
		fmt.Println("Error can't stat file: ", err)
		return
	}

	buf, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("Error can't read file: ", err)
		return
	}

	s := string(buf)
	r := strings.Replace(s, "(Skip = \"Remove to run test\")", "", -1)

	err = ioutil.WriteFile(inFile, []byte(r), info.Mode())
	if err != nil {
		fmt.Println("Error can't write file: ", err)
		return
	}

}
