package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/martinlindhe/notify"
)

func main() {
	messagePtr := flag.String("message", "I think you needed to remember something, mate!", "a string")
	cronPtr := flag.String("cron", "empty", "a string")
	flag.Parse()

	if *cronPtr == "empty" {
		notify.Notify("golarm", "Thing to remember", *messagePtr, "path/to/icon.png")
	} else {
		maybeInitCrontab()
		saveCronJob(*cronPtr, *messagePtr)
	}
}

// Creates a crontab in the scenarion in which the user doesn't have an existing one
func maybeInitCrontab() {
	cmd := exec.Command("sh", "-c", "crontab -l 2>/dev/null | crontab -")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func saveCronJob(cron string, message string) {
	command := fmt.Sprintf("(crontab -l 2>/dev/null && echo \"%s /Users/manzanit0/repositories/golarm/golarm -message=%s\") | crontab -", cron, message)
	cmd := exec.Command("sh", "-c", command)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
