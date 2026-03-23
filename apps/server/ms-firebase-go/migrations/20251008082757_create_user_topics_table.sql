-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_topics (
    user_uuid TEXT NOT NULL,
    topic_uuid UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_at TIMESTAMP,
    updated_by TEXT,
    PRIMARY KEY (user_uuid, topic_uuid),
    FOREIGN KEY (topic_uuid) REFERENCES topic(uuid) ON DELETE CASCADE
);

CREATE INDEX idx_user_topics_user_uuid ON user_topics(user_uuid);
CREATE INDEX idx_user_topics_topic_uuid ON user_topics(topic_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_topics;
-- +goose StatementEnd
