package protein

import (
	"errors"
	"slices"
)

var ErrStop = errors.New("STOP Codons")
var ErrInvalidBase = errors.New("invalid Base")

var ProteinTranslation = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCC", "UCA", "UCG"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
}
var Codons = []string{
	// Methionine
	"AUG",
	// Phenylalanine
	"UUU", "UUC",
	// Leucine
	"UUA", "UUG",
	// Serine
	"UUA", "UUG", "UCU", "UCC",
	// Tyrosine
	"UAU", "UAC",
	// Cysteine
	"UGU", "UGC",
	// Tryptophan
	"UGG",
}
var InvalidCodon = []string{"UAA", "UAG", "UGA"}

func FromRNA(rna string) ([]string, error) {
	output := []string{}
	for start, end := 0, 3; end <= len(rna); {
		r := rna[start:end]
		if slices.Contains(InvalidCodon, r) {
			return output, nil
		}
		switch slices.Index(Codons, r) {
		case 0:
			output = append(output, "Methionine")
		case 1, 2:
			output = append(output, "Phenylalanine")
		case 3, 4:
			output = append(output, "Leucine")
		case 5, 6, 7, 8:
			output = append(output, "Serine")
		case 9, 10:
			output = append(output, "Tyrosine")
		case 11, 12:
			output = append(output, "Cysteine")
		case 13:
			output = append(output, "Tryptophan")
		default:
			return output, ErrInvalidBase
		}
		start += 3
		end += 3
	}

	return output, nil
}

func FromCodon(codon string) (string, error) {
	for k, v := range ProteinTranslation {
		if slices.Contains(v, codon) {
			return k, nil
		}
	}
	if slices.Contains(InvalidCodon, codon) {
		return "", ErrStop
	}
	return "", ErrInvalidBase
}
