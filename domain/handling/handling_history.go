package handling

import "sort"

var EmptyHistory = NewHistory(make([]*Event, 0))

// History means the handling history of a cargo.
type History struct {
	events []*Event
}

func NewHistory(events []*Event) *History {
	if events == nil {
		events = make([]*Event, 0)
	}

	return &History{events: events}
}

func (h *History) Events() []*Event {
	return h.events
}

func (h *History) DistinctEventsByCompletionTime() []*Event {
	if h.events == nil || len(h.events) == 0 {
		return nil
	}

	sort.Slice(h.events, func(i, j int) bool {
		if i < len(h.events) && j < len(h.events) {
			return h.events[i].completionTime.Before(h.events[j].completionTime)
		}

		return false
	})

	return h.events
}

func (h *History) MostRecentlyCompletedEvent() *Event {
	events := h.DistinctEventsByCompletionTime()
	if events == nil || len(events) == 0 {
		return nil
	}

	return h.events[len(h.events)-1]
}
