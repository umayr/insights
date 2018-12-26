package insights

import (
	"time"
)

type Insights struct {
	First     *Message       `json:"first"`
	Last      *Message       `json:"last"`
	Duration  time.Duration  `json:"duration"`
	Frequency map[string]int `json:"frequency"`

	TotalMessages int `json:"totalMessages"`
	TotalWords    int `json:"totalWords"`
	TotalLetters  int `json:"totalLetters"`

	AverageWordsPerMessage   float64 `json:"averageWordsPerMessage"`
	AverageLettersPerMessage float64 `json:"averageLettersPerMessage"`

	Participants Participants `json:"participants"`

	Contribution          Contribution              `json:"contribution"`
	ContributionCount     map[string]int            `json:"contributionCount"`
	ContributionWords     map[string]int            `json:"contributionWords"`
	ContributionLetters   map[string]int            `json:"contributionLetters"`
	ContributionFrequency map[string]map[string]int `json:"contributionFrequency"`

	Timeline        Timeline          `json:"timeline"`
	TimelineCount   map[time.Time]int `json:"timelineCount"`
	TimelineWords   map[time.Time]int `json:"timelineWords"`
	TimelineLetters map[time.Time]int `json:"timelineLetters"`

	MostActiveDay    time.Time `json:"mostActiveDay"`
	MostActiveCount  int       `json:"mostActiveCount"`
	LeastActiveDay   time.Time `json:"leastActiveDay"`
	LeastActiveCount int       `json:"leastActiveCount"`

	AverageMessagesPerDay float64 `json:"averageMessagesPerDay"`
	AverageWordsPerDay    float64 `json:"averageWordsPerDay"`
	AverageLettersPerDay  float64 `json:"averageLettersPerDay"`

	EmojisUsed map[string]map[string]int `json:"emojisUsed"`
}

func New(c Conversation) *Insights {
	i := &Insights{}
	i.First, i.Last, i.Duration = c.First(), c.Last(), c.Duration()
	i.Frequency = c.Frequency()
	i.TotalMessages, i.TotalWords, i.TotalLetters = c.Count(), c.Words(), c.Letters()
	i.AverageWordsPerMessage, i.AverageLettersPerMessage = c.Average()
	i.Participants = c.Participants()
	i.Contribution = c.Contribution()
	i.ContributionCount, i.ContributionWords, i.ContributionLetters, i.ContributionFrequency = i.Contribution.Extract()

	i.Timeline = c.Timeline(time.Hour * 24 * 7)
	i.TimelineCount, i.TimelineWords, i.TimelineLetters = i.Timeline.Extract()

	i.MostActiveDay, i.MostActiveCount = i.Timeline.Most()
	i.LeastActiveDay, i.LeastActiveCount = i.Timeline.Least()

	i.AverageMessagesPerDay, i.AverageWordsPerDay, i.AverageLettersPerDay = i.Timeline.Average()
	i.EmojisUsed = c.Emoji()

	return i
}
