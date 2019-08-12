package main

import "fmt"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

 */




  func main() {

  	nums := []int{1,2,-3,9}
  	fmt.Println(maxSubArray3(nums))

  }


/*
遍历数组，分别找出以每一个元素为开头的最大连续和
 */
func maxSubArray(nums []int) int {

	max := 0
	negitive := -1<<63
	for index, num := range nums {
		if num <= 0 {
			if num > negitive {
				negitive = num
			}
			continue
		}
		sum := num
		maxVal := num
		for i := index+1; i<len(nums); i++ {
			sum += nums[i]
			if sum > maxVal {
				maxVal = sum
			}
		}

		if maxVal > max {
			max = maxVal
		}
	}
	if max == 0 {
		return negitive
	}
	return max
}

/*
分治法：
将nums[]分为左右两部分
那么最大的和要么出现在最左边，要么出现在最右边，要么出现在中间
 */

 func maxSubArray2(nums []int) int {

 	if len(nums) == 1 {
 		return nums[0]
	}
	maxLeft := maxSubArray(nums[:len(nums)/2])
	maxRight := maxSubArray(nums[len(nums)/2:])

	// 中间交叉的最大值
	// 分别求出以nums[len(nums)/2-1]结尾的左边的最大和，nums[len(nums)/2]开头的右边的最大和
	mLeftMax := nums[len(nums)/2-1]
	sum := mLeftMax
	for i:= len(nums)/2-2; i>=0; i-- {
		sum+=nums[i]
		if sum > mLeftMax {
			mLeftMax = sum
		}
	}

	mRightMax := nums[len(nums)/2]
	sum = mRightMax
	for i:= len(nums)/2+1; i< len(nums);i++ {
		sum += nums[i]
		if sum > mRightMax {
			mRightMax = sum
		}
	}

	mid := mLeftMax + mRightMax

	res := maxLeft
	if res < maxRight {
		res = maxRight
	}
	if res < mid {
		res = mid
	}
	return res
 }




func maxSubArray3(nums []int) int {

	max := nums[0]
	sum := max

	for i := 1; i< len(nums); i++ {
		if sum >=0 {
			sum += nums[i]
			if sum > max {
				max = sum
			}
		} else {
			if nums[i] > sum && nums[i] > max {
				max = nums[i]
			} else if sum > max {
				max = sum
			}
			sum = nums[i]
		}
	}
	return max
}