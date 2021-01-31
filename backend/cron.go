package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

type CronTab struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Recipient   string `json:"recipient"`
	Schedule    string `json:"schedule"`
	Type        string `json:"type"`
	Stopped     bool   `json:"stopped"`
}

func cronWatch() {
	w := watcher.New()
	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only notify rename and move events.
	w.FilterOps(watcher.Write)

	// Only files that match the regular expression during file listings
	// will be watched.
	r := regexp.MustCompile("^abc$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
				cronTabUpdater()
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("/crontabs/crontab.db"); err != nil {
		log.Fatalln(err)
	}

	// Trigger 2 events after watcher started.
	go func() {
		w.Wait()
		cronTabUpdater()
		c.Start()
	}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func cronTabUpdater() {
	f, err := os.Open("/crontabs/crontab.db")
	if err != nil {
		log.Fatal(err)
	}
	d := json.NewDecoder(f)
	for _, y := range c.Entries() {
		fmt.Printf("Removing Id: %s\n", y.Job)
		c.Remove(y.ID)
	}
	for {
		var conf CronTab
		if err := d.Decode(&conf); err == io.EOF {
			break // done decoding file
		} else if err != nil {
			log.Fatal(err)
		}
		//var output = fmt.Sprintf("There is a scheduled payment to %s of the amount %s. The current schedule is: %s", conf.Recipient, conf.Amount, conf.Schedule)
		if conf.Stopped == false {
			c.AddFunc(conf.Schedule, func() {
				venmoAction(conf.Recipient, conf.Description, conf.Type, conf.Amount)
			})
		}
	}
	for _, y := range c.Entries() {
		fmt.Printf("New Job Added:%s, Next Run:%s\n", y.Job, y.Next)
	}
}
