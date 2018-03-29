package insee

// INSEE codes
// https://en.wikipedia.org/wiki/INSEE_code
// https://fr.wikipedia.org/wiki/Code_Insee

type VAT struct {
	number string
}

func ParseVAT(input string) *VAT {
	return &VAT{
		number: input,
	}
}

func (s *VAT) Validate() bool {
	return len(s.number) == 13 && ValidateLuhn(s.GetSiren().Number()) && s.GetSiren().VATKey() == s.number[2:4]
}

func (s *VAT) Number() string {
	return s.number
}

func (s *VAT) GetSiren() *Siren {
	return ParseSiren(s.number[4:])
}
