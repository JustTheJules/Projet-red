package src

func IsAlpha(s string) bool {
	b := []byte(s)
	for i := range s {
		if (b[i] >= 97 && b[i] <= 122) || (b[i] >= 65 && b[i] <= 90) {
			i++
		} else {
			return false
		}
	}
	return true
}
