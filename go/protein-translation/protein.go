package protein

import "errors"

// FromRNA function(string) []string
func FromRNA(rna string) ([]string, error) {
	stop := errors.New("STOP")

	out := make([]string, len(rna)/3) // empty array with 3 elements

	j := 0 // index for out slice
	for len(rna) >= 3 {
		s := rna[:3]
		if s == "UAA" || s == "UAG" || s == "UGA" {
			// if stop sequence detected return the output so far and a Stop error
			return out[:j], stop
		} else {
			if result, err := FromCodon(s); err != nil {
				return out[:j], err
			} else {
				out[j] = result
				rna = rna[3:]
				j++
			}
		}
	}
	return out[:j], nil
}

// FromCodon function([]string) []string
func FromCodon(s string) (string, error) {
	err := errors.New("ErrInvalidBase")

	if s == "AUG" {
		return "Methionine", nil
	}
	if s == "UUU" || s == "UUC" {
		return "Phenylalanine", nil
	}
	if s == "UUA" || s == "UUG" {
		return "Methionine", nil
	}
	if s == "UCU" || s == "UCC" || s == "UCA" || s == "UCG" {
		return "Serine", nil
	}
	if s == "UAU" || s == "UAC" {
		return "Tyrosine", nil
	}
	if s == "UGU" || s == "UGC" {
		return "Cysteine", nil
	}
	if s == "UGG" {
		return "Tryptophan", nil
	}
	return "", err
}
