package hackerrank

import (
    "sort"
)

func MinimumSwaps(arr []int32) int32 {
    newArray := make([]int, len(arr))
    for index, num := range(arr) {
        newArray[index] = int(num)
    }
    
    unsorted := map[int]int{}
    indexToNum := map[int]int{}
    
    for index, num := range(newArray) {
        unsorted[num] = index
        indexToNum[index] = num
    }
    
    sorted := map[int]int{}
    sort.Ints(newArray)
    for index, num := range(newArray) {
        sorted[num] = index
    }
    
    var swaps int32
    for num, index := range(unsorted) {
        desiredIndex := sorted[num]
        if index != desiredIndex {
            swapNum := indexToNum[desiredIndex]
            
            temp := unsorted[num]
            unsorted[num] = desiredIndex
            unsorted[swapNum] = temp
            
            indexToNum[desiredIndex] = num
            indexToNum[temp] = swapNum
            swaps++
        }    
    }
    return swaps
}

