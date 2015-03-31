package main 

import (
	"fmt"
	"math"
	"sort"
	"http"
)

type statistics struct {
	numbers []float64
	mean float64
	median float64
}

const form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`;

const anError = `<p class="error">%s</p>`

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

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprint(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprint(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}
