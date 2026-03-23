package service

import (
	"fmt"
	"ms-tagpeak/external/notifications"
	"ms-tagpeak/pkg/logster"

	"github.com/samber/lo"
)

func UpdateUserTopic(userUuid string, level string) {
	logster.StartFuncLogMsg(fmt.Sprintf("user uuid: %s, level: %s", userUuid, level))

	topics, errGetTopics := notifications.GetAllTopics()

	if errGetTopics != nil {
		logster.Error(errGetTopics, "Error getting topics")
	}

	topicBase, _ := lo.Find(topics, func(item notifications.Topic) bool {
		return item.Name == *(Configuration.MembershipLevels.Base)
	})
	topicSilver, _ := lo.Find(topics, func(item notifications.Topic) bool {
		return item.Name == *(Configuration.MembershipLevels.Silver)
	})
	topicToSet, _ := lo.Find(topics, func(item notifications.Topic) bool {
		return item.Name == level
	})

	if level == *(Configuration.MembershipLevels.Silver) {
		errRemove := notifications.RemoveUserFromTopic(topicBase.UUID.String(), userUuid)
		if errRemove != nil {
			logster.Error(errRemove, "Error removing token from topic")
		}
	}

	if level == *(Configuration.MembershipLevels.Gold) {
		errRemove := notifications.RemoveUserFromTopic(topicSilver.UUID.String(), userUuid)
		if errRemove != nil {
			logster.Error(errRemove, "Error removing token from topic")
		}
	}

	errAdd := notifications.AddUserToTopic(topicToSet.UUID.String(), userUuid)
	if errAdd != nil {
		logster.Error(errAdd, "Error adding token to topic")
	}
}
