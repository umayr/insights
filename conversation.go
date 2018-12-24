package insights

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const regexSplit = `\[.*].*:\s`

const delimiter = string(0xEF) + "##" + string(0xEF)

type Conversation []Message

func NewConversation(src string, tz string) (c Conversation, err error) {
	raw := regexp.MustCompile(regexSplit).ReplaceAllStringFunc(src, func(s string) string {
		return delimiter + s
	})

	msgs := strings.Split(raw, delimiter)

	if tz == "" {
		tz = "Local"
	}

	l, err := time.LoadLocation(tz)
	if err != nil {
		return
	}

	for _, raw := range msgs {
		if raw == "" {
			continue
		}

		m, err := NewMessage(raw, l)
		if err != nil {
			return nil, err
		}
		c = append(c, m)
	}

	return
}

func (c Conversation) First() (m Message) {
	if len(c) == 0 {
		return
	}

	return c[0]
}

func (c Conversation) Last() (m Message) {
	if len(c) == 0 {
		return
	}

	return c[len(c)-1]
}

func (c Conversation) Duration() time.Duration {
	first, last := c.First(), c.Last()
	return last.Time.Sub(first.Time)
}

func (c Conversation) Count() int {
	return len(c)
}

func (c Conversation) combine() (str string) {
	for _, m := range c {
		if m.Type == MessageTypeText {
			str += m.Text + " "
		}
	}

	return str
}

func (c Conversation) Words() int {
	return len(
		strings.Fields(
			strings.Replace(
				c.combine(),
				"\n",
				"",
				-1),
		),
	)
}

func (c Conversation) Letters() int {
	return len(strings.Join(strings.Fields(c.combine()), ""))
}

func (c Conversation) Timeline() (t Timeline) {
	return NewTimeline(c)
}

func (c Conversation) Participants() (p Participants) {
	return NewParticipants(c)
}

func (c Conversation) Contribution() (contrib Contribution) {
	return NewContribution(c)
}

func (c Conversation) Average() (words, letters float64) {
	w, l := 0, 0
	for _, v := range c {
		w += len(
			strings.Fields(
				strings.Replace(
					v.Text,
					"\n",
					"",
					-1),
			),
		)
		l += len(strings.Join(strings.Fields(v.Text), ""))
	}

	words = float64(w) / float64(len(c))
	letters = float64(l) / float64(len(c))

	return
}

func (c Conversation) Frequency() (m map[string]int) {
	m = map[string]int{
		"0h":  0,
		"1h":  0,
		"2h":  0,
		"3h":  0,
		"4h":  0,
		"5h":  0,
		"6h":  0,
		"7h":  0,
		"8h":  0,
		"9h":  0,
		"10h": 0,
		"11h": 0,
		"12h": 0,
		"13h": 0,
		"14h": 0,
		"15h": 0,
		"16h": 0,
		"17h": 0,
		"18h": 0,
		"19h": 0,
		"20h": 0,
		"21h": 0,
		"22h": 0,
		"23h": 0,
	}

	for _, v := range c {
		k := fmt.Sprintf("%dh", v.Time.Hour())
		m[k] = m[k] + 1
	}

	return
}

func (c Conversation) Emoji() map[string]int {
	return CountEmoji(c.combine())
}
