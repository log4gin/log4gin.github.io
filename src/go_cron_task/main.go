package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New(
		cron.WithSeconds(),
	)

	c.AddFunc("@every 2s", func() {
		cmd := exec.Command("python", "main.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}
	})

	c.Start()

	select {}

}
