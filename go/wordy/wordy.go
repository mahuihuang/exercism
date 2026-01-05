package wordy

import (
	"strconv"
	"strings"
)

// 1. 移除首位前后缀
// 2. 转换单词为运算符
// 3. 转换句子为数组
// 4. 遍历数组，并计算结果
func Answer(question string) (int, bool) {
	// re := regexp.MustCompile(`(?<=What is ).*?(?=\?)`)
	question = strings.Trim(strings.Trim(question, "What is "), "?")

	calExp := replaceOperator(question)
	if len(calExp) == 0 {
		return 0, false
	}
	result, err := strconv.Atoi(calExp[0])
	if err != nil {
		return result, false
	}
	for i := 1; i < len(calExp); i = i + 2 {
		numberIndex := i + 1
		if numberIndex >= len(calExp) {
			return 0, false
		}
		// covert string to int
		number, err := strconv.Atoi(calExp[numberIndex])
		if err != nil {
			return 0, false
		}
		// convert string to operator
		switch calExp[i] {
		case "+":
			result += number
		case "-":
			result -= number
		case "*":
			result *= number
		case "/":
			result /= number
		default:
			return 0, false
		}
	}
	return result, true
}

// replaceOperator replace word to Operator
func replaceOperator(question string) []string {
	question = strings.ReplaceAll(question, "plus", "+")
	question = strings.ReplaceAll(question, "minus", "-")
	question = strings.ReplaceAll(question, "divided by", "/")
	question = strings.ReplaceAll(question, "multiplied by", "*")

	return strings.Fields(question)
}
