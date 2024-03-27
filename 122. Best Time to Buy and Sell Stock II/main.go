package main 


// 1 

func maxProfit(prices []int) int {
  min := prices[0]
  res := 0
  for _,price := range prices {
    if (price < min){
      min = price
    }else {
      res += price - min
      min = price
    }
    
  }
  return res
}



// 2 
func maxProfit(prices []int) int {
  res := 0
  for i,price := range prices {
    if i == 0 {continue}
    
    prev := prices[i-1]
    curr := price
    res += max(curr - prev ,0)
  }
  return res
}



func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
