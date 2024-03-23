package main


func containsDuplicate(nums []int) bool {

    dub := make([]int ,100000)
 
    for _,i := range nums {
        if (dub[i] != 0){
            return true
        }
        dub[i]++;
    }
    return false;

}
