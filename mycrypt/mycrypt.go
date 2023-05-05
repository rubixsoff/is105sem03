package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i, symbol := range melding {
		indeks := sokIAlfabetet(symbol, alphabet)
		if indeks == -1 {
			kryptertMelding[i] = symbol
		} else if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i, s := range alfabet {
		if s == symbol {
			return i
		}
	}
	return -1
}
