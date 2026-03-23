package dto

type CreateTopicDto struct {
	Name string `json:"name"`
}

type AddRemoveTokenToTopicDto struct {
	Token string `json:"token"`
}

type AddRemoveUserToTopicDto struct {
	UserUUID string `json:"userUuid"`
}
