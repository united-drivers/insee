package insee

import "strconv"

// INSEE codes
// https://en.wikipedia.org/wiki/INSEE_code
// https://fr.wikipedia.org/wiki/Code_Insee

type Siren struct {
	number string
}

func ParseSiren(input string) *Siren {
	return &Siren{
		number: input,
	}
}

func (s *Siren) Validate() bool {
	return len(s.number) == 9 && ValidateLuhn(s.number)
}

func (s *Siren) Number() string {
	return s.number
}

func (s *Siren) VATKey() string {
	number, err := strconv.Atoi(s.number)
	if err != nil {
		return ""
	}
	return strconv.Itoa((12 + 3*number%97) % 97)
}

func (s *Siren) VAT(countryPrefix string) *VAT {
	return ParseVAT(countryPrefix + s.VATKey() + s.number)
}
