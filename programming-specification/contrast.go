package programming_specification

import (
	"strings"
	"time"
)

// InvestV1 重构前
func InvestV1(userId, financialProductId string) {
	now := time.Now()
	lastDay := now.AddDate(0, 1, -now.Day())
	if lastDay.Day() == now.Day() {
		return
	}
}

// InvestV2 重构后
func InvestV2(userId, financialProductId string) {
	if isLastDayOfMonth(time.Now()) {
		return
	}
}

func isLastDayOfMonth(time time.Time) bool {
	lastDay := time.AddDate(0, 1, -time.Day())
	if lastDay.Day() == time.Day() {
		return true
	}
	return false
}

// ----------移除过深的嵌套----------

// 1. 去掉多余的 if 或 else 语句

type Order struct {
	count float64
	price float64
}

func (o *Order) Count() float64 {
	return o.count
}

func (o *Order) SetCount(count float64) {
	o.count = count
}

func (o *Order) Price() float64 {
	return o.price
}

func (o *Order) SetPrice(price float64) {
	o.price = price
}

func calculateTotalAmount(orders []*Order) float64 {
	if orders == nil || len(orders) == 0 {
		return 0.0
	} else { // 这里可以移除多余的 else
		amount := 0.0
		for _, order := range orders {
			if order != nil {
				amount += order.Count() * order.Price()
			}
		}
		return amount
	}
}

func MatchStringsV1(strList []string, subStr string) []string {
	matchedStrings := make([]string, 0)

	if strList != nil && len(strList) != 0 && len(subStr) != 0 {
		for _, s := range strList {
			if len(s) != 0 { // 可以与下面的 if 语句合并
				if strings.Contains(s, subStr) {
					matchedStrings = append(matchedStrings, s)
				}
			}
		}
	}

	return matchedStrings
}

// 2. 使用编程语言提供的 continue、break、return 关键字，提前退出嵌套
// 3. 调整执行语句减少嵌套

// MatchStringsV2 重构前
func MatchStringsV2(strList []string, subStr string) []string {
	matchedStrings := make([]string, 0)

	if strList != nil && len(strList) != 0 && len(subStr) != 0 {
		for _, s := range strList {
			if len(s) != 0 && strings.Contains(s, subStr) {
				matchedStrings = append(matchedStrings, s)
				// 这里还有其他逻辑
			}
		}
	}

	return matchedStrings
}

// MatchStringsV3 重构后
func MatchStringsV3(strList []string, subStr string) []string {
	matchedStrings := make([]string, 0)

	// 直接 return 减少嵌套
	// 调整执行语句减少嵌套
	if strList == nil || len(strList) == 0 || len(subStr) == 0 {
		return matchedStrings
	}

	for _, s := range strList {
		// 直接 continue 减少嵌套
		if len(s) == 0 || strings.Contains(s, subStr) {
			continue
		}

		matchedStrings = append(matchedStrings, s)
		// 这里还有其他逻辑
	}

	return matchedStrings
}

// 4. 将部分嵌套逻辑封装成函数调用，以此来减少嵌套

// AppendSaltsV1 重构前
func AppendSaltsV1(passwords []string) []string {
	if passwords == nil || len(passwords) == 0 {
		return []string{}
	}

	passwordsWithSalt := make([]string, 0)
	for _, password := range passwords {
		l := len(password)
		if l < 0 {
			continue
		}

		if l < 8 {
			// 执行处理逻辑
		} else {
			// 执行另外的处理逻辑
		}
	}

	return passwordsWithSalt
}

// AppendSaltsV2 重构后
func AppendSaltsV2(passwords []string) []string {
	if passwords == nil || len(passwords) == 0 {
		return []string{}
	}

	passwordsWithSalt := make([]string, 0)
	for _, password := range passwords {
		if len(password) < 0 {
			continue
		}

		passwordsWithSalt = append(passwordsWithSalt, appendSalt(password))
	}

	return passwordsWithSalt
}

func appendSalt(password string) string {
	passwordWithSalt := ""

	l := len(password)
	if l < 8 {
		// 执行处理逻辑
	} else {
		// 执行另外的处理逻辑
	}

	return passwordWithSalt
}

// ----------使用解释性变量----------

// 1. 常量替代魔法数字

// CalculateCircularAreaV1 重构前
func CalculateCircularAreaV1(radius float64) float64 {
	return 3.14 * radius * radius
}

const PI = 3.14 // 常量替代魔法数字

// CalculateCircularAreaV2 重构后
func CalculateCircularAreaV2(radius float64) float64 {
	return PI * radius * radius
}

// 2. 使用解释性变量来解释复杂表达式

var (
	SummerStart = time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC)
	SummerEnd   = time.Date(time.Now().Year(), 8, 31, 0, 0, 0, 0, time.UTC)
	isSummer    = time.Now().After(SummerStart) && time.Now().Before(SummerEnd)
)

func demo() {
	if isSummer {
		// do something
	} else {
		// do something
	}

}
