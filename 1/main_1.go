package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	var allTypes []interface{} = []interface{}{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	hashHex := GenerateHash(allTypes)
	fmt.Println("SHA256:", hashHex)
}

func GenerateHash(value []any) string {
	sb := strings.Builder{}

	for _, value := range value {
		typeStr := fmt.Sprintf("%T", value)
		sb.WriteString(typeStr)
		fmt.Println(typeStr)
	}

	str := sb.String()
	fmt.Println(str)

	runes := []rune(str)
	fmt.Printf("%c\n", runes)

	n := len(runes)
	middle := n / 2
	salt := "go-2024"
	saltRunes := []rune(salt)

	newRunes := make([]rune, 0, n+len(saltRunes))
	newRunes = append(newRunes, runes[:middle]...)
	newRunes = append(newRunes, saltRunes...)
	newRunes = append(newRunes, runes[middle:]...)

	resultString := string(newRunes)
	hash := sha256.Sum256([]byte(resultString))
	return hex.EncodeToString(hash[:])
}
