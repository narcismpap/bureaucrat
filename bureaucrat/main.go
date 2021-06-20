// Bureaucr.at Coding Challenge
// Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/
// London, Jun 2021
// github.com/narcismpap/bureaucrat

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var sourcePath string
	var ref1 string
	var ref2 string

	flag.StringVar(&sourcePath, "s", "/GoT.json", ".json staff directory file")
	flag.StringVar(&ref1, "l", "GoT-005", "Staff #1 ref")
	flag.StringVar(&ref2, "r", "GoT-006", "Staff #2 ref")

	flag.Parse()

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	sourceData, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	directory := &Staff{}
	if json.Unmarshal(sourceData, &directory) != nil {
		log.Fatal("Unable to parse source")
	}

	d := NewDirectoryQuery(directory, StaffReference(ref1), StaffReference(ref2))
	res, err := d.CommonManager()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
