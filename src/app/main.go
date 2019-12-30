package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type error interface {
	Error() string
}

// Sync CAs
func createCertificatesDir() (ok bool, message string) {
	newpath := filepath.Join(".", "grid-security/certificates")
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

func syncCAs() {
	ok, message := createCertificatesDir()
	if ok {
		fmt.Println(message)
	}
	fmt.Println(message)
}

// END Sync CAs

// Sync VOMS
type VO struct {
	voName string
}

// END Sync VOMS

func main() {
	//webRun()
	//syncCAs()

}
