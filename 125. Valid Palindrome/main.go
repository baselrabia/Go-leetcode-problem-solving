package main 

// two pointer 

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	arr := []rune(s)

	for i < j {
		left := unicode.ToLower(arr[i])
		right := unicode.ToLower(arr[j])

		if !isLetterOrDigit(left) {
			i++
			continue
		}

		if !isLetterOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s) || unicode.IsDigit(s)
}


// two pointer different approach  

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

