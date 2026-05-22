func isPalindrome(s string) bool {
    low := strings.ToLower(s)
    clean := []byte{}
    for _, b := range []byte(low) {
        if ('a' <= b && b <= 'z') || (b >= '0' && b <= '9') {
            clean = append(clean, b)
        }
    }

    for i := 0; i < len(clean)/2; i++ {
        if clean[i] != clean[len(clean)-i-1] {
            return false
        }
    }
    return true
}
