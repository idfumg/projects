package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("received timeout")
	ErrInterrupt = errors.New("received interrupt")
)

type Runner struct {
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
	tasks     []func(int)
}

func New(d time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(d),
		interrupt: make(chan os.Signal, 1),
		tasks:     nil,
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
