package demo_wallet

import (
	"errors"
	"time"
)

// 基于充血模型的DDD开发模式，和贫血模型的传统开发模式的主要区别是 Service 层。
// 把虚拟钱包 VirtualWallet 类 设计层一个充血的 Domain 领域模型，
// 将原来在 Service 类中的部分业务逻辑移动到 VirtualWallet 类中，
// 让 Service 类的实现依赖 VirtualWallet 类。

// VirtualWallet Domain 领域模型（充血模型），功能简单的时候，看起来来很淡薄。
// 增加一些复杂的功能是，优势就明显了，如增加透支和冻结功能。
// 功能继续演进，如增加更细化的冻结策略，透支策略，支持钱包账户ID自动生成逻辑（分布式 ID 生成算法）等，
// 那么就值得设计为充血模型，优势就更加明显了。
type VirtualWallet struct {
	id         string
	createTime time.Time
	balance    float64

	// 增加透支和冻结功能
	isAllowedOverdraft bool
	overdraftAmount    float64
	frozenAmount       float64
}

func NewVirtualWallet(preAllocatedId string) *VirtualWallet {
	return &VirtualWallet{
		id:              preAllocatedId,
		balance:         0,
		overdraftAmount: 0,
		frozenAmount:    0}
}

// 增加透支和冻结功能

func (w *VirtualWallet) Freeze(amount float64) {}

func (w *VirtualWallet) UnFreeze(amount float64) {}

func (w *VirtualWallet) IncreaseOverdraftAmount(amount float64) {}

func (w *VirtualWallet) DecreaseOverdraftAmount(amount float64) {}

func (w *VirtualWallet) CloseOverdraft() {}

func (w *VirtualWallet) OpenOverdraft() {}

func (w *VirtualWallet) Balance() float64 {
	return w.balance
}

func (w *VirtualWallet) GetAvailableAmount() float64 {
	totalAvailableBalance := w.balance - w.frozenAmount
	if w.isAllowedOverdraft {
		totalAvailableBalance += w.overdraftAmount
	}
	return totalAvailableBalance
}

func (w *VirtualWallet) Debit(amount float64) error {
	availableAmount := w.GetAvailableAmount()
	if availableAmount-amount < 0 {
		return errors.New("insufficient balance")
	}

	w.balance -= amount
	return nil
}

func (w *VirtualWallet) Credit(amount float64) error {
	availableAmount := w.GetAvailableAmount()
	if availableAmount < 0 {
		return errors.New("invalid amount")
	}

	w.balance += amount
	return nil
}

type DDDVirtualWalletService struct {
	walletRepo      VirtualWalletRepository
	transactionRepo VirtualWalletTransactionRepository
}

func (s *DDDVirtualWalletService) getVirtualWallet(walletId string) *VirtualWallet {
	walletEntity := s.walletRepo.GetWalletEntity(walletId)
	wallet := convert(walletEntity)
	return wallet
}

func (s *DDDVirtualWalletService) GetBalance(walletId string) float64 {
	return s.walletRepo.GetBalance(walletId)
}

func (s *DDDVirtualWalletService) Debit(walletId string, amount float64) error {
	walletEntity := s.walletRepo.GetWalletEntity(walletId)
	wallet := convert(walletEntity)
	if err := wallet.Debit(amount); err != nil {
		return err
	}

	transactionEntity := NewVirtualWalletTransactionEntity()
	transactionEntity.SetAmount(amount)
	transactionEntity.SetCreateTime(time.Now())
	transactionEntity.SetType(DEBIT)
	transactionEntity.SetFromWalletId(walletId)
	s.transactionRepo.SaveTransaction(transactionEntity)

	s.walletRepo.UpdateBalance(walletId, wallet.Balance())

	return nil
}

func (s *DDDVirtualWalletService) Credit(walletId string, amount float64) error {
	walletEntity := s.walletRepo.GetWalletEntity(walletId)
	wallet := convert(walletEntity)
	if err := wallet.Credit(amount); err != nil {
		return err
	}

	transactionEntity := NewVirtualWalletTransactionEntity()
	transactionEntity.SetAmount(amount)
	transactionEntity.SetCreateTime(time.Now())
	transactionEntity.SetType(DEBIT)
	transactionEntity.SetFromWalletId(walletId)
	s.transactionRepo.SaveTransaction(transactionEntity)

	s.walletRepo.UpdateBalance(walletId, wallet.Balance())

	return nil
}

func (s *DDDVirtualWalletService) Transfer() {
	// 与基于贫血模型的传统开发模式，代码一样
}

func convert(entity *VirtualWalletEntity) *VirtualWallet {
	return &VirtualWallet{}
}
