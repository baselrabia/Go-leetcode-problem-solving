package main 




// the mot aqurate approch  
func isAnagram(s string, t string) bool {
     if(len(s) != len(t)){
        return false 
    }

   var map1 [26]int
   

    for i := range s {
        map1[int(s[i])-int('a')]++
        map1[int(t[i])-int('a')]--
    }

 
    for _,v := range map1{
        if(v != 0){
            return false 
        }
    }

    return true 
}


// the first approch 


func isAnagram2(s string, t string) bool {
     if(len(s) != len(t)){
        return false 
    }

    map1 := make(map[string]int)

    for _, v := range s {
        map1[string(v)]++
    }



    map2 := make(map[string]int)

    var num1 int
    for _, v := range t {
        num1 =  map1[string(v)]
        if(num1 == 0){
            return false 
        }
         map2[string(v)]++
    }

    if(len(map1) != len(map2)){
        return false 
    }

    for i := range map1{
        if(map1[string(i)] != map2[string(i)]){
            return false 
        }
    }

    return true 
}
