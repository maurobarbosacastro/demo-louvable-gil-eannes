package jobs

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func StartJobs() {
	logster.Info("Starting jobs...")

	loc, err := time.LoadLocation("Europe/Lisbon")
	if err != nil {
		logster.Error(err, "Error loading timezone")
	}
	logster.Info(fmt.Sprintf("Timezone: %s", loc.String()))

	go example(loc)
}

func example(location *time.Location) {
	logster.StartFuncLog()

	timer := dotenv.GetEnv("EXAMPLE_JOB")

	c := cron.NewWithLocation(location)

	// Error handling for cron job addition
	err := c.AddFunc(timer, func() {
		logster.Info("exampleJob Started")

		logster.Info("exampleJob Ended")
	})
	if err != nil {
		logster.Error(err, "Error adding cron job exampleJob")
		return
	}

	// Start the cron scheduler
	c.Start()

	select {}
}
