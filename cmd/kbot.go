/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var TeleToken = os.Getenv("TELE_TOKEN") // TeleToken should be set in the environment

// kbotCmd represents the kbot command

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kbot is running...", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "https://api.telegram.org",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Failed to create bot instance. Check your TELE_TOKEN.", err)
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {

			log.Printf("Payload: %s, Text: %s", m.Message().Payload, m.Message().Text)
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprint("Welcome to KBot! How can I assist you today?", appVersion))
			}
			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
