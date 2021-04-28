package design_principles

import (
	"errors"
	"log"
	"net/http"
)

var NoAuthorizationErr = errors.New("no authorization runtime error")

type Transport interface {
	SendRequest(r *http.Request) error
}

type Transporter struct {
	httpClient *http.Client
}

func NewTransporter(httpClient *http.Client) *Transporter {
	return &Transporter{httpClient: httpClient}
}

func (t *Transporter) SendRequest(r *http.Request) error {
	return nil
}

type SecurityTransporter struct {
	appId    string
	appToken string
	*Transporter
}

func NewSecurityTransporter(appId string, appToken string, transporter *Transporter) *SecurityTransporter {
	return &SecurityTransporter{
		appId:       appId,
		appToken:    appToken,
		Transporter: transporter,
	}
}

// SendRequest
//
// - 修改前，如果 appId 或者 appToken 没有设置，不做校验
// - 修改后，如果 appId 或者 appToken 没有设置，则直接抛出 NoAuthorizationRuntimeException 未授权异常
func (s *SecurityTransporter) SendRequest(r *http.Request) error {
	// 这是修改前的代码
	//if s.appId != "" && s.appToken != "" {
	//	r.SetBasicAuth(s.appId, s.appToken)
	//}

	// 这是修改后的代码
	if s.appId == "" || s.appToken == "" {
		return NoAuthorizationErr
	}
	r.SetBasicAuth(s.appId, s.appToken)

	return s.Transporter.SendRequest(r)
}

type lspDemo struct{}

func (l *lspDemo) demoFunction(transport Transport) {
	request := new(http.Request)

	if err := transport.SendRequest(request); err != nil {
		log.Fatalln(err)
	}
}
