package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func get_freq(rune_slice []rune) map[rune]int {

	frequncy := make(map[rune]int, 1)
	for _, rune := range rune_slice {
		if _, ok := frequncy[rune]; ok {
			frequncy[rune] += 1
		} else {
			frequncy[rune] = 1
		}
	}

	return frequncy
}

func break_string(string_to_process string) []rune {
	var runes []rune
	for _, char := range string_to_process {
		if char == 10 {
			break
		}
		runes = append(runes, char)
	}

	return runes
}

func DecodeFrequency(frequency map[rune]int) map[rune]string {

	decode := make(map[rune]string, len(frequency))
	for k, val := range frequency {
		fmt.Printf("Известно что буква %v встречалась %v раз\nКакую букву подставим?", string(k), val)
		char := read_input()

		decode[k] = char
	}

	return decode
}

func CalcCharPercent(frequency map[rune]int) {

	var sum float32 = 0
	for _, val := range frequency {
		sum += float32(val)
	}

	for k, val := range frequency {
		var percent float32 = float32(val) * 100 / sum
		fmt.Printf("Частота буквы %v равна %v от общего количества \n", string(k), percent)
	}
}

func main() {
	var input string = read_input()
	runes := break_string(input)

	freq := get_freq(runes)

	CalcCharPercent(freq)

	decode := DecodeFrequency(freq)

	var output = make([]string, len(runes))
	for _, rune := range runes {
		output = append(output, strings.TrimSpace(decode[rune]))
	}
	fmt.Println("Результа дешифроки шифра Цезаря: ", output)
}
