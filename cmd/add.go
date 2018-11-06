package cmd

import (
	"fmt"
	"log"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"ad"},
	Short:   "Add a new todo",
	Long:    `Add will create a new todo item to the list`,
	Run:     addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todolist.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todolist.Item{Text: x}
		item.SetPriority(priority)

		item.SetTime(false)

		items = append(items, item)
	}

	if err := todolist.SaveItems(viper.GetString("datafile"), items); err != nil {
		fmt.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3, (1-High, 2-Norm, 3-Low)")
}
