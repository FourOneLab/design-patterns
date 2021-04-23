package demo_wallet

import (
	"errors"
	"time"
)

// 典型的三层结构，Controller 和 VO 负责暴露接口

type VirtualWalletController struct {
	walletService VirtualWalletService
}

// GetBalance 查询余额
func (c *VirtualWalletController) GetBalance(walletId string) {}

// Debit 出账
func (c *VirtualWalletController) Debit(walletId string, amount float64) {}

// Credit 入账
func (c *VirtualWalletController) Credit(walletId string, amount float64) {}

// Transfer 转账
func (c *VirtualWalletController) Transfer(fromWalletId, toWalletId string, amount float64) {}

//------------------------------
// Service 和 BO 负责核心业务逻辑

// TransactionType
const (
	DEBIT = iota
	CREDIT
	TRANSFER
)

type VirtualWalletBo struct {
	id         string
	createTime time.Time
	balance    float64
}

type VirtualWalletService struct {
	walletRepo      VirtualWalletRepository
	transactionRepo VirtualWalletTransactionRepository
}

func (s *VirtualWalletService) GetVirtualWallet(walletId string) *VirtualWalletBo {
	walletEntity := s.walletRepo.GetWalletEntity(walletId)

	convert := func(entity *VirtualWalletEntity) *VirtualWalletBo {
		return &VirtualWalletBo{}
	}

	walletBo := convert(walletEntity)
	return walletBo
}

func (s *VirtualWalletService) GetBalance(walletId string) float64 {
	return s.walletRepo.GetBalance(walletId)
}

func (s *VirtualWalletService) Debit(walletId string, amount float64) error {
	walletEntity := s.walletRepo.GetWalletEntity(walletId)
	balance := walletEntity.GetBalance()

	if balance-amount < 0 {
		return errors.New("insufficient balance")
	}

	transactionEntity := NewVirtualWalletTransactionEntity()
	transactionEntity.SetAmount(amount)
	transactionEntity.SetCreateTime(time.Now())
	transactionEntity.SetType(DEBIT)
	transactionEntity.SetFromWalletId(walletId)
	s.transactionRepo.SaveTransaction(transactionEntity)
	s.walletRepo.UpdateBalance(walletId, balance-amount)

	return nil
}

func (s *VirtualWalletService) Credit(walletId string, amount float64) {
	transactionEntity := NewVirtualWalletTransactionEntity()
	transactionEntity.SetAmount(amount)
	transactionEntity.SetCreateTime(time.Now())
	transactionEntity.SetType(CREDIT)
	transactionEntity.SetFromWalletId(walletId)
	s.transactionRepo.SaveTransaction(transactionEntity)

	walletEntity := s.walletRepo.GetWalletEntity(walletId)
	balance := walletEntity.GetBalance()
	s.walletRepo.UpdateBalance(walletId, balance+amount)
}

func (s *VirtualWalletService) Transfer(fromWalletId, toWalletId string, amount float64) error {
	transactionEntity := NewVirtualWalletTransactionEntity()
	transactionEntity.SetAmount(amount)
	transactionEntity.SetCreateTime(time.Now())
	transactionEntity.SetType(TRANSFER)
	transactionEntity.SetFromWalletId(fromWalletId)
	transactionEntity.SetToWalletId(toWalletId)
	s.transactionRepo.SaveTransaction(transactionEntity)

	err := s.Debit(fromWalletId, amount)
	if err != nil {
		return err
	}

	s.Credit(toWalletId, amount)
	return nil
}

//------------------------------
// Repository  和 Entity 负责数据存储

type VirtualWalletEntity struct {
	balance float64
}

func (e VirtualWalletEntity) GetBalance() float64 {
	return e.balance
}

type VirtualWalletRepository struct{}

func (r *VirtualWalletRepository) GetWalletEntity(walletId string) *VirtualWalletEntity {
	return &VirtualWalletEntity{}
}

func (r *VirtualWalletRepository) GetBalance(walletId string) float64 {
	return 0
}

func (r VirtualWalletRepository) UpdateBalance(walletId string, amount float64) {}

type VirtualWalletTransactionEntity struct {
	amount          float64
	createTime      time.Time
	transactionType int
	fromWalletId    string
	toWalletId      string
}

func NewVirtualWalletTransactionEntity() *VirtualWalletTransactionEntity {
	return &VirtualWalletTransactionEntity{}
}

func (e *VirtualWalletTransactionEntity) SetAmount(amount float64) {
	e.amount = amount
}

func (e *VirtualWalletTransactionEntity) SetCreateTime(now time.Time) {
	e.createTime = now
}

func (e *VirtualWalletTransactionEntity) SetType(transactionType int) {
	e.transactionType = transactionType
}

func (e *VirtualWalletTransactionEntity) SetFromWalletId(walletId string) {
	e.fromWalletId = walletId
}

func (e *VirtualWalletTransactionEntity) SetToWalletId(walletId string) {
	e.toWalletId = walletId
}

type VirtualWalletTransactionRepository struct{}

func (r VirtualWalletTransactionRepository) SaveTransaction(entity *VirtualWalletTransactionEntity) {}
