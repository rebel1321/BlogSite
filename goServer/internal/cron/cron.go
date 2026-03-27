package cron

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func StartCronJobs() {
	c := cron.New()

	// Every 10 minutes
	_, err := c.AddFunc("*/10 * * * *", func() {
		fmt.Println("⏰ Running every 10 minutes: Ping 🚀")
		
		// 👉 your logic here
		// example:
		// pingServer()
	})

	if err != nil {
		log.Println("Cron error:", err)
	}

	c.Start()

	fmt.Println("✅ Cron jobs started")
}