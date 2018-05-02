This module contains the API written in golang to handle 
the user registration, login and logout functionality of the 
Starbucks application (which our team is developing as part of Team Hackathon Project). 

MongoDB is used as the database for managing users.


| API Route           | Method           | Description                                                    |
| --------------------|------------------| ---------------------------------------------------------------|
| /signup      | POST                    | Responsible for getting a Registering a new user on the website          |
| /login       | POST                    | Responsible for signing a user and creating a session for that user      |
| /logout      | POST                   | Responsible for clearing the session and signing out the logged in user  |
