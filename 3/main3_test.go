package main

import (
	"maps"
	"testing"
)

func TestAdd(t *testing.T) {
	sim := NewStringIntMap()

	sim.Add("One", 1)

	if sim.hashMap["One"] != 1 {
		t.Errorf("expected %d, got %v", 1, sim.hashMap["One"])
	}
}

func TestAdd_AlreadyExisting(t *testing.T) {
	sim := NewStringIntMap()

	sim.Add("One", 1)
	sim.Add("One", 1)

	if sim.hashMap["One"] != 1 {
		t.Errorf("expected %d, got %v", 1, sim.hashMap["One"])
	}
}

func TestAdd_Rewrite(t *testing.T) {
	sim := NewStringIntMap()

	sim.Add("One", 1)
	sim.Add("One", 2)

	if sim.hashMap["One"] != 2 {
		t.Errorf("expected %d, got %v", 2, sim.hashMap["One"])
	}
}

func TestAdd_EmptyKey(t *testing.T) {
	sim := NewStringIntMap()

	sim.Add("", 1)

	if sim.hashMap[""] != 1 {
		t.Errorf("expected %d, got %v", 1, sim.hashMap["One"])
	}
}

func TestRemove(t *testing.T) {
	sim := NewStringIntMap()
	wantSim := NewStringIntMap()
	wantSim.hashMap["Two"] = 2

	sim.hashMap["One"] = 1
	sim.hashMap["Two"] = 2
	sim.Remove("One")

	if !maps.Equal(wantSim.hashMap, sim.hashMap) {
		t.Errorf("expected %v, got %v", wantSim.hashMap, sim.hashMap)
	}
}

func TestRemove_ElementNotExist(t *testing.T) {
	sim := NewStringIntMap()
	wantSim := NewStringIntMap()

	sim.Remove("One")

	if !maps.Equal(wantSim.hashMap, sim.hashMap) {
		t.Errorf("expected %v, got %v", wantSim.hashMap, sim.hashMap)
	}
}

func TestCopy(t *testing.T) {
	sim := NewStringIntMap()
	wantSim := NewStringIntMap()
	wantSim.hashMap["One"] = 1
	wantSim.hashMap["Two"] = 2

	sim.hashMap["One"] = 1
	sim.hashMap["Two"] = 2

	result := sim.Copy()

	sim.hashMap["Three"] = 3

	if !maps.Equal(result, wantSim.hashMap) {
		t.Errorf("expected %v, got %v", wantSim.hashMap, result)
	}
}

func TestExists(t *testing.T) {
	sim := NewStringIntMap()
	sim.hashMap["One"] = 1

	result := sim.Exists("One")

	if !result {
		t.Errorf("expected %t, got %t", true, result)
	}
}

func TestExists_NonExist(t *testing.T) {
	sim := NewStringIntMap()

	result := sim.Exists("Two")

	if result {
		t.Errorf("expected %t, got %t", false, result)
	}
}

func TestGet(t *testing.T) {
	sim := NewStringIntMap()
	sim.hashMap["One"] = 1

	result, ok := sim.Get("One")

	if result != 1 {
		t.Errorf("expected %d, got %d", 1, result)
	}
	if !ok {
		t.Errorf("expected %t, got %t", true, ok)
	}
}

func TestGet_NonExist(t *testing.T) {
	sim := NewStringIntMap()

	result, ok := sim.Get("Two")

	if result != 0 {
		t.Errorf("expected %d, got %d", 0, result)
	}
	if ok {
		t.Errorf("expected %t, got %t", false, ok)
	}
}
