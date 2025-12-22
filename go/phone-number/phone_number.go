package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

// NANP 电话系统， 一共有 10 未数字组成，前三个数字为区号，后七位为本地号码。
// 1. 有时电话号码前会带国家编码，数字或加号与数字的组合
// 2. 区号和本地号码的区间范围是 2-9
// 格式 NXX NXX-XXXX， N 代表 2-9，X 代表 0-9

// Number 格式电话号码，移除标点符号和国家代码，当前联系仅北美的国家的代号有效
func Number(phoneNumber string) (string, error) {
	errorMsg := errors.New("invalid phone number")
	if phoneNumber == "" {
		return "", errorMsg
	}
	// 获取所有的数字
	position := 0
	numericalPhoneNumber := []rune{}
	for _, digit := range phoneNumber {
		if unicode.IsDigit(digit) {
			numericalPhoneNumber = append(numericalPhoneNumber, digit)
			position++
		}
	}
	// 判断长度合法性
	if position == 11 {
		if numericalPhoneNumber[0] != '1' {
			return "", errorMsg
		}
		numericalPhoneNumber = numericalPhoneNumber[1:]
	} else if position != 10 {
		return "", errorMsg
	}
	// 判断区号和本地号码开头是否合法
	if numericalPhoneNumber[0] < '2' || numericalPhoneNumber[3] < '2' {
		return "", errorMsg
	}
	return string(numericalPhoneNumber), nil
}

// AreaCode 返回区号
func AreaCode(phoneNumber string) (string, error) {
	pn, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return pn[0:3], nil
}

// Format 返回被格式化的电话号码，(NXX) NXX-XXX
// 被格式化的电话号码的列表
// +1 (613)-995-0253
// 613-995-0253
// 1 613 995 0253
// 613.995.0253
// 格式化后的电话号码
// (613) 995-0253
func Format(phoneNumber string) (string, error) {
	pn, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", pn[0:3], pn[3:6], pn[6:]), nil
}
