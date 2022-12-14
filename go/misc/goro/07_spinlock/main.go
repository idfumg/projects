package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type SpinLock int32

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}

func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}

const (
	totalAccounts  = 50000
	maxAmountMoved = 10
	initialMoney   = 100
	threads        = 4
)

func performMovements(ledger *[totalAccounts]int32, locks *[totalAccounts]sync.Locker, totalTrans *int64) {
	for {
		accountA := rand.Intn(totalAccounts)
		accountB := rand.Intn(totalAccounts)
		for accountA == accountB {
			accountB = rand.Intn(totalAccounts)
		}
		amountToMove := rand.Int31n(maxAmountMoved)
		toLock := []int{accountA, accountB}
		sort.Ints(toLock)

		locks[toLock[0]].Lock()
		locks[toLock[1]].Lock()

		atomic.AddInt32(&ledger[accountA], -amountToMove)
		atomic.AddInt32(&ledger[accountB], amountToMove)
		atomic.AddInt64(totalTrans, 1)

		locks[toLock[1]].Unlock()
		locks[toLock[0]].Unlock()
	}
}

func main() {
	fmt.Println("Total accounts:", totalAccounts, ", total threads:", threads, ", using Spinlocks")

	var ledger [totalAccounts]int32
	var locks [totalAccounts]sync.Locker
	var totalTrans int64

	for i := 0; i < totalAccounts; i++ {
		ledger[i] = initialMoney
		locks[i] = NewSpinLock()
		// locks[i] = &sync.Mutex{}
	}

	for i := 0; i < threads; i++ {
		go performMovements(&ledger, &locks, &totalTrans)
	}

	for {
		time.Sleep(2 * time.Second)
		var sum int32
		for i := 0; i < totalAccounts; i++ {
			locks[i].Lock()
		}
		for i := 0; i < totalAccounts; i++ {
			sum += ledger[i]
		}
		for i := 0; i < totalAccounts; i++ {
			locks[i].Unlock()
		}
		fmt.Println(totalTrans, sum)
	}
}
