package notifications

import "fmt"

type NotificationContent struct {
	Title string
	Body  string
}

// GetNotificationContentWithValues returns notification content with values replaced using fmt.Sprintf
// The values array should contain values in the order they appear in the template:
// - LIVE: [refId, store]
// - TRACKED: [store]
// - FINISHED: [store, refId, value]
// - STOPPED: []
func GetNotificationContent(cashbackState string, values ...interface{}) *NotificationContent {
	switch cashbackState {
	case "LIVE":
		return &NotificationContent{
			Title: "Your cashback is live!",
			Body:  fmt.Sprintf("We've just launched your cashback (ref ID %s) on your purchase at %s. You may follow its performance daily from your dashboard and redeem it at any time.", values...),
		}
	case "TRACKED":
		return &NotificationContent{
			Title: "Purchase tracked!",
			Body:  fmt.Sprintf("Congratulations on your purchase at %s. We are now waiting for confirmation from the partner store before launching your cash reward.", values...),
		}
	case "FINISHED":
		return &NotificationContent{
			Title: "",
			Body:  fmt.Sprintf("Your purchase on %s with reference %s has produced a final cashback of %s. This has been added to your balance and is available for withdrawal.", values...),
		}
	case "STOPPED":
		return &NotificationContent{
			Title: "Stop Request received",
			Body:  "We will confirm the final cash reward amount in the next 2 business days.",
		}
	}
	return nil
}
