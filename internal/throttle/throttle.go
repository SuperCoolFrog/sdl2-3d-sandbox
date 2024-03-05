package throttle

import "time"

type Throttle struct {
	duration time.Duration
	last     time.Time
	counter  time.Duration
}

func New(duration time.Duration) *Throttle {
	return &Throttle{
		duration: duration,
		counter:  0,
		last:     time.Now(),
	}
}

func (t *Throttle) Do(f func()) {
	now := time.Now()
	t.counter += now.Sub(t.last)

	if t.counter >= t.duration {
		f()
		t.counter = time.Duration(0)
	}

	t.last = now
}

/**
var d time.Duration = 1000000000 // One second
th := throttle.New(d)
th.Do(func() {
	mouseX, mouseY, s := sdl.GetMouseState()
	fmt.Printf("(%d, %d, %d)\n", mouseX, mouseY, s)
})
**/
