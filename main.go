package main

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

const cscc = "C:\\windows\\cscc.dat"
const infpub = "C:\\windows\\infpub.dat"
const readOnly = os.FileMode(os.O_RDONLY)

func main() {
	if runtime.GOOS != "windows" {
		log.Fatalln("This killswitch/vaccine is only for Windows. No files created.")
	}

	if _, err := os.Stat(cscc); os.IsNotExist(err) {
		ioutil.WriteFile(cscc, []byte(""), readOnly)
	} else {
		log.Printf("File [%s] already exists. No action needed", cscc)
	}
	if _, err := os.Stat(infpub); os.IsNotExist(err) {
		ioutil.WriteFile(infpub, []byte(""), readOnly)
	} else {
		log.Printf("File [%s] already exists. No action needed", infpub)
	}
}
