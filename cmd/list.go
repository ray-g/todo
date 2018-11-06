package cmd

import (
	"log"

	"github.com/ray-g/todo/todolist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the todos",
	Long:    `Listing the todos`,
	Run:     listRun,
}

var printOpt todolist.PrintOption

func listRun(cmd *cobra.Command, args []string) {
	items, err := todolist.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	// sort.Sort(todolist.ByPri(items))

	todolist.PrettyPrint(items, printOpt)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&printOpt.Done, "done", "n", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVarP(&printOpt.All, "all", "a", false, "Show all Todos")
	listCmd.Flags().BoolVarP(&printOpt.LastWeek, "lastweek", "w", false, "Show 'Done' Todos in last week")
	listCmd.Flags().BoolVarP(&printOpt.LastMonth, "lastmonth", "m", false, "Show 'Done' Todos in last month")
	listCmd.Flags().Int64VarP(&printOpt.Days, "days", "d", 0, "Show 'Done' Todos in last n days")
	listCmd.Flags().BoolVarP(&printOpt.LongTime, "longtime", "l", false, "Show time in long format (including clock)")
}
