/*
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.
*/
package runningsumof1darray

func RunningSum(nums []int) []int {
	var res []int
	var acc int
	for _, num := range nums {
		acc += num
		res = append(res, acc)
	}
	return res
}
