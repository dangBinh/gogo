package main 

import (
	"fmt"
	"sort"
	"net/http"
	"log"
	"strings"
	"strconv"
	"math"
)

type statistics struct {
	numbers []float64
	mean float64
	median float64
	mode []float64
	stdDev float64
	// stdDev float64
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
	stats.mode = mode(numbers)
	stats.stdDev = stdDev(numbers)
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

func mode(numbers []float64) []float64{
	var result []float64
	var final []float64
	var (
		i int
		j int
		temp float64
		bound int
	)
	// sap xep lai mang theo thu tu tang dan 
	// tinh mode theo buoc nhay neu buoc nhay la 1 thi khong luu vao mang
	// neu buoc nhay > 1 thi luu vao mang 
	// lay gia tri lon nhat trong mang -> mode
	if len(numbers) > 0 {
		for i = 0; i < len(numbers) - 1; i++ {
			for j = 1; j < len(numbers); j++ {
				if numbers[i] > numbers [j] {
					temp = numbers[i]
					numbers[i] = numbers[j]
					numbers[j] = temp
				}
			}
		}
		i = 0
		for i < len(numbers) {
			bound = 0
			for j = i; j < len(numbers); j++ {
				if numbers[i] == numbers[j] {
					bound++
				}
			}
			if bound >= 2 {
				result = append(result, numbers[i], float64(bound))
				i = i + bound
			} else {
				i ++
			}
		}
		if len(result) > 0 {
			i = 0
			max := result[i+1]
			for i = 0; i < len(result); i += 2 {
				if max < result[i+1] {
					max = result[i+1]
				}
			}
			for i = 0; i < len(result); i+= 2 {
				if max == result[i+1] {
					final = append(final, result[i])
				}
			}
		}
	}
	return final
}

func stdDev(numbers []float64) (result float64) {
	result = 0.0
	median := median(numbers)
	for _, x := range numbers {
		result = result + math.Pow(x - median, 2)
	}
	result = math.Sqrt(result / float64(len(numbers)) )
	return result
}

func main() {
	http.HandleFunc("/", homePage) // goi toi function nay khi duoc requrest 
	if err := http.ListenAndServe(":9001", nil); err != nil { // dung de start server
		log.Fatal("failed to start server", err) // tuong duong voi Printf() va goi toi os.Exit(1)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) { //ResponseWriter de construct mot HTTP response
	pageTop := "<html><body>"
	pageBottom := "</body></html>"
	err := request.ParseForm() // Must be called before writing response phan tinh raw query tu form 
	fmt.Fprint(writer, pageTop, form) // Writer interface la write co ban ghi do dai p bytes tu p dua tren data stream
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
	<tr><th colspan="2">Results</th></tr>
	<tr><td>Numbers</td><td>%v</td></tr>
	<tr><td>Count</td><td>%d</td></tr>
	<tr><td>Mean</td><td>%f</td></tr>
	<tr><td>Median</td><td>%f</td></tr>
	<tr><td>Mode</td><td>%f</td></tr>
	<tr><td>Std. Dev</td><td>%f</td></tr>
	</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.stdDev)
}
