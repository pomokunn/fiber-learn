package exercise

import (
	"reflect"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Test with name Poom", input: "Poom", expected: "Hello, Poom"},
		{name: "Test with empty name", input: "", expected: "Hello, "},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Greeting(tt.input)
			if tt.expected != actual {
				t.Errorf("greeting(%s) = %s; want %s", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestGreetingWithAge(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   uint
		expected string
	}{
		{name: "Test with name Poom and age 24", input1: "Poom", input2: 24, expected: "Hello, Poom. You are 24 years old."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GreetingWithAge(tt.input1, tt.input2)
			if tt.expected != actual {
				t.Errorf("greetingWithAge(%s, %d) = %s; want %s", tt.input1, tt.input2, actual, tt.expected)
			}
		})
	}
}

func TestGreetingWithAgeAndDrink(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   uint
		input3   string
		expected string
	}{
		{name: "Test with name Poom, age 24, and drink water", input1: "Poom", input2: 24, input3: "water", expected: "Hello, Poom. You are 24 years old and your favorite drink is water."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GreetingWithAgeAndDrink(tt.input1, tt.input2, tt.input3)
			if tt.expected != actual {
				t.Errorf("greetingWithAgeAndDrink(%s, %d, %s) = %s; want %s", tt.input1, tt.input2, tt.input3, actual, tt.expected)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "Test with odd number 5", input: 5, expected: true},
		{name: "Test with even number 6", input: 6, expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsOdd(tt.input)
			if tt.expected != actual {
				t.Errorf("isOdd(%d) = %t; want %t", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestSumOfFirst(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "Test with n = 3", input: 3, expected: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SumOfFirst(tt.input)
			if tt.expected != actual {
				t.Errorf("isOdd(%d) = %d; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{name: "Test with prime number 3", input: 3, expected: true},
		{name: "Test with non-prime number 6", input: 6, expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsPrime(tt.input)
			if tt.expected != actual {
				t.Errorf("isOdd(%d) = %t; want %t", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestDouble(t *testing.T) {
	// Test case 1: Doubling a positive number
	n := 2
	Double(&n)
	if n != 4 {
		t.Errorf("double(&n) failed, expected %d, got %d", 4, n)
	}
	// Test case 2: Doubling a negative number
	n = -3
	Double(&n)
	if n != -6 {
		t.Errorf("double(&n) failed, expected %d, got %d", -6, n)
	}
	// Test case 3: Doubling zero
	n = 0
	Double(&n)
	if n != 0 {
		t.Errorf("double(&n) failed, expected %d, got %d", 0, n)
	}
}

func TestAppendGreeting(t *testing.T) {
	// Test case 1: Appending greeting to a non-empty string
	name := "Poom"
	AppendGreeting(&name)
	expected := "Hi, Poom"
	if name != expected {
		t.Errorf("appendGreeting(&s) failed, expected %s, got %s", expected, name)
	}
	// Test case 2: Appending greeting to an empty string
	name = ""
	AppendGreeting(&name)
	expected = "Hi, "
	if name != expected {
		t.Errorf("appendGreeting(&s) failed, expected %s, got %s", expected, name)
	}
}

func TestAbc(t *testing.T) {
	expected := []string{"apple", "banana", "coconut"}
	result := []string{"apple", "banana", "coconut"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Abc() failed, expected %v, got %v", expected, result)
	}
}

func TestEfg(t *testing.T) {
	expected := []string{"elderberries", "figs", "grapes"}
	result := []string{"elderberries", "figs", "grapes"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Efg() failed, expected %v, got %v", expected, result)
	}
}

func TestCde(t *testing.T) {
	expected := []string{"coconut", "durian", "elderberries"}
	result := []string{"coconut", "durian", "elderberries"}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Cde() failed, expected %v, got %v", expected, result)
	}
}

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		// Simple case with repeated and unique words
		{name: "simple case", input: "hello world hello", expected: map[string]int{"hello": 2, "world": 1}},
		// Empty string should result in an empty map
		{name: "empty string", input: "", expected: map[string]int{}},
		// Handling punctuation as part of words
		{name: "punctuation", input: "hello, world! hello.", expected: map[string]int{"hello,": 1, "world!": 1, "hello.": 1}},
		// Multiple spaces between words
		{name: "multiple spaces", input: "hello   world    hello", expected: map[string]int{"hello": 2, "world": 1}},
		// Only whitespace should result in an empty map
		{name: "whitespace only", input: "   ", expected: map[string]int{}},
		// Large input to test performance
		{name: "large input", input: strings.Repeat("word ", 1000), expected: map[string]int{"word": 1000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := WordCount(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("wordCount(%q) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestBookString(t *testing.T) {
	// สร้างตัวแปร book
	book := Book{Name: "Harry Potter", Author: "J. K. Rowling"}

	// คาดหวังผลลัพธ์
	expected := "Harry Potter by J. K. Rowling"

	// เรียกใช้งาน method String() และตรวจสอบผลลัพธ์
	if result := book.String(); result != expected {
		t.Errorf("String() failed, expected %s, got %s", expected, result)
	}
}

func TestBookSetName(t *testing.T) {
	// สร้างตัวแปร book
	book := Book{Name: "Harry Potter", Author: "J. K. Rowling"}

	// เปลี่ยนชื่อหนังสือ
	newName := "Harry Potter and the Goblet of Fire"
	book.SetName(newName)

	// คาดหวังผลลัพธ์
	expected := "Harry Potter and the Goblet of Fire"

	// ตรวจสอบผลลัพธ์
	if result := book.Name; result != expected {
		t.Errorf("SetName() failed, expected %s, got %s", expected, result)
	}
}
