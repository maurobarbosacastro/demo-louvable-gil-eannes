package service

import (
	"context"
	"fmt"
	"ms-firebase-go/external/firebase"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/models"
	"ms-firebase-go/internal/repository"
	"ms-firebase-go/pkg/logster"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func GetAllTopics() ([]models.TopicWithoutTarget, error) {
	topics, err := repository.GetAllTopics()
	if err != nil {
		return nil, err
	}

	topicsMapped := lo.Map(topics, func(item models.Topic, index int) models.TopicWithoutTarget {
		return item.ToTopicWithoutTarget()
	})

	return topicsMapped, nil
}

func CreateTopic(data dto.CreateTopicDto) (*models.Topic, error) {
	return repository.CreateTopic(models.Topic{Name: data.Name})
}

func GetTopicByName(name string) (*models.Topic, error) {
	return repository.GetTopicByName(name)
}

func AddTokenToTopic(topicUuid uuid.UUID, token string) error {
	firebase.SubscribeTokensToTopic(context.Background(), topicUuid.String(), []string{token})

	return AddTokensToTopic(topicUuid, []string{token})
}

func RemoveTokenFromTopic(topicUuid uuid.UUID, token string) error {
	firebase.UnsubscribeTokensFromTopic(context.Background(), topicUuid.String(), []string{token})

	return RemoveTokensFromTopic(topicUuid, []string{token})
}

func AddTokensToTopic(topicUuid uuid.UUID, tokens []string) error {
	userTokens, err := repository.GetUserTokensByTokens(tokens)
	if err != nil {
		return err
	}

	return repository.AddUserTokensToTopic(topicUuid, userTokens)
}

func RemoveTokensFromTopic(topicUuid uuid.UUID, tokens []string) error {
	userTokens, err := repository.GetUserTokensByTokens(tokens)
	if err != nil {
		return err
	}

	return repository.RemoveUserTokensFromTopic(topicUuid, userTokens)
}

func AddUserToTopic(topicUuid uuid.UUID, user string) error {
	logster.StartFuncLogMsg(fmt.Sprintf("%s, %s", topicUuid, user))

	err := repository.AddUserToTopic(user, topicUuid, user)
	if err != nil {
		logster.Error(err, "Error adding user to topic")
		logster.EndFuncLog()
		return err
	}

	userToken, errGetUserToken := GetUserTokenFromUserUuid(user)
	if errGetUserToken != nil {
		logster.Error(errGetUserToken, "Error getting user token")
		logster.EndFuncLog()
		return err
	}

	logster.Info(fmt.Sprintf("User token: %+v", *userToken))

	tokens := lo.Map(*userToken, func(item models.UserToken, index int) string {
		return item.Token
	})

	firebase.SubscribeTokensToTopic(context.Background(), topicUuid.String(), tokens)

	logster.EndFuncLog()
	return nil
}

func RemoveUserFromTopic(topicUuid uuid.UUID, user string) error {
	err := repository.RemoveUserFromTopic(user, topicUuid)

	if err != nil {
		return err
	}

	userToken, errGetUserToken := GetUserTokenFromUserUuid(user)
	if errGetUserToken != nil {
		return err
	}
	tokens := lo.Map(*userToken, func(item models.UserToken, index int) string {
		return item.Token
	})
	firebase.UnsubscribeTokensFromTopic(context.Background(), topicUuid.String(), tokens)

	return nil
}
