package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	tmpArr := make([]int, len(nums1)+len(nums2))
	ti,ai,bi := 0,0,0
	for {
		if ai >= len(nums1) {
			for ;bi<len(nums2);bi++{
				tmpArr[ti]=nums2[bi]
				ti++
			}
			break
		}
		if bi >= len(nums2) {
			for;ai<len(nums1);ai++ {
				tmpArr[ti]=nums1[ai]
				ti++
			}
			break
		}

		if nums1[ai] <= nums2[bi] {
			tmpArr[ti] = nums1[ai]
			ai++
		} else {
			tmpArr[ti] = nums2[bi]
			bi++
		}
		ti++
	}
	nums1=nil
	nums2=nil
	count := len(tmpArr)

	if count == 0 {
		return 0
	}

	if count == 1 {
		return float64(tmpArr[0])
	}

	if count%2==0 {
		return float64(tmpArr[count/2] + tmpArr[count/2-1])/2
	}
	return float64(tmpArr[count/2])
}