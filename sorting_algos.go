package main

import (
	"fmt"
	"math/rand"
)
import "time"

func main() {
	// Testing sorting with 100.000 numbers shuffled randomly
	a := MakeRange(0, 15)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	start := time.Now()
	fmt.Println("Bubble sort")
	fmt.Printf("%v\n", BubbleSort(a))
	duration := time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	start = time.Now()
	fmt.Println("Merge sort")
	fmt.Printf("%v\n", MergeSort(a))
	duration = time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	start = time.Now()
	fmt.Println("Quick sort")
	fmt.Printf("%v\n", QuickSort(a))
	duration = time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	start = time.Now()
	fmt.Println("Selection sort")
	fmt.Printf("%v\n", SelectionSort(a))
	duration = time.Since(start)
	fmt.Println("Time lapsed: ", duration)
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func BubbleSort(elements []int) []int {
	for i := len(elements); i > 0; i-- {
		for j := 1; j < i; j++ {
			if elements[j-1] > elements[j] {
				innerElement := elements[j]
				elements[j] = elements[j-1]
				elements[j-1] = innerElement
			}
		}
	}
	return elements
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

func QuickSort(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}

	left, right := 0, len(elements) - 1

	pivot := rand.Int() % len(elements)
	elements[pivot], elements[right] = elements[right], elements[pivot]

	for i := range elements {
		if elements[i] < elements[right] {
			elements[i], elements[left] = elements[left], elements[i]
			left++
		}
	}

	elements[left], elements[right] = elements[right], elements[left]

	QuickSort(elements[:left])
	QuickSort(elements[left + 1:])

	return elements
}


func SelectionSort(elements []int) []int {
	for i := 0; i < len(elements); i++ {
		for j:= i + 1; j < len(elements); j++ {
			for elements[j] < elements[i] {
				elements[i], elements[j] = elements[j], elements[i]
			}
		}
	}
	return elements
}
