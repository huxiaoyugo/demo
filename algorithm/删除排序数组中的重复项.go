package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	var p,k = 1,1
	for ;p<len(nums);p++ {
		if nums[p] == nums[k-1] {
			continue
		}
		if k !=p {
			nums[k] = nums[p]
			k++
		}
	}
	return k
}


func removeElement(nums []int, val int) int {
	var p,k = 0,0
	for ;p<len(nums); p++ {
		if nums[p] == val {
			continue
		}
		if k !=p {
			nums[k] = nums[p]
		}
		k++
	}

	return k
}

func main() {

	nums := []int{1,2,3,4}

	fmt.Println(removeElement(nums,5))

	fmt.Println(nums)
}