package main

import (
	"golang.org/x/tour/wc"
	"strings" //need this for Fields
)

func WordCount(s string) map[string]int {
	string_slice := strings.Fields(s) //declare a slice containing all words in string s (Fields breaks the string into a slice, minus the annoying ' ') 
	result := make(map[string]int)
	n := len(string_slice) //get length of slice (basically length of original string, minus ' ' )
	for i :=0; i<n; i++ {
		result[string_slice[i]] ++ //counts the amount of time
								// each word appears in the string
	}
	
	return result
}

func main() {
	wc.Test(WordCount)
}
