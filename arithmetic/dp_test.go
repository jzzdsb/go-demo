/*
 * @Author: vic
 * @Date: 2021-12-26 19:59:34
 * @LastEditTime: 2021-12-26 21:09:03
 * @LastEditors: Please set LastEditors
 * @Description:测试背包算法
 * @FilePath: /src/demo/arithmetic/dp_test.go
 */
package arithmetic

import (
	"testing"
)

func TestBasePackage(t *testing.T) {
	//定义一个内置对象
	type dpParam struct {
		w      []int
		v      []int
		target int
	}

	//定义参数
	good := dpParam{
		w:      []int{2, 1, 3},
		v:      []int{4, 2, 3},
		target: 4,
	}

	tests := []struct {
		name   string
		args   dpParam
		answer int
	}{
		{
			name:   "good",
			args:   good,
			answer: 6,
		},
	}

	t.Logf("------------>begin")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := BasePackage(tt.args.w, tt.args.v, tt.args.target); result != tt.answer {
				t.Errorf("basePackage() result = %d answer = %d", result, tt.answer)
			}
		})
	}
	t.Logf("end<--------------")

}
