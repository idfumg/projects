package main

import "fmt"

func binarySearchRec(arr []int, key int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] == key {
		return mid
	}
	if arr[mid] < key {
		return binarySearchRec(arr, key, mid+1, high)
	}
	return binarySearchRec(arr, key, low, mid-1)
}

func binarySearchIter(arr []int, key int, low int, high int) int {
	for low != high {
		mid := low + (high-low)/2
		if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if arr[low] == key {
		return low
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The position of 8 is:", binarySearchRec(arr, 8, 0, len(arr)-1))
	fmt.Println("The position of 8 is:", binarySearchIter(arr, 8, 0, len(arr)-1))
}
