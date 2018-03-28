package insee

// Luhn algorithm
// http://en.wikipedia.org/wiki/Luhn_algorithm
// http://rosettacode.org/wiki/Luhn_test_of_credit_card_numbers#Go

var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

func ValidateLuhn(s string) bool {
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}
