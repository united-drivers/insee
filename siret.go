package insee

// INSEE codes
// https://en.wikipedia.org/wiki/INSEE_code
// https://fr.wikipedia.org/wiki/Code_Insee

type Siret struct {
	number string
}

func ParseSiret(input string) *Siret {
	return &Siret{
		number: input,
	}
}

func (s *Siret) Validate() bool {
	return len(s.number) == 14 && ValidateLuhn(s.number)
}

func (s *Siret) Number() string {
	return s.number
}

func (s *Siret) GetSiren() *Siren {
	if !s.Validate() {
		return nil
	}
	return ParseSiren(s.number[:9])
}
