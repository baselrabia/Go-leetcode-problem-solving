func reversePrefix(word string, ch byte) string {
    wordbyt := []byte(word)
    var rev []byte
    for i, v := range wordbyt {
        rev = append([]byte{v}, rev...)
        if (ch == v){
            return string(rev) + string(wordbyt[i+1:])
        }

    }

    return word
}