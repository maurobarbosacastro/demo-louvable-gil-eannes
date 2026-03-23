package jobs

import (
	"fmt"
	"ms-firebase-go/internal/service"
	"ms-firebase-go/pkg/dotenv"
	"ms-firebase-go/pkg/logster"
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

	go sendScheduledNotifications(loc)
}

func sendScheduledNotifications(location *time.Location) {
	timer := dotenv.GetEnv("SEND_SCHEDULED_NOTIFICATIONS_JOB")
	logster.StartFuncLogMsg(timer)

	// c := cron.NewWithLocation(location)
	c := cron.New()

	// Error handling for cron job addition
	err := c.AddFunc(timer, func() {
		logster.Info("sendScheduledNotificationsJob Started")

		service.ProcessCurrentNotifications()

		logster.Info("sendScheduledNotificationsJob Ended")
	})
	if err != nil {
		logster.Error(err, "Error adding cron job sendScheduledNotificationsJob")
		return
	}

	// Start the cron scheduler
	c.Start()

	select {}
}
