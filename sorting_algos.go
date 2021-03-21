package main

import (
	"fmt"
	"math/rand"
)
import "time"

func main() {
	// Testing sorting with 100.000 numbers shuffled randomly
	a := MakeRange(0, 100000)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	start := time.Now()
	fmt.Println("Bubble sort")
	BubbleSort(a)
	duration := time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	start = time.Now()
	fmt.Println("Merge sort")
	MergeSort(a)
	duration = time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	start = time.Now()
	fmt.Println("Quick sort")
	QuickSort(a)
	duration = time.Since(start)
	fmt.Println("Time lapsed: ", duration)

	start = time.Now()
	fmt.Println("Selection sort")
	SelectionSort(a)
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

func MergeSort(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}

	mid := len(elements) / 2
	left := elements[:mid]
	right := elements[mid:]

	size, i, j := len(left)+len(right), 0, 0

	slice := make([]int, size, size)

	for el := 0; el < size; el++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[el] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[el] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[el] = left[i]
			i++
		} else {
			slice[el] = right[j]
			j++
		}
	}
	return slice
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
