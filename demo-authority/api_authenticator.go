package demo_authority

import (
	"database/sql"
	"errors"
	"time"
)

var (
	TokenExpiredErr  = errors.New("token expired")
	TokenVerifiedErr = errors.New("token verification failed")
)

type ApiAuthenticator interface {
	HTTPAuth(url string) error
	RPCAuth(request ApiRequest) error
}

type DefaultApiAuthenticator struct {
	CredentialStorage
}

func NewDefaultApiAuthenticator(db *sql.DB) *DefaultApiAuthenticator {
	return &DefaultApiAuthenticator{CredentialStorage: NewMySQLCredentialStorage(db)}
}

func NewApiAuthenticator(credentialStorage CredentialStorage) *DefaultApiAuthenticator {
	return &DefaultApiAuthenticator{CredentialStorage: credentialStorage}
}

func (d *DefaultApiAuthenticator) HTTPAuth(url string) error {
	apiRequest, err := BuildFromURL(url)
	if err != nil {
		return err
	}

	return d.auth(apiRequest)
}

func (d *DefaultApiAuthenticator) RPCAuth(request ApiRequest) {
	panic("implement me")
}

func (d DefaultApiAuthenticator) auth(request *ApiRequest) error {
	appId := request.GetAppId()
	token := request.GetToken()
	timestamp := request.GetTimestamp()
	createTime := time.Unix(timestamp, 0)
	baseUrl := request.GetBaseUrl()

	clientAuthToken := NewAuthToken(token, createTime)
	if clientAuthToken.IsExpired() {
		return TokenExpiredErr
	}

	password, err := d.GetPassWordByAppId(appId)
	if err != nil {
		return err
	}

	serverAuthToken := Generate(baseUrl, appId, password, createTime)
	if !serverAuthToken.Match(clientAuthToken) {
		return TokenVerifiedErr
	}

	return nil
}
