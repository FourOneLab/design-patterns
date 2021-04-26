package demo_authority

import (
	"crypto/md5"
	"fmt"
	"time"
)

const DefaultExpiredTimeInterval = 60 * time.Second

type AuthToken struct {
	token               string
	createTime          time.Time
	expiredTimeInterval time.Duration
}

func NewAuthToken(token string, createTime time.Time) *AuthToken {
	return &AuthToken{
		token:               token,
		createTime:          createTime,
		expiredTimeInterval: DefaultExpiredTimeInterval,
	}
}

// Generate helper function to generate token for server side.
func Generate(baseURL, appId, password string, createTime time.Time) *AuthToken {
	return &AuthToken{
		token:               generateToken(baseURL, appId, password, createTime),
		createTime:          createTime,
		expiredTimeInterval: DefaultExpiredTimeInterval,
	}
}

func generateToken(baseURL, appId, password string, createTime time.Time) string {
	raw := fmt.Sprintf("%s-%s-%s-%d", baseURL, appId, password, createTime.Unix())
	sum := md5.Sum([]byte(raw))
	return fmt.Sprintf("%x", sum)
}

func (t *AuthToken) GetToken() string {
	return t.token
}

func (t *AuthToken) IsExpired() bool {
	now := time.Now()
	if t.createTime.Add(t.expiredTimeInterval).Before(now) {
		return true
	}
	return false
}

func (t *AuthToken) Match(authToken *AuthToken) bool {
	return t.token == authToken.token
}
