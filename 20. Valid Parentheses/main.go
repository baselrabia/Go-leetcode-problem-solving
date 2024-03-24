package main 

// first approach 
func isValid(s string) bool {

    pairs := map[byte]byte{
        '}' : '{',
        ']' : '[',
        ')' : '(',
    }
 
	 stack := make([]byte,0)

     for _,v := range []byte(s) {
        pair, ok := pairs[v]
        if !ok {
            stack = append(stack, v)
            continue;
        }

        if (len(stack) == 0){
            return false
        }

        if(stack[len(stack)-1] != pair){
            return false
        } 
        
        stack = stack[:len(stack)-1]

     }

     return len(stack) == 0
}


// second approach 

type Stack []byte

func (s *Stack) append(data byte) {
    *s = append(*s, data)
}

func (s *Stack) pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) last() byte {
    return (*s)[len(*s)-1]
}

func isValid(s string) bool {

    pairs := map[byte]byte{
        '}' : '{',
        ']' : '[',
        ')' : '(',
    }
 
	stack := Stack{}

    for _,v := range []byte(s){
        pair, ok := pairs[v]
        if !ok{
            stack.append(v)
            continue;
        } 

        if len(stack) == 0 {
            return false
        }

        if stack.last() != pair{
            return false
        }

        stack.pop()
    }

    return len(stack) == 0 
}
