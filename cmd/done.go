package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as Done",
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todolist.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		items[i-1].SetTime(true)
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		todolist.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
