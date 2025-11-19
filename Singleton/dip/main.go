package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

// This will only be called once and then its done
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("./database.txt")
		db := singletonDatabase{caps}
		if e != nil {
			panic(e)
		}
		db.capitals = caps
		instance = &db
	})

	// we keep returning the same instance pointer when multiple calls are made
	return instance
}

func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (db *DummyDatabase) GetPopulation(name string) int {
	if len(db.dummyData) == 0 {
		db.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}

	return db.dummyData[name]
}

// The "scary" part of a singleton is dependency on a global state, which makes testing difficult. Using DIP you solve this issue, by providing an API that depends on abstractions (interface) rather than concretions (singleton instance)
func main() {
	db := &DummyDatabase{}
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulation(db, names)
	fmt.Println(tp == 4) // alpha + gamma == 1 + 3 == 4
}
