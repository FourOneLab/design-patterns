package specification_refactoring

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

const transactionIdPrefix = "t_"

var InvalidTransactionErr = errors.New("invalid transaction")

type Status int

const (
	Initial Status = iota
	Executed
	Expired
	Failed
)

type Transaction struct {
	Id                  string
	BuyerId             string
	SellerId            string
	ProductId           string
	OrderId             string
	createdTime         time.Time
	Amount              int64
	WalletTransactionId string

	WalletRPC

	status Status
	mu     sync.Mutex
}

func NewTransaction(id string, buyerId string, sellerId string, productId string, orderId string, rpc WalletRPC) *Transaction {
	return &Transaction{
		Id:          fillTransactionId(id),
		BuyerId:     buyerId,
		SellerId:    sellerId,
		ProductId:   productId,
		OrderId:     orderId,
		createdTime: time.Now(),
		WalletRPC:   rpc,
		mu:          sync.Mutex{},
	}
}

// fillTransactionId helper function for transaction id
func fillTransactionId(id string) string {
	if id == "" {
		uuid, err := uuid.NewUUID()
		if err != nil {
			return ""
		}
		id = uuid.String()
	}

	if !strings.HasPrefix(id, transactionIdPrefix) {
		id = transactionIdPrefix + id
	}

	return id
}

func (t *Transaction) Execute() (bool, error) {
	if t.Id == "" || t.SellerId == "" || t.Amount < 0 {
		return false, InvalidTransactionErr
	}

	if t.status == Executed {
		return true, nil
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.status == Executed {
		// double check
		return true, nil
	}

	if t.isExpired() {
		t.status = Expired
		return false, nil
	}

	//walletRPC := NewWalletRPCService()
	walletTransactionId := t.WalletRPC.MoveMoney(t.Id, t.BuyerId, t.SellerId, t.Amount)
	if walletTransactionId != "" {
		t.WalletTransactionId = walletTransactionId
		t.status = Executed
		return true, nil
	}

	t.status = Failed
	return false, nil
}

func (t *Transaction) isExpired() bool {
	return t.createdTime.Add(14 * 24 * time.Hour).Before(time.Now())
}

type WalletRPC interface {
	MoveMoney(id, buyerId, sellerId string, amount int64) string
}

type WalletRPCService struct{}

func NewWalletRPCService() WalletRPC {
	return &WalletRPCService{}
}

func (s *WalletRPCService) MoveMoney(id, buyerId, sellerId string, amount int64) string {
	return ""
}

type MockWalletRPCServiceOne struct{}

func (m *MockWalletRPCServiceOne) MoveMoney(id, buyerId, sellerId string, amount int64) string {
	return "123abc"
}

type MockWalletRPCServiceTwo struct{}

func (m *MockWalletRPCServiceTwo) MoveMoney(id, buyerId, sellerId string, amount int64) string {
	return ""
}
