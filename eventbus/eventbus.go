package eventbus

import "sync"

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, rawData interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		channels := append(DataChannelSlice{}, chans...)
		go func(dataEvent DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- dataEvent
			}
		}(DataEvent{Data: rawData, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

//S is stand for Singleton/Super/Server EventBus
var S = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

// New You can create new EventBus as you want
func New() *EventBus {
	return &EventBus{
		subscribers: map[string]DataChannelSlice{},
	}
}
