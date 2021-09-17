package main

import "fmt"

func main() {

	var input string
	fmt.Scan(&input)

	var result []string
	m := make(map[string]int)

	var greska bool

	for i := 0; i < len(input); i++ {
		if i%3 == 0 {
			m[string(input[i])] += 1

			var card string = input[i : i+3]
			if contains(result, card) {
				fmt.Println("GRESKA")
				greska = true
				break
			} else {
				result = append(result, card)
			}
		}
	}
	if !greska {
		fmt.Printf("%d"+" "+"%d"+" "+"%d"+" "+"%d", 13-m["P"], 13-m["K"], 13-m["H"], 13-m["T"])
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
