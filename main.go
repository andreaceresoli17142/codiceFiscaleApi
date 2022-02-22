package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// commento pazzo

var vocali = map[rune]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

var monthToLetter = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "H", 7: "L", 8: "M", 9: "P", 10: "R", 11: "S", 12: "T"}

var CARATTERE_CONTROLLO_PARI = map[string]int{
	"0": 1,
	"1": 0,
	"2": 5,
	"3": 7,
	"4": 9,
	"5": 13,
	"6": 15,
	"7": 17,
	"8": 19,
	"9": 21,
	"I": 19,
	"R": 8,
	"A": 1,
	"J": 21,
	"S": 12,
	"B": 0,
	"K": 2,
	"T": 14,
	"C": 5,
	"L": 4,
	"U": 16,
	"D": 7,
	"M": 18,
	"V": 10,
	"E": 9,
	"N": 20,
	"W": 22,
	"F": 13,
	"O": 11,
	"X": 25,
	"G": 15,
	"P": 3,
	"Y": 24,
	"H": 17,
	"Q": 6,
	"Z": 23,
}

var CARATTERE_CONTROLLO_DISPARI = map[string]int{
	"0": 0,
	"9": 9,
	"I": 8,
	"R": 17,
	"1": 1,
	"A": 0,
	"J": 9,
	"S": 18,
	"2": 2,
	"B": 1,
	"K": 10,
	"T": 19,
	"3": 3,
	"C": 2,
	"L": 11,
	"U": 20,
	"4": 4,
	"D": 3,
	"M": 12,
	"V": 21,
	"5": 5,
	"E": 4,
	"N": 13,
	"W": 22,
	"6": 6,
	"F": 5,
	"O": 14,
	"X": 23,
	"7": 7,
	"G": 6,
	"P": 15,
	"Y": 24,
	"8": 8,
	"H": 7,
	"Q": 16,
	"Z": 25,
}

var CARATTERE_CONTROLLO_RESTO = []string{
	"A",
	"H",
	"O",
	"V",
	"B",
	"I",
	"P",
	"W",
	"C",
	"J",
	"Q",
	"X",
	"D",
	"K",
	"R",
	"Y",
	"E",
	"L",
	"S",
	"Z",
	"F",
	"M",
	"T",
	"G",
	"N",
	"U",
}

func readCsvFromFile(file_path string) ([][]string, error) {

	ret := [][]string{}

	file, err := os.Open(file_path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowContent := strings.Split(scanner.Text(), ",")
		ret = append(ret, rowContent)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func isCons(char rune) bool {

	_, ok := vocali[char]
	return !ok
}

func calcolaPrimeTreCons(stringa string) string {

	ret := ""

	for _, crune := range stringa {
		if isCons(crune) {
			ret += string(crune)
		}
	}
	return ret
}

func calcolaCodiceFiscale(nome string, cognome string, luogo_di_nascita string, sigla_provincia string, sesso string, data_di_nascita string) string {

	ret := ""

	sessoSanitized := strings.ToUpper(sesso)

	// splitto la data di nascita salvata dd/mm/yyyy
	dataNascitaSplitted := strings.Split(data_di_nascita, "/")

	// prime tre lettere del cognome
	consCognome := calcolaPrimeTreCons(strings.ToUpper(cognome))

	for len(consCognome) < 3 {
		consCognome += "X"
	}

	ret += consCognome

	// prime tre lettere del nome
	consNome := calcolaPrimeTreCons(strings.ToUpper(nome))

	for len(consNome) < 3 {
		consNome += "X"
	}

	ret += consNome

	// ultime due cifre dell' anno di nascita
	ret += dataNascitaSplitted[2][2:4]

	//mese dell' anno trasformata in lettera
	intMonth, _ := strconv.Atoi(dataNascitaSplitted[1])

	ret += monthToLetter[intMonth]

	// giorno di nascita (+40 se donna)
	giornoNascita, _ := strconv.Atoi(dataNascitaSplitted[0])

	if sessoSanitized == "F" {
		giornoNascita += 40
	}

	ret += strconv.Itoa(giornoNascita)

	// placeholder codice catastale
	ret += "X000"

	// calcolo carattere di controllo
	resto := 0
	for i, char := range ret {
		if i%2 == 0 {
			resto += CARATTERE_CONTROLLO_PARI[string(char)]
		} else {
			resto += CARATTERE_CONTROLLO_DISPARI[string(char)]
		}
	}

	ret += CARATTERE_CONTROLLO_RESTO[resto%26]

	return ret
}

func main() {

	csvContent, err := readCsvFromFile("testfile.csv")

	if err != nil {
		fmt.Print(err)
		return
	}

	// fmt.Printf("slice content: %v\n", csvContent)

	for _, rowSlice := range csvContent {
		fmt.Println(calcolaCodiceFiscale(rowSlice[0], rowSlice[1], rowSlice[2], rowSlice[3], rowSlice[4], rowSlice[5]))
	}
}
