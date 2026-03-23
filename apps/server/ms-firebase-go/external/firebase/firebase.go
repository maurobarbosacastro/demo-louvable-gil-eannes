package firebase

import (
	"context"
	"fmt"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/responses"
	"ms-firebase-go/pkg/dotenv"
	"ms-firebase-go/pkg/logster"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var Client *messaging.Client

func Init() {
	var err error
	Client, err = initFirebase()
	if err != nil {
		logster.Error(err, "Failed to initialize Firebase")
	}
}

func initFirebase() (*messaging.Client, error) {
	logster.StartFuncLog()
	ctx := context.Background()

	pathFile := dotenv.GetEnv("FIREBASE_CREDENTIALS_PATH")

	// Check if credentials file exists
	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("firebase credentials file not found at path %s: %w", pathFile, err)
	}

	opt := option.WithCredentialsFile(pathFile)
	config := &firebase.Config{ProjectID: dotenv.GetEnv("PROJECT_ID")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting messaging client: %w", err)
	}

	logster.EndFuncLogMsg("Connected to Firebase")
	return client, nil
}

func SendNotification(ctx context.Context, token string, messageDto *dto.MessageDTO) (*responses.NotificationResponse, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("%+v", *messageDto))

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: messageDto.Title,
			Body:  messageDto.Body,
		},
		Data: map[string]string{
			"message": messageDto.Body,
		},
		Token: token,
	}

	messageID, err := Client.Send(ctx, message)
	if err != nil {
		logster.Error(err, "Error sending notification")
		return &responses.NotificationResponse{
			Success:   false,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Message sent with ID: %s", messageID))
	return &responses.NotificationResponse{
		Success:   true,
		MessageID: messageID,
		Timestamp: time.Now(),
	}, nil
}

func SendGroupMessage(ctx context.Context, group string, messageDto *dto.MessageDTO) (*responses.NotificationResponse, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("%+v", *messageDto))

	topic := fmt.Sprintf("/topics/%s", group)
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: messageDto.Title,
			Body:  messageDto.Body,
		},
		Data: map[string]string{
			"message": messageDto.Body,
		},
		Topic: topic,
	}

	messageID, err := Client.Send(ctx, message)
	if err != nil {
		logster.Error(err, "Error sending group notification")
		return &responses.NotificationResponse{
			Success:   false,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}, err
	}

	logster.Info(fmt.Sprintf("Group message sent to topic %s with ID: %s", topic, messageID))

	logster.EndFuncLog()
	return &responses.NotificationResponse{
		Success:   true,
		MessageID: messageID,
		Timestamp: time.Now(),
	}, nil
}

func SubscribeTokensToTopic(ctx context.Context, topic string, tokens []string) {
	logster.StartFuncLogMsg(fmt.Sprintf("topic: %s | tokens: %v", topic, tokens))

	// topic = "/topics/" + topic
	response, err := Client.SubscribeToTopic(ctx, tokens, topic)
	if err != nil {
		logster.Panic(err, "Error subscribing to topic")
	}
	logster.Warnf("%d tokens were subscribed to %s", response.SuccessCount, topic)
	logster.Warnf("%d tokens failed to subscribe to %s", response.FailureCount, topic)
	if response.FailureCount > 0 {
		for i, err := range response.Errors {
			if err != nil {
				error := fmt.Errorf("%+v", *err)
				logster.Error(error, fmt.Sprintf("Error %d", i))
			}
		}
	}
}

func UnsubscribeTokensFromTopic(ctx context.Context, topic string, tokens []string) {
	logster.StartFuncLogMsg(fmt.Sprintf("topic: %s | tokens: %v", topic, tokens))

	response, err := Client.UnsubscribeFromTopic(ctx, tokens, topic)
	if err != nil {
		logster.Panic(err, "Error subscribing to topic")
	}
	logster.Warnf("%d tokens were unsubscribed to %s", response.SuccessCount, topic)
	logster.Warnf("%d tokens failed to unsubscribe to %s", response.FailureCount, topic)
	if response.FailureCount > 0 {
		for i, err := range response.Errors {
			if err != nil {
				error := fmt.Errorf("%+v", *err)
				logster.Error(error, fmt.Sprintf("Error %d", i))
			}
		}
	}
}
