package strand

import "strings"

func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		translate := map[rune]rune{
			'G': 'C',
			'C': 'G',
			'T': 'A',
			'A': 'U',
		}
		return translate[r]
	}, dna)
}
