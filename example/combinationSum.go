package example

/**
 * 39. 组合总和
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * candidates 中的数字可以无限制重复被选取。
 * 搜索回溯?=
 */
func combinationSum(candidates []int, target int) [][]int {
	comb := []int{}
	ans := [][]int{}
	var dfs func(target, i int)
	dfs = func(target, i int) {
		if len(candidates) == i {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		//直接跳过
		dfs(target, i+1)
		//// 选择当前数
		if target-candidates[i] >= 0 {
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
