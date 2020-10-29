//Create more sophisticated but still simple programs using Go and data structure usage.
package main

import  (
	"fmt"
	"unicode"
	)




func main() {
	fmt.Println("Hello, world!")

	fmt.Println("ParsePhone() test")
	fmt.Printf("ParsePhone(%q) = %q\n", "123-456-7890", ParsePhone("123-456-7890"))
	fmt.Printf("ParsePhone(%q) = %q\n", "1 2 3 4 5 6 7 8 9 0", ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	fmt.Println("Anagram() test")
	fmt.Printf("Anagram(%q, %q) = %v\n", "12345", "52314", Anagram("12345", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "21435", "53241", Anagram("21435", "53241"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "12346", "52314", Anagram("12346", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "123456", "52314", Anagram("123456", "52314"))

	fmt.Println("FindEvens() test")
	fmt.Printf("FindEvens(%v) = %v\n", []int{1, 2, 3, 4}, FindEvens([]int{1, 2, 3, 4}))

	fmt.Println("SliceProduct() test")
	fmt.Printf("SliceProduct(%v) = %v\n", []int{5, 6, 8}, SliceProduct([]int{5, 6, 8}))

	fmt.Println("Unique() test")
	fmt.Printf("Unique(%v) = %v\n", []int{1, 2, 3, 4, 4, 5, 6, 6}, Unique([]int{1, 2, 3, 4, 4, 5, 6, 6}))

	fmt.Println("InvertMap() test")
	fmt.Printf("InvertMap(%v) = %v\n", map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}, InvertMap(map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}))
}

// ParsePhone parses a string of numbers into the format 06 22 14 33 44.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "12 34 56 78 90"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "12 34 56 78 90"
func ParsePhone(phone string) string {

	var phoneParsed string
	var count int = 1
	for _, i := range phone {
		if unicode.IsDigit(i) {
			if count%2 == 0 || count == 1 {
				phoneParsed += string(i)
			} else {
				phoneParsed += " " + string(i)
			}
			count++
		}
	}
	return  phoneParsed
}

// Write a function to check whether two given strings are anagram of each other or not. 
// An anagram of a string is another string that contains same characters, 
// only the order of characters can be different. For example, “abcd” and “dabc” are anagram of each other.
// This function is NOT case sensitive and should handle UTF-8

func hash(str1, str2 string) map[string]int {

	hash := make(map[string]int)

	for _, r := range str1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range str2 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}


	return hash
}

func Anagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	} else {

		hash := hash(s1,s2)
		var isAnagram bool = true
		for _, value := range hash {
			if value%2 != 0 {
				isAnagram = false
				break
			}

		}
		return isAnagram

	}
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {

	var res []int
	for _, value := range e {
		if value%2 == 0 {
			res = append(res, value)
		}
	}
	return res
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var res int
	for _, value := range e {
		res += value
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {

	var res []int 
	var comp bool = true

	for i, value := range e {
		comp = true
		for j, value2 := range e {
			if value==value2 && i!=j {
				comp = false
			} 
		}
		if comp {
			res = append(res, value)
		}

	}
	return res
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	var vk = make(map[int]string)

	for j, value2 := range kv {
		vk[value2] = j
	}
	return vk
}