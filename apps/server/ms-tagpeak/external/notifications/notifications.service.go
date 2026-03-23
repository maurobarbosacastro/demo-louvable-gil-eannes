package notifications

import (
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/http_client"
	"net/http"
)

func SaveToken(body TokenSave) (*UserToken, error) {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "token"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	var response UserToken

	_, err := internalHttpClient.PostJson(url, nil, &body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func DeleteToken(token string) error {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "token/" + token

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	_, err := internalHttpClient.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTokensByUser(userUuid string) error {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "token/user/" + userUuid

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	_, err := internalHttpClient.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}

func CreateNotification(body CreateNotificationDto) (*Notification, error) {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "notification"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	var response Notification

	_, err := internalHttpClient.PostJson(url, nil, &body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func SendNotification(notificationUuid string) error {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "notification/" + notificationUuid + "/send"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	// Send empty body since the notification UUID is in the URL path
	_, err := internalHttpClient.PostJson(url, nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTopics() ([]Topic, error) {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "topic"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}
	var response []Topic

	_, err := internalHttpClient.Get(url, nil, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func AddUserToTopic(topicUuid string, userUuid string) error {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "topic/" + topicUuid + "/user"

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	_, err := internalHttpClient.PostJson(url, nil, &AddRemoveUserToTopicDto{
		UserUuid: userUuid,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func RemoveUserFromTopic(topicUuid string, userUuid string) error {
	baseUrl := dotenv.GetEnv("MS_FIREBASE_GO_URL")
	url := baseUrl + "topic/" + topicUuid + "/user/" + userUuid

	internalHttpClient := &http_client.InternalHttpClient{InternalHttpClient: &http.Client{}}

	_, err := internalHttpClient.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}
