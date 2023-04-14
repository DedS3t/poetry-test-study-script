package pkg

import "math"

func CosineSimilarity(s1 string, s2 string) float64 {
	freq1 := getCharFrequency(s1)
	freq2 := getCharFrequency(s2)

	return cosineSimilarity(freq1, freq2)
}

func getCharFrequency(str string) map[rune]int {
	// count character frequency
	freq := make(map[rune]int)
	for _, char := range str {
		freq[char]++
	}

	return freq
}

func cosineSimilarity(freq1, freq2 map[rune]int) float64 {
	// create a set of all characters in the two frequency vectors
	chars := make(map[rune]bool)
	for char := range freq1 {
		chars[char] = true
	}
	for char := range freq2 {
		chars[char] = true
	}

	// calculate dot product and magnitude of vectors
	dotProduct := 0.0
	mag1 := 0.0
	mag2 := 0.0
	for char := range chars {
		dotProduct += float64(freq1[char] * freq2[char])
		mag1 += math.Pow(float64(freq1[char]), 2)
		mag2 += math.Pow(float64(freq2[char]), 2)
	}

	// calculate cosine similarity
	if mag1 == 0 || mag2 == 0 {
		return 0
	}
	return dotProduct / (math.Sqrt(mag1) * math.Sqrt(mag2))
}
