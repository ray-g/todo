package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rmDaysAgo bool
	rmRange   bool
	rmAll     bool
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove an existing todo",
	Long:    `Remove will remove an existing todo item from the list`,
	Run:     removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {
	items, err := todolist.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalln("failed to open datafile:", viper.GetString("datafile"))
	}

	if rmDaysAgo {
		removeDaysAgo(args[0], items)
	} else if rmAll {
		removeAll(items)
	} else {
		removeByIdx(args[0], items)
	}
}

func removeDaysAgo(dayStr string, items []todolist.Item) {
	days, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Fatalln(dayStr, "is not a valid number\n", err)
	}

	tm := time.Unix(time.Now().Unix()-60*60*24*int64(days), 0)

	length := len(items)
	for i := 0; i < length; i++ {
		item := items[i]
		if item.Done && item.TimeDone.Before(tm) {
			text := item.Text
			items = append(items[:i], items[i+1:]...)
			fmt.Printf("%q, %v\n", text, "removed")
			i--
			length--
		}
	}
	todolist.SaveItems(viper.GetString("datafile"), items)
}

func removeAll(items []todolist.Item) {
	length := len(items)
	for i := 0; i < length; i++ {
		item := items[i]
		if item.Done {
			text := item.Text
			items = append(items[:i], items[i+1:]...)
			fmt.Printf("%q, %v\n", text, "removed")
			i--
			length--
		}
	}
	todolist.SaveItems(viper.GetString("datafile"), items)
}

func removeByIdx(idx string, items []todolist.Item) {
	i, err := strconv.Atoi(idx)
	if err != nil {
		log.Fatalln(idx, "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		text := items[i-1].Text
		items = append(items[:i-1], items[i:]...)
		fmt.Printf("%q %v\n", text, "removed")

		todolist.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVar(&rmDaysAgo, "days-ago", false, "Remove done todos which X days ago, by done time.")
	// removeCmd.Flags().BoolVar(&rmRange, "range", false, "Remove done todos between the number range.")
	removeCmd.Flags().BoolVar(&rmAll, "all-done", false, "Remove all done todos")
}
