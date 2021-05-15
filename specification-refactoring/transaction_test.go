package specification_refactoring

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestTransaction_Execute(t1 *testing.T) {
	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			"normal",
			fields{
				Id:                  "111",
				BuyerId:             "222",
				SellerId:            "333",
				ProductId:           "444",
				OrderId:             "555",
				createdTime:         time.Now(),
				Amount:              666,
				WalletTransactionId: "777",
				WalletRPC:           new(MockWalletRPCServiceOne),
				status:              Initial,
				mu:                  sync.Mutex{},
			},
			true,
			false,
		},
		{
			"invalid transaction",
			fields{
				Id:                  "111",
				BuyerId:             "",
				SellerId:            "",
				ProductId:           "222",
				OrderId:             "333",
				createdTime:         time.Now(),
				Amount:              -1,
				WalletTransactionId: "444",
				WalletRPC:           new(MockWalletRPCServiceOne),
				status:              Initial,
				mu:                  sync.Mutex{},
			},
			false,
			true,
		},
		{
			"expired",
			fields{
				Id:                  "111",
				BuyerId:             "222",
				SellerId:            "333",
				ProductId:           "444",
				OrderId:             "555",
				createdTime:         time.Now().Add(-15 * 24 * time.Hour),
				Amount:              666,
				WalletTransactionId: "777",
				WalletRPC:           new(MockWalletRPCServiceOne),
				status:              Initial,
				mu:                  sync.Mutex{},
			},
			false,
			false,
		},
		{
			"executed",
			fields{
				Id:                  "111",
				BuyerId:             "222",
				SellerId:            "333",
				ProductId:           "444",
				OrderId:             "555",
				createdTime:         time.Now(),
				Amount:              666,
				WalletTransactionId: "777",
				WalletRPC:           new(MockWalletRPCServiceOne),
				status:              Executed,
				mu:                  sync.Mutex{},
			},
			true,
			false,
		},
		{
			"rpc failed",
			fields{
				Id:                  "111",
				BuyerId:             "222",
				SellerId:            "333",
				ProductId:           "444",
				OrderId:             "555",
				createdTime:         time.Now(),
				Amount:              666,
				WalletTransactionId: "777",
				WalletRPC:           new(MockWalletRPCServiceTwo),
				status:              Initial,
				mu:                  sync.Mutex{},
			},
			false,
			false,
		},
		{
			"executing",
			fields{
				Id:                  "111",
				BuyerId:             "222",
				SellerId:            "333",
				ProductId:           "444",
				OrderId:             "555",
				createdTime:         time.Now(),
				Amount:              666,
				WalletTransactionId: "777",
				WalletRPC:           new(MockWalletRPCServiceOne),
				status:              0,
				mu:                  sync.Mutex{},
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				Id:                  tt.fields.Id,
				BuyerId:             tt.fields.BuyerId,
				SellerId:            tt.fields.SellerId,
				ProductId:           tt.fields.ProductId,
				OrderId:             tt.fields.OrderId,
				createdTime:         tt.fields.createdTime,
				Amount:              tt.fields.Amount,
				WalletTransactionId: tt.fields.WalletTransactionId,
				WalletRPC:           tt.fields.WalletRPC,
				status:              tt.fields.status,
				mu:                  tt.fields.mu,
			}
			got, err := t.Execute()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransaction(t *testing.T) {
	type args struct {
		id        string
		buyerId   string
		sellerId  string
		productId string
		orderId   string
		rpc       WalletRPC
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{
			"empty id",
			args{
				id:        "",
				buyerId:   "111",
				sellerId:  "222",
				productId: "333",
				orderId:   "444",
				rpc:       nil,
			},
			&Transaction{
				Id:                  "t_8b13ba9c-b3b4-11eb-b91b-2af8bc697ffd",
				BuyerId:             "111",
				SellerId:            "222",
				ProductId:           "333",
				OrderId:             "444",
				createdTime:         time.Now(),
				Amount:              0,
				WalletTransactionId: "",
				WalletRPC:           nil,
				status:              Initial,
				mu:                  sync.Mutex{},
			},
		},
		{
			"normal",
			args{
				id:        "t_b7edfd84-b3b4-11eb-9570-2af8bc697ffd",
				buyerId:   "111",
				sellerId:  "222",
				productId: "333",
				orderId:   "444",
				rpc:       nil,
			},
			&Transaction{
				Id:                  "t_b7edfd84-b3b4-11eb-9570-2af8bc697ffd",
				BuyerId:             "111",
				SellerId:            "222",
				ProductId:           "333",
				OrderId:             "444",
				createdTime:         time.Now(),
				Amount:              0,
				WalletTransactionId: "",
				WalletRPC:           nil,
				status:              Initial,
				mu:                  sync.Mutex{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransaction(tt.args.id, tt.args.buyerId, tt.args.sellerId, tt.args.productId, tt.args.orderId, tt.args.rpc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
