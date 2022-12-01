package websocket

import (
	"time"
)

type TimerType string

const (
	None       TimerType = "none"
	START_GAME TimerType = "start_game"
	QUEUE      TimerType = "queue"
)

type Timer struct {
	Type     TimerType
	Expired  bool `json:"expired"`
	Duration int  `json:"duration"`
	interval int
	stop     chan bool
}

func newTimer(duration, interval int, timerType TimerType) *Timer {
	return &Timer{
		stop:     make(chan bool),
		Duration: duration,
		interval: interval,
		Expired:  true,
		Type:     timerType,
	}
}

func (t *Timer) start(pool *Pool) {
	t.Expired = false
	pool.Timer <- Message{Type: "TIMER", Timer: t} // send message to render timer on frontend instantly
	ticker := time.NewTicker(time.Duration(t.interval) * time.Second * 1)
	go func() {
	F:
		for {
			select {
			case <-t.stop:
				t.Expired = true
				ticker.Stop()
				break F
			case <-ticker.C:
				t.Duration--
				if t.Duration == 0 {
					ticker.Stop()
					t.Expired = true
					break F
				}
				pool.Timer <- Message{Type: "TIMER", Timer: t}
			}
		}
		// send last message to channel after we break at line 49 or 43
		pool.Timer <- Message{Type: "TIMER", Timer: t}
	}()
}
