package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Lists struct {
	Do        string
	Added     string
	Completed string
	TC        string
}

var Todo []Lists
var add = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Adds a thing tothe list",
	Run: func(cmd *cobra.Command, args []string) {
		to := Lists{args[0], (time.Now().Format(time.DateTime)), "No", "nil"}
		Todo = append(Todo, to)
		fmt.Println("Added: ", args[0])

		// fmt.Println(Todo)
	},
}

var list = &cobra.Command{
	Use:   "list",
	Short: "Lists all the items currently existing in the system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Id| Job\t\t| Entry Time \t\t| Done\t| Completed Time       |")
		fmt.Println(`------------------------------------------------------------------------`)
		for i, V := range Todo {
			fmt.Println(i+1, "| ", V.Do, "\t| ", V.Added, "\t| ", V.Completed, "\t| ", V.TC, "|")
		}
	},
}

var done = &cobra.Command{
	Use:   "done [done element]",
	Short: "Marks one duty as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i, V := range Todo {
			if string(args[0]) == V.Do {
				Todo[i].TC = string(time.Now().Format(time.DateTime))
				Todo[i].Completed = "Yes"
				fmt.Println("Done: ", args[0])

			}
		}
	},
}

var remove = &cobra.Command{
	Use:   "remove [Item]",
	Short: "Removes an item from the list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i, V := range Todo {
			if V.Do == string(args[0]) {
				fmt.Println("Removed: ", args[0])
				Todo = append(Todo[:i], Todo[i+1:]...)
			}
		}
	},
}

func makeroot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ToDo",
		Short: "For adding your To-Do List ",
	}
	rootCmd.AddCommand(add, list, done, remove)
	return rootCmd

}

func execute(cmd *cobra.Command) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		check(err)
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Nothing inputed")
			continue
		}
		if input == "quit" {
			fmt.Println("Exiting!!!!")
			break
		}
		args := strings.Fields(input)

		cmd.SetArgs(args)

		if err := cmd.Execute(); err != nil {
			fmt.Println("Error")
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
func main() {
	rootCmd := makeroot()
	execute(rootCmd)
}
