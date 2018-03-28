package insee

import "fmt"

func ExampleSiret_Validate() {
	fmt.Println(
		ParseSiret("").Validate(),
		ParseSiret("404 833 048 00022").Validate(),
		ParseSiret("40483304800022").Validate(),
		ParseSiret("404 833 048 00023").Validate(),
		ParseSiret("40483304900023").Validate(),
		ParseSiret("1").Validate(),
		ParseSiret("1111111111").Validate(),
	)
	// Output: false false true false false false false
}

func ExampleSiret_GetSiren() {
	fmt.Println(
		ParseSiret("").GetSiren(),
		ParseSiret("404 833 048 00022").GetSiren(),
		ParseSiret("40483304800022").GetSiren().Number(),
		ParseSiret("404 833 048 00023").GetSiren(),
		ParseSiret("40483304900023").GetSiren(),
		ParseSiret("1").GetSiren(),
		ParseSiret("1111111111").GetSiren(),
	)
	// Output: <nil> <nil> 404833048 <nil> <nil> <nil> <nil>
}
