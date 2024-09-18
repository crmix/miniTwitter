CREATE TYPE state AS ENUM ('original', 'retweet');
CREATE TABLE tweets (
    id uuid PRIMARY KEY,
    content TEXT,
    image_url TEXT,
    video_url TEXT,
    user_id uuid REFERENCES users(id),
    state state NOT NULL,
    original_tweet_id uuid REFERENCES tweets(id), -- NULL for original tweets, references the original tweet for retweets
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
