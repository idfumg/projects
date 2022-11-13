package sort

import (
	"sort"
)

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	// time.Sleep(1*time.Second)
}

func StdSort(nums []int) {
	sort.Ints(nums)
}
