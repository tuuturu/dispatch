package cmd

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"github.com/tuuturu/dispatch/pkg/handler/tick"
	"github.com/tuuturu/dispatch/pkg/watcher/mqtt"
	"log"
	"os"
	"time"
)

var rootCmd = cobra.Command{
	Use:  "dispatch",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		ctx := &daemon.Context{
			PidFileName: "dispatcher.pid",
			PidFilePerm: 0644,
			LogFileName: "dispatcher.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
		}

		d, err := ctx.Reborn()
		if err != nil {
			return err
		}

		if d != nil {
			return
		}

		defer func() {
			_ = ctx.Release()
		}()

		log.Print("- - - - - - - - - - - - - - - -")
		log.Print("daemon started")

		err = runDispatcher(args)
		if err != nil {
			return fmt.Errorf("starting dispatcher: %w", err)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

func runDispatcher(args []string) error {
	watcher := mqtt.NewMQTTWatcher(args[0])

	err := watcher.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = watcher.Close()
	}()

	watcher.RegisterHandler(tick.NewTickHandler())

	for {
		time.Sleep(2 * time.Second)
	}

}
