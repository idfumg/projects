package store

import (
	"fmt"
	"log"
	"myapp/config"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

var testpg *Pg
var once sync.Once

const (
	WaitTillTestDoneDuration = 3 * time.Second
)

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

type TestLogger struct {
}

func (l *TestLogger) Printf(s string, params ...any) {

}

func setup() {
	once.Do(func() {
		// logger := log.New(os.Stdout, "Service: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger := &TestLogger{}

		config, err := config.New("../.env")
		if err != nil {
			log.Fatalf("Error! Could not init the config: %v\n", err)
		}

		testpg, err = NewPg(logger, config)
		if err != nil {
			log.Fatalf("Error! Could not init a store: %v\n", err)
		}
	})
}

func teardown() {
	if testpg == nil {
		// No tests used the database!
		return
	}

	defer testpg.db.(*sqlx.DB).Close()

	startTime := time.Now()
	for {
		open := testpg.db.(*sqlx.DB).Stats().OpenConnections
		if open > 0 {
			if time.Since(startTime) > WaitTillTestDoneDuration {
				printStats()
				log.Panicf("failed to close %d connections", open)
			}
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}
}

func printStats() {
	stats := testpg.db.(*sqlx.DB).Stats()
	fmt.Printf("MaxOpenConnections: %d\n", stats.MaxOpenConnections)
	fmt.Printf("OpenConnections: %d\n", stats.OpenConnections)
	fmt.Printf("InUse: %d\n", stats.InUse)
	fmt.Printf("Idle: %d\n", stats.Idle)
	fmt.Printf("WaitCount: %d\n", stats.WaitCount)
	fmt.Printf("WaitDuration: %d\n", stats.WaitDuration)
	fmt.Printf("MaxIdleClosed: %d\n", stats.MaxIdleClosed)
	fmt.Printf("MaxIdleTimeClosed: %d\n", stats.MaxIdleTimeClosed)
	fmt.Printf("MaxLifetimeClosed: %d\n", stats.MaxLifetimeClosed)
}
