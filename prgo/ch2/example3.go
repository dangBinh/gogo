package main 

import (
	"fmt"
	// "sort"
	"net/http"
	"log"
	// "strings"
	// "strconv"
	"math"
	"math/cmplx"
)

const form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="arg[]" size="30"> x^2 + 
<input type="text" name="arg[]" size="30"> x + 
<input type="text" name="arg[]" size="30">  ->
<input type="submit" value="Calculate">
</form>`;

const anError = `<p class="error">%s</p>`

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
		if arg, message, ok := processRequest(request); ok{
			x1, x2 := solve(arg)
			fmt.Println(x1, x2)
			// fmt.Fprint(writer, formatQuestion(question))
			// fmt.Fprint(writer, formatSolution(result))
		} else {
			// fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var x []float64
	req := request.Form["arg[]"]
	for _, x := range req {
		if x, err := strconv.ParseFloat(x); if err != nil {
			return x, x + "Not invalid", false
		}
	}	
	fmt.Println(req)
	x = append(x, req)
	return x, "", true
}

func formatQuestion() {

}

func solve(question []float64) (complex128, complex128){
	if len(question) > 0 {
		a := complex(question[1], 0)
		b := complex(question[2], 0)
		c := complex(question[3], 0)
		delta := cmplx.Sqrt(cmplx.Pow(b, 2) - 4 * a * c)
			x1 := ( -b + cmplx.Sqrt(delta) ) / (2 * a)
			x2 := ( -b - cmplx.Sqrt(delta) ) / (2 * a)
			return x1 ,x2 
	}
}

func formatSolution() {
	
}

// EqualFloat() tra ve gia tri dung x xap xi gan bang y voi limit cho truoc 
func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x - y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func EqualComplex(x, y) {

}