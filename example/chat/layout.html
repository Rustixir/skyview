package Admin

import (
	Base "BuyCoffee/cmd"
	"bytes"
	"log"
	"os"
	"slices"
	"strconv"
	"sync/atomic"
	"time"
)

type Tracker struct {
	// number of streamers or youtubers
	Users atomic.Int64
	// number of payers
	Payers atomic.Int64
	// number of requests
	Requests atomic.Int64
	// number of views
	Views atomic.Int64

	interval time.Duration
	maxFiles int
}

func NewTracker(persistent bool, interval time.Duration, maxFiles int) *Tracker {
	tracker := &Tracker{
		interval: interval,
		maxFiles: maxFiles,
	}

	if persistent {
		tracker.startReporter()
	}
	return tracker
}

func (t *Tracker) IncUsers() {
	t.Users.Add(1)
}
func (t *Tracker) DecUsers() {
	t.Users.Add(-1)
}

func (t *Tracker) IncPayers() {
	t.Payers.Add(1)
}
func (t *Tracker) DecPayers() {
	t.Payers.Add(-1)
}

func (t *Tracker) IncRequests() {
	t.Requests.Add(1)
}
func (t *Tracker) DecRequests() {
	t.Requests.Add(-1)
}

func (t *Tracker) IncViews() {
	t.Views.Add(1)
}
func (t *Tracker) DecViews() {
	t.Views.Add(-1)
}

func (t *Tracker) startReporter() {
	ticker := time.Tick(t.interval)
	go func() {
		for range ticker {
			t.process()
		}
	}()
}

func (t *Tracker) process() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Receovered Panic!! Reason: %v", r)
		}
	}()

	path := Base.ReportPath
	file := path + "/report-" + time.Now().Format(time.DateOnly)
	os.WriteFile(file, t.report(), os.ModeAppend)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("Reporter error: " + err.Error())
		continue
	}
	if len(files) >= t.maxFiles {
		slices.SortFunc(files, func(a, b os.DirEntry) int {
			infoA, _ := a.Info()
			infob, _ := b.Info()
			return infoA.ModTime().Compare(infob.ModTime())
		})
		os.Remove(path + "/" + files[0].Name())
	}
}

func (t *Tracker) report() []byte {
	var (
		buf bytes.Buffer
	)

	users := t.Users.Swap(0)
	payers := t.Payers.Swap(0)
	requests := t.Requests.Swap(0)
	views := t.Views.Swap(0)

	buf.WriteString("Users: " + strconv.FormatInt(users, 10))
	buf.WriteString("Payers: " + strconv.FormatInt(payers, 10))
	buf.WriteString("Requests: " + strconv.FormatInt(requests, 10))
	buf.WriteString("Views: " + strconv.FormatInt(views, 10))
	return buf.Bytes()
}
