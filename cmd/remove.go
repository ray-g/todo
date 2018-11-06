package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
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
}
