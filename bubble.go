package main

import (
    "fmt"
)

//bubble sorting
func sort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := 1; j < len(arr)-i; j++ {
            if arr[j] < arr[j-1] {
                arr[j], arr[j-1] = arr[j-1], arr[j]
            }
        }
    }
}

func main() {
    var arr = []int{9, 6, 2, 10, 16, 7, 23, 11}
    sort(arr)
    fmt.Println(arr)
}
