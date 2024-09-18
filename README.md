# miniTwitter
miniTwitter App

This is a mini Twitter-like application written in Golang, designed to implement basic social media functionalities such as posting tweets, following users, and viewing timelines.

Features

User Registration: Users can register with unique usernames and passwords.

Authentication: Secure user login using JWT tokens.

Tweet Posting: Users can post short messages (tweets).

Timeline View: Users can view their own timeline, which consists of tweets from users they follow.

Follow/Unfollow Users: Users can follow or unfollow other users.

Like Tweets: Users can like tweets.


Tech Stack

Backend: Golang

Database: PostgreSQL

Authentication: JWT tokens

API Design: RESTful API

Frameworks: Standard Go libraries, Gin and Gorm frameworks (or any other if you used one)


Installation

1. Clone the repository:

git clone https://github.com/yourusername/minitwitter.git


2. Navigate to the project directory:

cd miniTwitter


3. Install dependencies:

go mod tidy


4. Set up PostgreSQL:

Ensure you have PostgreSQL installed and running.

Create a new database for the app.

Update the database connection details in config.json or .env file.



5. Run the application:

go run main.go



API Endpoints
http://localhost:8080/api/auth/registr  method POST
http://localhost:8080/api/auth/login    method POST
http://localhost:8080/api/auth/logout   method POST
http://localhost:8080/api/tweets        method POST
http://localhost:8080/api/follow        method POST
http://localhost:8080/api/followers     method GET
http://localhost:8080/api/followings    method GET
http://localhost:8080/api/unfollow      method POST
http://localhost:8080/api/like          method POST
http://localhost:8080/api/unlike        method POST
http://localhost:8080/api/retweet       method POST
http://localhost:8080/api/search?q=me   method GET

Environment Variables

Make sure to set up the following environment variables:

DB_HOST

DB_USER

DB_PASSWORD

DB_NAME

JWT_SECRET


Future Improvements

Add comment functionality.

Implement real-time updates using WebSockets.

Enhance security with password encryption (if not done already).


License

This project is licensed under the MIT License - see the LICENSE file for details.


---

You can modify the placeholders or sections according to your actual implementation and project structure.
