package main 

import (
	"fmt"
	"math"
	"sort"
)

type statistics struct {
	numbers []float64
	mean float64
	median float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return stats
}

func sum(numbers []float64) (total float64){
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64{
	middle := len(numbers) / 2
	result := numbers[middle]
	if(len(numbers) % 2 == 0) {
		result = (result + numbers[middle - 1]) / 2
	}
	return result
}

