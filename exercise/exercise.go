package exercise

import (
	"fmt"
	"strings"
)

// greeting("Pallat") should return "Hello, Pallat"
func Greeting(name string) string {
	return "Hello, " + name
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func GreetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func GreetingWithAgeAndDrink(name string, age uint, drink string) string {
	if drink == "" {
		return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
	}
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
}

func IsOdd(n int) bool {
	return n%2 != 0
}

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6
func SumOfFirst(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 { // Increment by 2 to check only odd numbers
		if n%i == 0 {
			return false
		}
	}
	return true
}

// n := 2
// double(&n)
// n == 4
func Double(n *int) {
	if n != nil {
		*n = *n * 2
	}
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
func AppendGreeting(s *string) {
	if s != nil {
		*s = "Hi, " + *s
	}
}

func Abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed
	s = s[0:3]
	fmt.Println(s)
	// [apple banana coconut]
}

func Efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed
	s = s[4:7]
	fmt.Println(s)
	// * [elderberries figs grapes]
}

func Cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed
	s = s[2:5]
	fmt.Println(s)
	// * [coconut durian elderberries]
}

// TODO: split words and count them
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

func (b Book) String() string {
	// คืนค่า string ที่ประกอบด้วย Name และ Author
	return b.Name + " by " + b.Author
}

func (b *Book) SetName(name string) {
	// เปลี่ยนค่า Name ของ book
	b.Name = name
}
