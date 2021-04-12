package object_oriented

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// 在金融系统中给每个用户创建一个虚拟钱包，用来记录用户在系统中的虚拟货币量。
//
// 包含四个属性，参照封装特性，对这四个属性对访问进行限制。
// 对于封装这个属性，需要编程语言本身提供一种语法机制（访问权限控制）来支持。
//
// Golang 中首字母小写对属性只能在包内访问，首字母大写则包外可以直接访问。
//
// 封装是为了达到隐藏信息和保护数据的目的。

type Wallet struct {
	id                      uuid.UUID
	createTime              time.Time
	balance                 float64
	balanceLastModifiedTime time.Time

	mu sync.RWMutex
}

func NewWallet() *Wallet {
	return &Wallet{
		id:                      uuid.New(),
		createTime:              time.Now(),
		balance:                 0,
		balanceLastModifiedTime: time.Now(),
	}
}

func (w *Wallet) String() string {
	return fmt.Sprintf("id: %s \ncreateTime: %s \nbalance: %f \nmodifiedTime: %s",
		w.id.String(), w.createTime.String(), w.balance, w.balanceLastModifiedTime.String())
}

func (w *Wallet) GetID() uuid.UUID {
	return w.id
}

func (w *Wallet) GetCreateTime() time.Time {
	return w.createTime
}

func (w *Wallet) GetBalance() float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.balance
}

func (w *Wallet) GetBalanceLastModifiedTime() time.Time {
	return w.balanceLastModifiedTime
}

func (w *Wallet) IncreaseBalance(increaseAmount float64) error {
	if increaseAmount < 0 {
		return errors.New("increase amount less than zero")
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	w.balance += increaseAmount
	w.balanceLastModifiedTime = time.Now()
	return nil
}

func (w *Wallet) DecreaseBalance(decreaseAmount float64) error {
	if decreaseAmount < 0 {
		return errors.New("decrease amount less than zero")
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	w.balance -= decreaseAmount
	w.balanceLastModifiedTime = time.Now()
	return nil
}
