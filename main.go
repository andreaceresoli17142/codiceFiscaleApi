package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var vocali = map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func input(input_text string) string {
	var inp string
	fmt.Print(input_text)
	fmt.Scanf("%v", &inp)
	return inp
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

	_, ok := a[char]
	return ok
}

func calcolaPrimeTreCons(stringa string) string {

	ret := ""

	for _, crune := range stringa {
		if isCons(crune) {
			ret += crune
		}
	}

}

func calcolaCodiceFiscale(nome string, cognome string) string {

	//for len(nome) < 3 {
	//	 nome += "X"
	//}

	//return nome + " " + cognome
	return calcolaPrimeTreCons(nome)
}

func main() {
	//fmt.Println("helo")
	//fmt.Printf("output: %v", input("insert your name"))

	csvContent, err := readCsvFromFile("testfile.csv")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("slice content: %v\n", csvContent)

	for _, rowSlice := range csvContent {
		fmt.Println(calcolaCodiceFiscale(rowSlice[0], rowSlice[1]))
	}
}
