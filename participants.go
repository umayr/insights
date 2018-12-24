package insights

type Participants []string

func NewParticipants(c Conversation) (p Participants) {
	for _, m := range c {
		if !p.includes(m.Sender) {
			p = append(p, m.Sender)
		}
	}

	return
}

func (p Participants) includes(name string) bool {
	for _, n := range p {
		if n == name {
			return true
		}
	}
	return false
}
