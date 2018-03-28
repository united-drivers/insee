package insee

import "fmt"

func ExampleValidateLuhn() {
	fmt.Println(
		ValidateLuhn("49927398716"),
		ValidateLuhn("49927398717"),
		ValidateLuhn("1234567812345678"),
		ValidateLuhn("1234567812345670"),
	)
	// Output: true false false true
}
