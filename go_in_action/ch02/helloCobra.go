package main

import (
	"github.com/spf13/cobra"
)

var helloCommand *cobra.Command

func init() {
	helloCommand = &cobra.Command{
		Use:   "hello",
		Short: "Print hello world",
		Run:   sayHello,
	}
	helloCommand.Flags().StringP("name", "n", "World", "Who to say hello to.")
	helloCommand.MarkFlagRequired("name")
	helloCommand.Flags().StringP("language", "l", "en", "Which language to say hello in.")
}

func sayHello(cmd *cobra.Command, args []string) {

}
