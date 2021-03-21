package main

import "fmt"
import "time"

func main() {
    inElements := []int{67, 23, 11, 9075, 34, 26, 88, 97, 124, 9845, 5409, 36, 9884, 45879, 2144, 250, 760}

    start := time.Now()
    fmt.Println("Bubble sort")
    fmt.Println(Bubble(inElements))
    duration := time.Since(start)
    fmt.Println("Time lapsed: ", duration)
    
    start = time.Now()
    fmt.Println("Merge sort")
    fmt.Println(Merge(inElements))
    duration = time.Since(start)
    fmt.Println("Time lapsed: ", duration)
}

func Bubble(elements []int) []int {
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

func Merge(elements []int)  []int {
    if (len(elements) < 2) {
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


