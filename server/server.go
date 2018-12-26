package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/umayr/insights"
)

type Section struct {
	Heading string
	Values  map[string]interface{}
}

type Payload struct {
	Title    string
	Sections []Section
	JSON     string
}

type Server struct {
	Data    *insights.Insights
	payload *Payload
}

func New(i *insights.Insights) *Server {
	return &Server{Data: i}
}

func (s *Server) process() error {
	b, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}

	f := "Monday, 02 Jan 2006"
	p := &Payload{}
	p.JSON = string(b)

	sec := make([]Section, 3)

	s0 := Section{Heading: "General"}
	m0 := make(map[string]interface{})
	m0["Total Messages"] = s.Data.TotalMessages
	m0["Total Words"] = s.Data.TotalWords
	m0["Total Letters"] = s.Data.TotalLetters
	s0.Values = m0
	sec[0] = s0

	s1 := Section{Heading: "Activity"}
	m1 := make(map[string]interface{})
	m1["Most Active Day"] = s.Data.MostActiveDay.Format(f)
	m1["Most Active Count"] = s.Data.MostActiveCount
	m1["Least Active Day"] = s.Data.LeastActiveDay.Format(f)
	m1["Least Active Count"] = s.Data.LeastActiveCount
	s1.Values = m1
	sec[1] = s1

	s2 := Section{Heading: "Average"}
	m2 := make(map[string]interface{})
	m2["Words Per Message"] = fmt.Sprintf("%.2f", s.Data.AverageWordsPerMessage)
	m2["Letters Per Message"] = fmt.Sprintf("%.2f", s.Data.AverageLettersPerMessage)
	m2["Messages Per Day"] = fmt.Sprintf("%.2f", s.Data.AverageMessagesPerDay)
	m2["Words Per Day"] = fmt.Sprintf("%.2f", s.Data.AverageWordsPerDay)
	m2["Letters Per Day"] = fmt.Sprintf("%.2f", s.Data.AverageLettersPerDay)
	s2.Values = m2
	sec[2] = s2

	p.Sections = sec
	p.Title = "Your Insights"

	s.payload = p
	return nil
}

func (s *Server) Start(port string) error {
	if err := s.process(); err != nil {
		return err
	}

	pubBox := packr.NewBox("./public")
	tmplBox := packr.NewBox("./tmpl")

	fmt.Printf("Insights are being served at http://localhost:%s\n", port)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(pubBox)))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		raw, err := tmplBox.FindString("index.html")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong"))
			return
		}

		tmpl := template.Must(template.New("").Parse(raw))
		tmpl.Execute(w, s.payload)
	})

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
