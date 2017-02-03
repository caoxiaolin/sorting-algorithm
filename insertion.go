package main

import (
    "fmt"
)

//insertion sorting
func sort(arr []int) {
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i-1] {
            j := i - 1
            temp := arr[i]
            for j >= 0 && arr[j] > temp {
                arr[j+1] = arr[j]
                j--
            }
            arr[j+1] = temp
        }
    }
}

func main() {
    var arr = []int{9, 6, 2, 10, 16, 7, 23, 11}
    sort(arr)
    fmt.Println(arr)
}
