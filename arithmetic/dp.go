/*
 * @Author: vic
 * @Date: 2021-12-26 18:40:59
 * @LastEditTime: 2021-12-26 21:28:29
 * @LastEditors: Please set LastEditors
 * @Description: 简单的算法
 * @FilePath: /src/demo/arithmetic/dp.go
 */
package arithmetic

/**
 * @description:
 * @param {[]int} w 重量
 * @param {*} v 价值
 * @param {int} 最大重量
 * @return {*} 最大价值
 */
func BasePackage(w []int, v []int, target int) int {

	//物品数
	n := len(w)

	//重量 Math.max(int,int),太简单，所以要用户自己实现，笑死我了
	dp := make([]int, target+1)

	for i := 0; i < n; i++ {
		for j := target; j > 0; j-- {
			if j-w[i] >= 0 {
				//go没有三元运算符
				if (dp[j-1] - (dp[j-w[i]] + v[i])) > 0 {
					dp[j] = dp[j-1]
				} else {
					dp[j] = dp[j-w[i]] + v[i]
				}

			} else {
				dp[j] = dp[j-1]
			}
		}
	}

	return dp[target]
}
