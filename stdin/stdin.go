package stdin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// StrStdin 一列の文字列を標準入力から取得する
func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput := scanner.Text()

	return strings.TrimSpace(stringInput)
}

// IntStdin 整数値を標準入力から取得する
func IntStdin() (int, error) {
	str := StrStdin()
	return strconv.Atoi(str)
}

// SplitWithoutEmpty 空白や空文字だけの値を除去する
func SplitWithoutEmpty(stringTarget string, delim string) []string {
	stringSplited := strings.Split(stringTarget, delim)
	stringReturned := []string{}

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}
	return stringReturned
}

// SplitStrStdin デリミタで分割して文字列スライスを取得する
func SplitStrStdin(delim string) []string {
	str := StrStdin()
	return SplitWithoutEmpty(str, delim)
}

// SplitIntStdin デリミタで分割して整数値スライスを取得する
func SplitIntStdin(delim string) ([]int, error) {
	strs := SplitStrStdin(" ")
	intReturned := []int{}
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		intReturned = append(intReturned, n)
	}
	return intReturned, nil
}
