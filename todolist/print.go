package todolist

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

// PrintOption for print to console, used in list command
type PrintOption struct {
	All       bool
	Done      bool
	LastWeek  bool
	LastMonth bool
	LongTime  bool
	Days      int64
}

var timeFormat = ""

func initFormat(long bool) string {
	timeFormat = "[2006-01-02]"
	if long {
		timeFormat = "[2006-01-02 15:04:05]"
	}
	return timeFormat
}

// PrettyPrint prints todo items prettily
func PrettyPrint(items []Item, opt PrintOption) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	initFormat(opt.LongTime)

	compareDays := false
	days := opt.Days
	if opt.LastWeek {
		days = 7
	}
	if opt.LastMonth {
		days = 31
	}
	if days > 0 {
		compareDays = true
	}
	tm := time.Unix(time.Now().Unix()-60*60*24*days, 0)

	for _, i := range items {
		if opt.All || i.Done == opt.Done ||
			(compareDays && i.Done && i.TimeDone.After(tm)) {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", i.label(), i.prettyTimeAdd(), i.prettyDone(), i.prettyPri(), i.Text)
		}
	}

	w.Flush()
}

func (i *Item) label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) prettyDone() string {
	if i.Done {
		return "- " + i.TimeDone.Format(timeFormat)
	}
	return ""
}

func (i *Item) prettyPri() string {
	pri := ""
	switch i.Priority {
	case 1:
		pri = "High"
	case 3:
		pri = "Low"
	default:
		pri = "Norm"
	}

	return fmt.Sprintf("[%s]", pri)
}

func (i *Item) prettyTimeAdd() string {
	return i.TimeAdd.Format(timeFormat)
}
