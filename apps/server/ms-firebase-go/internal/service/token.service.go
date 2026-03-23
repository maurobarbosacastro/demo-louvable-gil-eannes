package service

import (
	"context"
	"fmt"
	"ms-firebase-go/external/firebase"
	"ms-firebase-go/internal/dto"
	"ms-firebase-go/internal/models"
	"ms-firebase-go/internal/repository"
	"ms-firebase-go/pkg/logster"
)

func CreateToken(createDto dto.TokenSave) (*models.UserToken, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("%+v", createDto))

	token, err := repository.CreateUserToken(models.UserToken{
		UserUUID: createDto.UserUUID,
		Token:    createDto.Token,
	})
	if err != nil {
		logster.Error(err, "Error creating token")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return token, nil
}

func RemoveToken(token string) error {
	logster.StartFuncLogMsg(token)

	// ToDO: Remove all tokens.
	userToken, errUserToken := repository.GetUserTokenByToken(token)
	if errUserToken != nil {
		logster.Error(errUserToken, "Error getting user token entity")
		logster.EndFuncLog()
		return errUserToken
	}

	errRemoveTopics := removeUserTokenFromTopics(*userToken)
	if errRemoveTopics != nil {
		logster.Error(errRemoveTopics, "Error removing topics associated with token")
		logster.EndFuncLog()
		return errRemoveTopics
	}

	return repository.DeleteUserToken(userToken.UUID)
}

func RemoveTokenForUser(userUuid string) error {
	logster.StartFuncLogMsg(userUuid)

	userToken, errUserToken := repository.GetUserTokenByUserUuid(userUuid)
	if errUserToken != nil {
		logster.Error(errUserToken, "Error getting user token entity")
		logster.EndFuncLog()
		return errUserToken
	}

	errRemoveTopics := removeUserTokenFromTopics(*userToken)
	if errRemoveTopics != nil {
		logster.Error(errRemoveTopics, "Error removing topics associated with token")
		logster.EndFuncLog()
		return errRemoveTopics
	}

	return repository.DeleteUserToken(userToken.UUID)
}

func removeUserTokenFromTopics(user models.UserToken) error {
	topics, errRemoveTopics := repository.RemoveUserTokenFromAllTopics(user)

	if errRemoveTopics != nil {
		logster.Error(errRemoveTopics, "Error removing topics associated with token")
		logster.EndFuncLog()
		return errRemoveTopics
	}

	for _, topic := range *topics {
		firebase.UnsubscribeTokensFromTopic(context.Background(), topic.UUID.String(), []string{user.Token})
	}

	return nil
}

func GetUserTokenFromUserUuid(userUuid string) (*[]models.UserToken, error) {
	logster.StartFuncLogMsg(userUuid)

	userToken, err := repository.GetAllUserTokensByUser(userUuid)
	if err != nil {
		logster.Error(err, "Error while getting userToken entity")
		logster.EndFuncLog()
		return nil, err
	}

	return userToken, nil
}
