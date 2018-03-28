package insee

import "fmt"

func ExampleSiren_Validate() {
	fmt.Println(
		ParseSiren("").Validate(),
		ParseSiren("404 833 048").Validate(),
		ParseSiren("404833048").Validate(),
		ParseSiren("404 833 049").Validate(),
		ParseSiren("404833049").Validate(),
		ParseSiren("1").Validate(),
		ParseSiren("1111111111").Validate(),
	)
	// Output: false false true false false false false
}
