package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func bitCoinExchange(word rune) int {
	switch word {
	case 'a', 'A':
		return 1
	case 'e', 'E':
		return 1
	case 'i', 'I':
		return 2
	case 'o', 'O':
		return 3
	case 'u', 'U':
		return 4
	default:
		return 0
	}
}

func main() {
	for _, user := range users {
		total := 0
		for _, word := range user {
			total += bitCoinExchange(word)
		}
		distribution[user] = total
	}
	fmt.Printf("%s", distribution)
}
