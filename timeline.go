package insights

import (
	"math"
	"sort"
	"time"
)

type Timeline map[time.Time]Conversation

func NewTimeline(c Conversation, d time.Duration) (t Timeline) {
	if d == 0 {
		d = time.Hour * 24
	}
	first, last := c.First(), c.Last()

	start := time.Date(first.Time.Year(), first.Time.Month(), first.Time.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(last.Time.Year(), last.Time.Month(), last.Time.Day(), 23, 59, 59, 999999999, time.UTC)

	days := int(math.Ceil(float64(end.Sub(start)) / float64(d)))

	cursor := start
	limit := start.Add(d).Add(time.Nanosecond * -1)

	t = make(Timeline, days)

	for i := 0; i < days; i++ {
		t[cursor] = Conversation{}

		for _, msg := range c {
			if msg.Time.After(cursor) && msg.Time.Before(limit) {
				t[cursor] = append(t[cursor], msg)
			}
		}

		cursor = cursor.Add(d)
		limit = cursor.Add(d).Add(time.Nanosecond * -1)
	}

	return
}

func (t Timeline) Extract() (count, words, letters map[time.Time]int) {
	count = make(map[time.Time]int, len(t))
	words = make(map[time.Time]int, len(t))
	letters = make(map[time.Time]int, len(t))

	for k, v := range t {
		count[k] = v.Count()
		words[k] = v.Words()
		letters[k] = v.Letters()
	}

	return
}

func (t Timeline) Days() (keys []time.Time) {
	for k := range t {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return keys
}

func (t Timeline) Most() (date time.Time, count int) {
	keys := t.Days()
	if len(keys) > 0 {
		date, count = keys[0], t[keys[0]].Count()
	}

	for _, k := range keys {
		if len(t[k]) > count {
			date, count = k, t[k].Count()
		}
	}

	return
}

func (t Timeline) Least() (date time.Time, count int) {
	keys := t.Days()
	if len(keys) > 0 {
		date, count = keys[0], t[keys[0]].Count()
	}

	for _, k := range keys {
		if len(t[k]) < count {
			date, count = k, t[k].Count()
		}
	}

	return
}

func (t Timeline) Average() (messages, words, letters float64) {
	m, w, l := 0, 0, 0
	for _, v := range t {
		m += v.Count()
		w += v.Words()
		l += v.Letters()
	}

	messages = float64(m) / float64(len(t))
	words = float64(w) / float64(len(t))
	letters = float64(l) / float64(len(t))

	return
}
