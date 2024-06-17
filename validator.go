package main

func Validator(cardNumber string) bool {
	total := 0
	var digit int
	alt := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit = int(cardNumber[i] - '0')

		if alt {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit

		alt = !alt
	}

	return total%10 == 0
}
