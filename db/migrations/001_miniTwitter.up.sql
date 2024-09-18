CREATE TABLE users (
    id uuid PRIMARY KEY,                      
    username VARCHAR(100) UNIQUE NOT NULL,      
    password TEXT NOT NULL,               
    name VARCHAR(100),                          
    bio TEXT,                                   
    profile_image TEXT,                        
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP        
);

