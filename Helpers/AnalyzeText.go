package helpers

import (
	"regexp"
	"strings"
)

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false

}

func CheckRegexp(String string) bool {
	return regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})`).MatchString(String)
}

func AnalyzeText(prompt string) string {

	options := map[string][]string{
		"Aquí estan tus citas agendadas":          []string{"MIS", "CITAS", "VER", "AGENDADAS", "AGENDE", "MOSTRAR", "MUESTRAME"},
		"Para que fecha quieres agendar la cita?": []string{"AGENDAR", "CITA", "APARTAR", "CREAR", "RESERVAR", "HACER"},
		"Que cita quieres reagendar?":             []string{"CAMBIAR", "REAGENDAR", "NUEVA", "FECHA", "CAMBIO", "MODIFICAR", "EDITAR"},
		"Aquí esta el calendario del doctor":      []string{"CALENDARIO", "AGENDA", "DOCTOR", "DISPONIBLE"},
	}

	palabras := strings.Split(strings.ToUpper(prompt), " ")

	counter := make(map[string]int)

	for key, value := range options {
		counter[key] = 0
		for _, parts := range palabras {

			if contains(value, parts) {
				counter[key]++
			}
		}
	}
 	return findMaxValue(counter)

}
