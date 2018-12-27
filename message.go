package insights

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const regexMetaInfo = `^(\[.*\])\s([a-zA-Z\s]+)\:`
const regexASCII = `[[:^ascii:]]`

type Type int

const (
	MessageTypeText Type = iota
	MessageTypeVideo
	MessageTypeAudio
	MessageTypeImage
	MessageTypeContact
)

func (t Type) MarshalJSON() (b []byte, err error) {
	var dict = map[Type]string{
		MessageTypeText:    "text",
		MessageTypeVideo:   "video",
		MessageTypeAudio:   "audio",
		MessageTypeImage:   "image",
		MessageTypeContact: "contact",
	}

	if v, exists := dict[t]; exists {
		b = []byte(fmt.Sprintf(`"%s"`, v))
	} else {
		err = fmt.Errorf("unable to marshal type: invalid type %d", t)
	}
	return
}

func (t *Type) UnmarshalJSON(b []byte) (err error) {
	var dict = map[string]Type{
		`"text"`:    MessageTypeText,
		`"video"`:   MessageTypeVideo,
		`"audio"`:   MessageTypeAudio,
		`"image"`:   MessageTypeImage,
		`"contact"`: MessageTypeContact,
	}

	if v, exists := dict[string(b)]; exists {
		*t = v
	} else {
		err = fmt.Errorf("unable to unmarshal type: invalid type: %s", b)
	}

	return
}

type Message struct {
	Sender string    `json:"sender"`
	Time   time.Time `json:"time"`
	Text   string    `json:"text"`
	Type   Type      `json:"type"`

	rawText string
}

var ErrInvalidMessage = fmt.Errorf("invalid message")

func NewMessage(raw string, loc *time.Location) (*Message, error) {
	m := &Message{}
	m.rawText = regexp.MustCompile(regexASCII).ReplaceAllLiteralString(raw, "")

	if strings.Contains(raw, "Messages to this group are now secured with end-to-end encryption.") {
		return nil, ErrInvalidMessage
	}

	r := regexp.MustCompile(regexMetaInfo)
	t := r.Split(raw, -1)
	if len(t) < 2 {
		return nil, ErrInvalidMessage
	}

	msg := t[1]

	if strings.Contains(msg, "omitted") {
		if strings.Contains(msg, "image") {
			m.Type = MessageTypeImage
		}
		if strings.Contains(msg, "audio") {
			m.Type = MessageTypeAudio
		}
		if strings.Contains(msg, "video") {
			m.Type = MessageTypeVideo
		}
		if strings.Contains(msg, "card") {
			m.Type = MessageTypeContact
		}
	} else {
		m.Type = MessageTypeText
	}

	m.Text = strings.TrimSpace(msg)
	g := r.FindStringSubmatch(raw)
	g1, g2 := g[1], g[2]

	m.Sender = strings.TrimSpace(g2)

	g1 = strings.TrimPrefix(g1, "[")
	g1 = strings.TrimSuffix(g1, "]")
	g1 = strings.TrimSpace(g1)

	tm, err := time.Parse("1/2/06, 3:04:05 PM", g1)
	if err != nil {
		return nil, err
	}

	m.Time = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(), tm.Nanosecond(), loc)

	return m, nil
}
