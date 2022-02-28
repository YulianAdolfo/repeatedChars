package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := strings.Split("Yulian Adolfo Rojas Gañan", "")
	temporal := []string{}
	var char string
	var count int
	for i := 0; i < len(sentence); i++ {
		char = strings.ToLower(sentence[i])
		if char != " " {
			for j := 0; j < len(sentence); j++ {
				if char == strings.ToLower(sentence[j]) {
					count++
				} else {
					if sentence[j] != " " {
						temporal = append(temporal, sentence[j])
					}
				}
			}
			if count == 1 {
				count = 0
			}
			sentence = temporal
			i--
			fmt.Println("Letra: ", char, " se repite en ", count, " veces")
			count = 0
			temporal = []string{}
		}
	}
}
