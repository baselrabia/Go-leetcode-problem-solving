package main 

// 1

func subsets(nums []int) [][]int {
    var res [][]int 
    var subset []int 
    var dfs func (i int)
    dfs = func (i int) {
        //  condition to stop the recursion
        if (i >= len(nums)){
            res = append(res, append([]int{}, subset...))
            return
        }

        // add nums i 
        subset = append(subset,nums[i])
        dfs(i+1)


        // dont add 
        subset = subset[:len(subset)-1]
        dfs(i+1)

    }

    dfs(0)
    return res
}

 
