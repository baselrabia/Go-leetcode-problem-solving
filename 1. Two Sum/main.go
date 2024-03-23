package main

import "fmt"

//latest approch 

func twoSum(nums []int, target int) []int {

    sumarr := make( map[int]int)

     for i,v := range nums {
        if  val, ok := sumarr[target-v]; ok {
            return []int{val,i}
        }
        sumarr[v] = i
    }

    return []int{}
    
}




// Approach 1: Brute Force

func twoSumForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// Approach 3: One-pass Hash Table


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		comp := target - nums[i]

		_, exist := m[comp]

		if exist {
			return []int{m[comp], i}
		}

		m[nums[i]] = i
	}

	fmt.Println("map:", m)
	return nil

	 
}



func main() {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(r)
}
