package design_principles

import (
	"errors"
	"strings"
)

var (
	InvalidUsernameErr = errors.New("invalid username")
	InvalidPasswordErr = errors.New("invalid password")
)

type UserAuthenticator struct{}

func (u *UserAuthenticator) Authenticate(username, password string) error {
	if !u.isValidUsername(username) {
		return InvalidUsernameErr
	}

	if !u.isValidPassword(password) {
		return InvalidPasswordErr
	}

	return nil
}

func (u *UserAuthenticator) isValidUsername(username string) bool {
	// 用户名不能为空
	if username == "" {
		return false
	}

	// 用户名长度限制
	i := len(username)
	if i < 4 || i > 64 {
		return false
	}

	// 用户名只能小写
	lower := strings.ToLower(username)
	if lower != username {
		return false
	}

	// 用户名只能是数字、小写字母和点（.）组成
	for _, c := range []byte(username) {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.') {
			return false
		}
	}

	return true
}

func (u *UserAuthenticator) isValidPassword(password string) bool {
	// 用户名不能为空
	if password == "" {
		return false
	}

	// 用户名长度限制
	i := len(password)
	if i < 4 || i > 64 {
		return false
	}

	// 用户名只能小写
	lower := strings.ToLower(password)
	if lower != password {
		return false
	}

	// 用户名只能是数字、小写字母和点（.）组成
	for _, c := range []byte(password) {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.') {
			return false
		}
	}

	return true
}

// ------------------------------
// 合并重复代码

type UserAuthenticatorV2 struct{}

func (u *UserAuthenticatorV2) Authenticate(username, password string) error {
	if !u.isValidUsernameOrPassword(username) {
		return InvalidUsernameErr
	}

	if !u.isValidUsernameOrPassword(password) {
		return InvalidPasswordErr
	}

	return nil
}

func (u *UserAuthenticatorV2) isValidUsernameOrPassword(str string) bool {
	// 用户名不能为空
	if str == "" {
		return false
	}

	// 用户名长度限制
	i := len(str)
	if i < 4 || i > 64 {
		return false
	}

	// 用户名只能小写
	lower := strings.ToLower(str)
	if lower != str {
		return false
	}

	// 用户名只能是数字、小写字母和点（.）组成
	for _, c := range []byte(str) {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '.') {
			return false
		}
	}

	return true
}
