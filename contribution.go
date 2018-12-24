package insights

type Contribution map[string]Conversation

func NewContribution(c Conversation) (contrib Contribution) {
	participants := c.Participants()

	contrib = make(Contribution, len(participants))

	for _, p := range participants {
		contrib[p] = Conversation{}

		for _, msg := range c {
			if msg.Sender == p {
				contrib[p] = append(contrib[p], msg)
			}
		}
	}

	return
}

func (c Contribution) Extract() (count, words, letters map[string]int, frequency map[string]map[string]int) {
	var participants Participants
	for k := range c {
		participants = append(participants, k)
	}

	count = make(map[string]int, len(c))
	words = make(map[string]int, len(c))
	letters = make(map[string]int, len(c))
	frequency = make(map[string]map[string]int, len(c))

	for _, k := range participants {
		count[k] = c[k].Count()
		words[k] = c[k].Words()
		letters[k] = c[k].Letters()
		frequency[k] = c[k].Frequency()
	}

	return
}
