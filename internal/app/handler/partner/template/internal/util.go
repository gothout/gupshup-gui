package util

import "regexp"

// substitui {{1}}, {{2}}, ... por "Variavel1", "Variavel2", ...
func PreencherExemploComVariaveis(texto string) string {
	re := regexp.MustCompile(`{{\s*(\d+)\s*}}`)
	return re.ReplaceAllStringFunc(texto, func(match string) string {
		num := re.FindStringSubmatch(match)[1]
		return "Variavel" + num
	})
}
