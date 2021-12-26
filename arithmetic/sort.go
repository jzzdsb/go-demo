/*
 * @Author: vic
 * @Date: 2021-12-26 21:38:30
 * @LastEditors: vic
 * @LastEditTime: 2021-12-26 22:37:59
 * @Description:
 */
package arithmetic

/**
 * @description: 快排
 * @param {[]int} nums 排列数组
 * @return {*}
 */
func QuickSort(nums []int, left int, right int) {

	if left >= right {
		return
	}

	index, d, g := left, left, right

	for d < g {
		for d < g && nums[g] >= nums[index] {
			g--
		}

		for d < g && nums[d] <= nums[index] {
			d++
		}

		d, g = g, d
	}

	nums[index], nums[d] = nums[d], nums[index]
	index = d
	QuickSort(nums, left, index)
	QuickSort(nums, index+1, right)
}
