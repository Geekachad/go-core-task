package main

import (
	"testing"
)

func TestGenerateHash(t *testing.T) {
	values := []any{
		42,
		052,
		0x2A,
		3.14,
		"Golang",
		true,
		complex64(1 + 2i),
	}
	expectedHash := "3f57eddce9cfca3285a46322d2480f81ea058c318b53da8a493ea0d2cc75f733"

	hash := GenerateHash(values)

	if hash != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, hash)
	}
}

func TestGenerateHash_NilInput(t *testing.T) {
	var values []any

	expectedHash := "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6"

	hash := GenerateHash(values)

	if hash != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, hash)
	}
}

func TestGenerateHash_EmptyInput(t *testing.T) {
	values := []any{}
	expectedHash := "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6"

	hash := GenerateHash(values)

	if hash != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, hash)
	}
}

func TestGenerateHash_SameInput(t *testing.T) {
	values := []any{
		42,
		"Golang",
		true,
	}

	hash1 := GenerateHash(values)
	hash2 := GenerateHash(values)

	if hash1 != hash2 {
		t.Errorf("hashes are different")
	}
}

func TestGenerateHash_DifferentInput(t *testing.T) {
	values1 := []any{"43"}
	values2 := []any{43}

	hash1 := GenerateHash(values1)
	hash2 := GenerateHash(values2)

	if hash1 == hash2 {
		t.Error("hashes should be different")
	}
}

func TestGenerateHash_SameTypeInput(t *testing.T) {
	values1 := []any{42}
	values2 := []any{43}

	hash1 := GenerateHash(values1)
	hash2 := GenerateHash(values2)

	if hash1 != hash2 {
		t.Error("hashes should be same")
	}
}

func TestGenerateHash_OrderMatters(t *testing.T) {
	values1 := []any{42, "Golang"}
	values2 := []any{"Golang", 42}

	hash1 := GenerateHash(values1)
	hash2 := GenerateHash(values2)

	if hash1 == hash2 {
		t.Error("hashes should be different")
	}
}

func TestGenerateHash_Length(t *testing.T) {
	hash := GenerateHash([]any{42})

	if len(hash) != 64 {
		t.Errorf("expected length 64, got %d", len(hash))
	}
}
