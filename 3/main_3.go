package main

import "fmt"

type StringIntMap struct {
	hashMap map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		hashMap: make(map[string]int),
	}
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.hashMap[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.hashMap, key)
}

func (sim *StringIntMap) Copy() map[string]int {
	mapCopy := make(map[string]int, len(sim.hashMap))
	for key, value := range sim.hashMap {
		mapCopy[key] = value
	}
	return mapCopy
}

func (sim *StringIntMap) Exists(key string) bool {
	_, ok := sim.hashMap[key]
	return ok

}

func (sim *StringIntMap) Get(key string) (int, bool) {
	value, ok := sim.hashMap[key]
	return value, ok
}

func main() {
	sim := NewStringIntMap()

	sim.Add("One", 1)
	sim.Add("Two", 2)
	fmt.Println(sim.hashMap)

	sim.Remove("One")
	fmt.Println(sim.hashMap)

	newSim := sim.Copy()
	sim.Add("Three", 3)
	fmt.Println(sim)
	fmt.Println(newSim)

	fmt.Println(sim.Exists("One"))
	fmt.Println(sim.Exists("Two"))

	fmt.Println(sim.Get("One"))
	fmt.Println(sim.Get("Two"))

}
