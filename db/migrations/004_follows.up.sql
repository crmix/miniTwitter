CREATE TABLE follows (
    id uuid PRIMARY KEY,
    follower_id uuid REFERENCES users(id) on DELETE CASCADE,
    followed_id uuid REFERENCES users(id) on DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (follower_id, followed_id)
);
