package jobs

import (
	"fmt"
	"github.com/robfig/cron"
	"ms-cj/internal/services"
	"ms-cj/pkg/dotenv"
	"ms-cj/pkg/logster"
	"time"
)

func StartJobs() {
	logster.Info("Starting jobs...")

	loc, err := time.LoadLocation("Europe/Lisbon")
	if err != nil {
		logster.Error(err, "Error loading timezone")
	}
	logster.Info(fmt.Sprintf("Timezone: %s", loc.String()))

	go checkCjTransactions(loc)
}

func checkCjTransactions(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("CJ_TRANSACTIONS_CHECK_TIMER")
	logster.Info("Timer: " + timer)

	c := cron.NewWithLocation(location)

	err := c.AddFunc(timer, func() {
		logster.Info("checkCjTransactions Start")

		services.ProcessTransactions()

		logster.Info("checkCjTransactions end")
	})

	if err != nil {
		logster.Error(err, "Error adding cron job checkCjTransactions")
		return
	}

	c.Start()
	select {}
}
