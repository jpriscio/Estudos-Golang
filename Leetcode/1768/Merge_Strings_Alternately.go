package main

func mergeAlternately(word1 string, word2 string) string {

	a := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			a = append(a, word1[i])
			i++
		}

		if j < len(word2) {
			a = append(a, word2[j])
			j++
		}
	}

	return string(a)
}
