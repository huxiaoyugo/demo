package main

import "fmt"

func main()  {

	fmt.Println(searchInsert([]int{1,3}, 4))
}

func searchInsert(nums []int, target int) int {

	return search(nums,0,len(nums)-1, target)
}


func search(nums []int, start,end,target int) int {


	mid := (start+end)/2
	midVal := nums[mid]
	if target == midVal {
		return mid
	}

	if start == end {
		if midVal > target {
			return start
		} else {
			return start+1
		}
	}

	if start + 1 == end {
		if target < nums[start] {
			return start
		}
		if target <= nums[end] {
			return end
		}
		if target > nums[end] {
			return end+1
		}
	}

	if midVal > target {
		return search(nums, start, mid-1, target)
	} else {
		return search(nums, mid+1, end, target)
	}

}