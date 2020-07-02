package main
import "fmt"

func main() {
    var n int
    if m, err := Scan(&n); m != 1 {
        panic(err)
    }

	elements := make([]int, n)
	readCases(elements, 0, n)
	printCases(elements, 0, n)
}

func printCases(elements []int, i, n int) {
	if n == 0 {
		return
	}

    fmt.Println(elements[i])
    printCases(elements, i+1, n-1)
}

func readCases(elements []int, i, n int) {
	if n == 0 {
		return
	}
	var m int
	Scan(&m)
	ReadN(elements, i, 0, m)
    readCases(elements, i+1, n-1)
}


func ReadN(all []int, c, i, n int) {
    if n == 0 {
        return
    }

    var x int
    if m, err := Scan(&x); m != 1 {
        panic(err)
    }

    if x > 0 {
    	all[c] += x*x
	}

    ReadN(all, c, i+1, n-1)
}

func Scan(a *int) (int, error) {
    return fmt.Scan(a)
}
