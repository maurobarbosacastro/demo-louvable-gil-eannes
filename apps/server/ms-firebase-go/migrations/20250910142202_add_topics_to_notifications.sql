-- +goose Up
-- +goose StatementBegin
CREATE TABLE notification_user_tokens (
    notification_uuid UUID NOT NULL,
    user_token_uuid UUID NOT NULL,
    PRIMARY KEY (notification_uuid, user_token_uuid),
    FOREIGN KEY (notification_uuid) REFERENCES notification(uuid) ON DELETE CASCADE,
    FOREIGN KEY (user_token_uuid) REFERENCES user_token(uuid) ON DELETE CASCADE
);

CREATE INDEX idx_notification_user_tokens_notification_uuid ON notification_user_tokens(notification_uuid);
CREATE INDEX idx_notification_user_tokens_user_token_uuid ON notification_user_tokens(user_token_uuid);

CREATE TABLE notification_topics (
    notification_uuid UUID NOT NULL,
    topic_uuid UUID NOT NULL,
    PRIMARY KEY (notification_uuid, topic_uuid),
    FOREIGN KEY (notification_uuid) REFERENCES notification(uuid) ON DELETE CASCADE,
    FOREIGN KEY (topic_uuid) REFERENCES topic(uuid) ON DELETE CASCADE
);

CREATE INDEX idx_notification_topics_notification_uuid ON notification_topics(notification_uuid);
CREATE INDEX idx_notification_topics_topic_uuid ON notification_topics(topic_uuid);

-- This is a more granular table, it can control per device
CREATE TABLE user_token_topics (
    user_token_uuid UUID NOT NULL,
    topic_uuid UUID NOT NULL,
    PRIMARY KEY (user_token_uuid, topic_uuid),
    FOREIGN KEY (user_token_uuid) REFERENCES user_token(uuid) ON DELETE CASCADE,
    FOREIGN KEY (topic_uuid) REFERENCES topic(uuid) ON DELETE CASCADE
);

CREATE INDEX idx_user_token_topics_user_token_uuid ON user_token_topics(user_token_uuid);
CREATE INDEX idx_user_token_topics_topic_uuid ON user_token_topics(topic_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_token_topics;
DROP TABLE IF EXISTS notification_topics;
DROP TABLE IF EXISTS notification_user_tokens;
-- +goose StatementEnd
