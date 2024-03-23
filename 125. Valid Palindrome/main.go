package main 



func isPalindrome(s string) bool {
 	s = strings.ToLower(s)

    var new string
    for _,r := range s {
        char := string(r)
        if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
            new+=char
        }
    }
    if len(new) == 1{
        return true
    }

    itr := (len(new)/2)
    for i := range itr {
        j := len(new) - i -1
        if (new[i] != new[j]){
            return false 
        }
    }



    fmt.Println("line exec: ",new)

    return true

}

