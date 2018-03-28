package insee

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
