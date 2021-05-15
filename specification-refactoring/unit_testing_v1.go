package specification_refactoring

import (
	"errors"
	"strconv"
	"strings"
)

type Text struct {
	content string
}

func NewText(content string) *Text {
	return &Text{content: content}
}

// ToNumber 将字符串转化成数字，忽略字符串中的首尾空格；
// 如果字符串中包含除首尾空格之外的非数字字符，则返回null。
func (t Text) ToNumber() (int, error) {
	if t.content == "" {
		return 0, errors.New("empty content")
	}

	return strconv.Atoi(strings.TrimSpace(t.content))
}
