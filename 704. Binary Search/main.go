package main 

// 1 
 
func search(nums []int, target int) int {
    l := len(nums)
    r := 0
    var mid int
    for i := range nums {   
        mid = (l + r) /2
        fmt.Println("mid  : ",mid)
        _ = i
        if target == nums[mid] {
            return mid
        }else if target <  nums[mid]{
            l = mid
        } else if target > nums[mid]{
            r = mid  
        }
    }

    return -1
}

// 2

func search(nums []int, target int) int {
    l := 0
    r := len(nums)-1
    var mid int
    for l <= r {    
        mid = (l + r) /2
          if target == nums[mid] {
            return mid
        }else if target <  nums[mid]{
            r = mid -1
        } else if target > nums[mid]{
            l = mid +1
        }
    }

    return -1
}
