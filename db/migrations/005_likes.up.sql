CREATE TABLE likes (
    id uuid PRIMARY KEY,
    tweet_id uuid NOT NULL REFERENCES tweets(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (tweet_id, user_id) 
);