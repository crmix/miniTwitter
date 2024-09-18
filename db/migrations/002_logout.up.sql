CREATE TABLE blocked_tokens (
    id uuid PRIMARY KEY,           
    token VARCHAR(255) UNIQUE NOT NULL, 
    expires_at BIGINT NOT NULL       
);