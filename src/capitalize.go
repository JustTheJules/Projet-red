package src

func Capitalize(s string) string {
	b := []rune(s)
	maj := 0
	for ind, let := range s {
		if let >= 'A' && let <= 'Z' {
			b[ind] = let + 32
		}
	}
	for i, l := range b {
		if (l >= 'a' && l <= 'z') && maj == 0 {
			b[i] = l - 32
			maj++
		}
		if maj == 0 && b[i] >= '0' && b[i] <= '9' {
			maj = 1
		}
		if (b[i] >= 9 && b[i] <= 47) || (b[i] >= 58 && b[i] <= 64) || (b[i] >= 91 && b[i] <= 96) || (b[i] >= 123 && b[i] <= 126) {
			maj = 0
		}
	}
	return string(b)
}
