package strand

func ToRNA(dna string) string {
	var rnaStrand = map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	str := ""

	for _, s := range dna {
		str += rnaStrand[s]
	}

	return str
}
