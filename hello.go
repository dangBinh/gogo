package main

import "fmt"
import "math"
import "time"

const s string = "costant"

func main() {
	// Value
	fmt.Println("go" + "lang");
	fmt.Println("1+1 = ", 1+1);
	fmt.Println("7.0 / 3.0 = ", 7.0 / 3.0);

	fmt.Println(true && false);
	fmt.Println(true || false);
	fmt.Println(!true);

	// Variable
	var f string = "string" 
	var a1,b1 int = 1, 2
	d:= "string" // short hand d string = "string"
 	var e int // mac dinh la 0 
 	fmt.Println(e);
	fmt.Println(f, a1, b1, d)

	// Constant 
	const n = 5000000
	fmt.Println(s);
	fmt.Println(math.Sin(n));

	// For 
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} 
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}

	// If Else
	if 1 == 1 {
		fmt.Println("True")
	} else if 2 == 1 {
		fmt.Println("false")
	} else {
		fmt.Println("yeah")
	}

	// Switch 
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Before noon")
		default: 
			fmt.Println("After noon")
	}

	// Array 
	var a [5]int 
	a[4] = 100
	b := [5]int{1 ,2 , 3, 4, 5}
	fmt.Println(b);
	var twoD [4][5]int
	fmt.Println(twoD);

	// Slices 
	// Slice giong voi array nhung them nhieu chuc nang hon
	// Slice duoc dinh kieu boi data type no chua 
	// Tao slice voi make 
	s := make([]string, 3)
	fmt.Println(s)

	s = append(s, "a")
	fmt.Println(s)

	fmt.Println(len(s)) // cha ve chieu dai cua s
	var sliceBetween = s[2:5]
	fmt.Println(sliceBetween)
	var sliceUpTo = s[:5]
	fmt.Println(sliceUpTo)
	var sliceUpFrom = s[2:]
	fmt.Println(sliceUpFrom)
	
	twoD2 := make([][]int, 3)
	fmt.Println(twoD2)

	// Map 
	// associative data type 
	sMap := make(map([string]int))
	sMap["k1"] = 7
	sMap["k2"] = 6
	fmt.Println(sMap)

	delete(sMap, "k2")
	_, prs := sMap["k2"]; // loai bo truong hop thieu key
	fmt.Println(sMap);

	// Range
	// Lap qua cac phan tu
	myNums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += sum
	}
	fmt.Println(sum)

	//Closure giong js
}
func plus(a, b, c int){
	return a + b + c
} 

func vals(int, int) {
	return a, b // a, b := vals(1, 2)
}

// goi voi nhieu doi so 
func sum(nums ...int) { // su dung slice func(slice ...)
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}	