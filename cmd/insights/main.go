package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/umayr/insights"
	"github.com/umayr/insights/server"
)

func main() {
	var (
		flagTimezone = flag.String("timezone", "", "Timezone of the exported chat like Asia/Dubai")
		flagServer   = flag.Bool("server", false, "Serve the insights in graphical representation")
		flagPretty   = flag.Bool("pretty", true, "Display the JSON in pretty format")
		flagPort     = flag.String("port", "9000", "Port for the server")
	)

	flag.Parse()

	src := flag.Arg(0)
	if src == "" {
		fmt.Println("error: path to exported chat is empty")
		os.Exit(1)
	}

	f, err := os.Open(src)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	c, err := insights.NewConversation(string(buf), *flagTimezone)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	i := insights.New(c)

	if *flagServer {
		s := server.New(i)
		if err := s.Start(*flagPort); err != nil {
			fmt.Printf("error: %s\n", err.Error())
			os.Exit(1)
		}
		return
	}

	var b []byte
	if *flagPretty {
		b, err = json.MarshalIndent(i, "", "  ")
	} else {
		b, err = json.Marshal(i)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	os.Stdout.Write(b)
}
