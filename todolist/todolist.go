package todolist

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

// Item for todo
type Item struct {
	Text     string
	Priority int
	Done     bool
	TimeAdd  time.Time
	TimeDone time.Time
	position int
}

// SetPriority sets the priority of todo item
func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

// SetTime set the time of todo item
// `done` true  - set TimeDone
//        false - set TimeAdd
func (i *Item) SetTime(done bool) {
	tm := time.Now()
	if done {
		i.TimeDone = tm
	} else {
		i.TimeAdd = tm
	}
}

// SaveItems save the todo items to the datafile
func SaveItems(filename string, items []Item) error {
	sort.Sort(ByPri(items))

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReadItems reads the todo items from the datafile
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

// ByPri implements sort.Interface for []Item based on the Priority & position field.
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].position < s[j].position
		}
		return s[i].Priority < s[j].Priority
	}
	return !s[i].Done
}

// ByTimeAdd implements sort.Interface for []Item based on add time
type ByTimeAdd []Item

func (s ByTimeAdd) Len() int      { return len(s) }
func (s ByTimeAdd) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTimeAdd) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].TimeAdd.After(s[j].TimeAdd)
		}
		return s[i].Priority < s[j].Priority
	}
	return !s[i].Done
}

// ByTimeDone implements sort.Interface for []Item based on done time
type ByTimeDone []Item

func (s ByTimeDone) Len() int      { return len(s) }
func (s ByTimeDone) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTimeDone) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].TimeDone.After(s[j].TimeDone)
		}
		return s[i].Priority < s[j].Priority
	}
	return !s[i].Done
}
