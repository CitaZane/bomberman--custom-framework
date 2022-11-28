package websocket

import (
	"strconv"
	"time"
)

type Timer struct {
	stop     chan bool
	expired  bool
	interval int
	duration int
}

func newTimer(duration, interval int) *Timer {
	return &Timer{
		stop:     make(chan bool),
		duration: duration,
		interval: interval,
		expired:  true,
	}
}
func (t *Timer) start(pool *Pool) {
	t.expired = false
	ticker := time.NewTicker(time.Duration(t.interval) * time.Second)
	go func() {
		for {
			select {
			case <-t.stop:
				t.expired = true
				return
			case <-ticker.C:
				t.duration--
				pool.Timer <- Message{Type: "TIMER", Body: strconv.Itoa(int(t.duration))}
				if t.duration == 0 {
					ticker.Stop()
					t.expired = true
					return
				}
			}
		}
	}()
	time.Sleep(time.Duration(t.duration) * time.Second)
	ticker.Stop()
	t.expired = true
	t.stop <- true
}
