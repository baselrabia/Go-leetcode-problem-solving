func reversePrefix(word string, ch byte) string {
    wordbyt := []byte(word)
    var rev string
    for i, v := range wordbyt {
        rev =  string(v) + rev

        if (ch == v){
            return  rev + word[i+1:]
        }

    }

    return word
}