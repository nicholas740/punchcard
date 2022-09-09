package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

// Punchcard instantiates a basic `Punchcard` structure for now. It implements `PunchcardService`.
type Punchcard struct {
	PunchcardService
}

// PunchcardService defines the methods for all of the functions, this is for clean architecture.
type PunchcardService interface {
	add(punchcard []string, data string) []string
	read(filename string) ([]byte, error)
	parse(data []byte) ([]string, error)
	write(punchcard []string, filename string) error
}

// NewPunchcard creates a new punchcard instance, the methods has to be defined.
func NewPunchcard() PunchcardService {
	return &Punchcard{}
}

// add appends a new data to the makeshift database / punchcard.
func (p *Punchcard) add(punchcard []string, data string) []string {
	updatedData := make([]string, 0)
	updatedData = append(updatedData, punchcard...)
	updatedData = append(updatedData, data)

	return updatedData
}

// read opens and reads the makeshift database / punchcard.
func (p *Punchcard) read(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// parse parses the data in `[]byte` format properly. It will be encoded into a struct.
func (p *Punchcard) parse(data []byte) ([]string, error) {
	attendances := []string{}
	if err := json.Unmarshal(data, &attendances); err != nil {
		return nil, err
	}

	return attendances, nil
}

// write writes data to a file with buffers. `MarshalIndent` is 2 spaces, while
// the file permission is the usual Linux file permission, or 0644 (`-rw-r--r--`).
func (p *Punchcard) write(data []string, filename string) error {
	indentedData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, indentedData, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}

// main is the driver code to run the whole script properly.
//
// Algorithm:
//  1. Define or instantiate all required variables.
//  2. Open the file.
//  3. Parse the file, store it into a slice of structures.
//  4. Append new data to that slice of structures.
//  5. Store that slice of structures to the previous file (we will replace the whole contents).
func main() {
	filename := path.Join("data", "punchcard.json")
	timezone := os.Getenv("TZ")
	punchcard := NewPunchcard()

	if timezone == "" {
		os.Setenv("TZ", "Asia/Jakarta")
	} else {
		os.Setenv("TZ", timezone)
	}

	rawPunchcard, err := punchcard.read(filename)
	if err != nil {
		panic(err.Error())
	}

	parsedPunchcard, err := punchcard.parse(rawPunchcard)
	if err != nil {
		panic(err.Error())
	}

	updatedData := punchcard.add(parsedPunchcard, time.Now().String())
	if err := punchcard.write(updatedData, filename); err != nil {
		panic(err.Error())
	}

	fmt.Println("→ Punchcard script has finished running!")
}
