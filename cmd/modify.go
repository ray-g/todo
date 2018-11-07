package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var modifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"mod"},
	Short:   "Modify an existing Item's Text and Priority",
	Run:     modifyRun,
}

var (
	modPriority int
	modText     string
)

func modifyRun(cmd *cobra.Command, args []string) {
	items, err := todolist.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		if modText != "" {
			oldText := items[i-1].Text
			items[i-1].Text = modText
			fmt.Printf("'%v' modified to '%v'\n", oldText, items[i-1].Text)
		}

		if items[i-1].Priority != modPriority {
			oldPriority := items[i-1].Priority
			items[i-1].Priority = modPriority
			fmt.Printf("'%v' priority modified from [%d] to [%d]\n", items[i-1].Text, oldPriority, modPriority)
		}

		todolist.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(modifyCmd)

	modifyCmd.Flags().IntVarP(&modPriority, "priority", "p", 2, "Priority:1,2,3, (1-High, 2-Norm, 3-Low)")
	modifyCmd.Flags().StringVarP(&modText, "text", "t", "", "New todo details")
}
