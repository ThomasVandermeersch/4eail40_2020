package main

import (
	"fmt"
	"time"
)

type (
	eventObserver struct {
		id int
	}

	eventNotifier struct {
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
	// Initialize a new Notifier.
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
