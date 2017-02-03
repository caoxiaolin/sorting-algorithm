package main

import (
    "fmt"
)

//quick sorting
func sort(arr []int) {
    if len(arr) <= 1 {
        return
    }
    mid, i := arr[0], 1
    head, tail := 0, len(arr)-1
    for head < tail {
        if arr[i] > mid {
            arr[i], arr[tail] = arr[tail], arr[i]
            tail--
        } else {
            arr[i], arr[head] = arr[head], arr[i]
            head++
            i++
        }
    }
    arr[head] = mid
    sort(arr[:head])
    sort(arr[head+1:])
}

func main() {
    var arr = []int{9, 6, 2, 10, 16, 7, 23, 11}
    sort(arr)
    fmt.Println(arr)
}
