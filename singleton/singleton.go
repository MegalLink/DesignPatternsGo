package singleton

import "sync"

// singletons are used when we need just one instance in entire project, like when we are using databse
// but i dont recommend use this patter because is difficult to test, we can't mock this dabase, it violates dependecy inversion principle
//, so better use dependency injection

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once // we can use a initMethod too, but sync.Once is thread safety , an guarante laziness
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		capitals := map[string]int{"zero": 0, "one": 1, "two": 2}
		db := singletonDatabase{capitals: capitals}
		db.capitals = capitals
		instance = &db
	})

	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city) //violates dependecy inversion principle
	}
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city) //we can provide a mock of database
	}
	return result
}
