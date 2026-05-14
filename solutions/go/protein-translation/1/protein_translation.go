package proteintranslation

import "errors"

// Global sentinels required for test suite compliance.
var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

func FromRNA(rna string) ([]string, error) {
	var result []string

	// Process strand in 3-byte blocks to match codon structure.
	for i := 0; i < len(rna); i += 3 {
		end := i + 3

		// Bounds check: verify if a full 3-character window exists.
		// Prevents runtime panic on incomplete trailing sequences.
		if end > len(rna) {
			return nil, ErrInvalidBase
		}

		chunk := rna[i:end]
		protein, err := FromCodon(chunk)

		// Priority 1: Exit immediately on STOP sentinel.
		// Validates cases where a STOP appears before junk/incomplete data.
		if err == ErrStop {
			return result, nil
		}

		// Priority 2: Standard error propagation for malformed codons.
		if err != nil {
			return nil, err
		}

		// Commit valid translation to the results slice.
		result = append(result, protein)
	}

	return result, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		// Trigger the ErrStop signal to terminate the FromRNA loop.
		return "", ErrStop
	default:
		// Default case handles unexpected character sequences.
		return "", ErrInvalidBase
	}
}