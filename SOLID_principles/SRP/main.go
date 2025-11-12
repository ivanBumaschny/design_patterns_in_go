// SRP
// Single Responsibility Principle (SRP)

// A type must have only one responsibility
package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Separation Of Concerns
// God Object
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 1)
}

func (j *Journal) Load(filename string) {
	// ..
}

// You dont want to have persistence methods within the class
// You would do it like this, if needed
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")

	//
	//SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
