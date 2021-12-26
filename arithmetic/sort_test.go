/*
 * @Author: vic
 * @Date: 2021-12-26 22:18:23
 * @LastEditors: vic
 * @LastEditTime: 2021-12-26 22:53:24
 * @Description:
 */
package arithmetic

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		name  string
		nums  []int
		left  int
		right int
	}{
		{"one", []int{2, 1, 3, 4}, 0, 3},
		{"two", []int{2, 1, 1, 3, 3, 4}, 0, 5},
		{"three", []int{}, 0, 0}}

	t.Logf("--------->start")
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.nums, tt.left, tt.right)
			t.Logf("name:%v nums sort:%v", tt.name, tt.nums)
		})
	}
	t.Logf("end<----------")
}
