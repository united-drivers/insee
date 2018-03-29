package insee

import "fmt"

func ExampleVAT_Validate() {
	fmt.Println(
		ParseVAT("FR 83 404 833 048").Validate(),
		ParseVAT("FR83404833048").Validate(),
		ParseVAT("FR84404833048").Validate(),
	)
	// Output: false true false
}
