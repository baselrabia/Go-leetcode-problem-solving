package main 
// 1

func maxProfit(prices []int) int {
	min := math.MaxUint32
    res := 0
    fmt.Println(min)
    for _, price := range prices{
        fmt.Println("p m r", price, " ", min, " ", res)
        if price < min {
               min = price
        } else { 
             if (price - min > res){
                res = price - min
            }
        }

    }
     
     return res
}



// 2


func maxProfit(prices []int) int {

    max := 0
    var s, pf, ps int
    for f := range len(prices)-1{
        pf = prices[f] 
        s = f+1
        for s < len(prices){
            // fmt.Println("f s : ", f , " ", s )
            ps = prices[s] 
            if (ps - pf > max) {
                max = ps - pf
            }
            s++
        }

    }

    return max;
}
