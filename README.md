# Project API using Golang

This repository contains a Golang-based API project that provides endpoints for managing users and their associated photos. The API follows a RESTful architecture and includes the following features:

## *User Management*

### User Attributes
- ID (primary key, required)
- Username (required)
- Email (unique & required)
- Password (required & minlength 6)
- Relationship with the Photo model (using constraint cascade)
- Created At (timestamp)
- Updated At (timestamp)

### User Endpoints

1. **POST - /users/register:**
   - *Description:* Register a new user.
   - *Request Body:*
     ```json
     {
       "username": "example_username", // required
       "email": "example@example.com", // required
       "password": "password123" // required
     }
     ```
   - *Response:* 201 Created
   - *Response Body:*
     ```json
     {
       "message": "registration success", 
     }
     ```

2. **POST - /users/login:**
   - *Description:* Login an existing user and obtain an authentication token.
   - *Request Body:*
     ```json
     {
       "email": "example@example.com", // required
       "password": "password123"  // required
     }
     ```
   - *Response:* 200 OK
   - *Response Body:*
     ```json
     {
       "message": "login success", 
     }
     ```
   
3. **GET - /users/logout:**
   - *Description:* Logout the currently logged-in user.
   - *Response:* 200 OK
   - *Response Body:*
     ```json
     {
       "message": "logout success", 
     }
     ```

4. **PUT - /users/:userId:**
   - *Description:* Update the currently logged-in user. User must provide the current password for verification.
   - *Request Body:*
     ```json
     {
       "new_username": "user123",  // optional
       "new_email": "user123@example.com",  // optional
       "new_password": "newpassword456", // optional
       "password": "admin123" // required
     }
     ```
   - *Response:* 200 OK
   - *Response Body:*
     ```json
     {
       "message": "update success", 
     }
     ```

5. **DELETE - /users/:userId:**
   - *Description:* Delete the currently logged-in user.
   - *Response:* 204 No Content

## *Photo Management*

### Photo Attributes
- ID (primary key, required)
- Title (required)
- Caption (optional)
- PhotoUrl (required)
- UserID (foreign key, required)
- Relationship with the User model
- Created At (timestamp)
- Updated At (timestamp)

### Photo Endpoints (User must be logged in)

1. **POST - /photos:**
   - *Description:* Upload a new photo.
   - *Request Body:*
     ```json
     {
       "title": "Example Photo", // required
       "caption": "A description of the photo",
       "photoUrl": "https://example.com/photo.jpg" // required
     }
     ```
   - *Response:* 201 Created
   - *Response Body:*
     ```json
     {
       "message": "photo uploaded", 
     }
     ```

2. **GET - /photos:**
   - *Description:* Get all photos uploaded by the currently logged-in user.
   - *Response:* 200 OK
   - *Response Body:*
     ```json
     {
        "photos": [
            {
                "id": 15,
                "user_id": 5,
                "title": "image",
                "caption": "user not set photo's caption",
                "photo_url": "images3.jpg",
                "created_at": "2024-01-24T06:31:24.111+07:00",
                "updated_at": "2024-01-24T06:31:24.111+07:00"
            },
            {
                "id": 14,
                "user_id": 5,
                "title": "image",
                "caption": "user not set photo's caption",
                "photo_url": "images2.jpg",
                "created_at": "2024-01-24T06:31:18.425+07:00",
                "updated_at": "2024-01-24T06:31:18.425+07:00"
            }
        ]
    }
     ```

3. **PUT - /photos/photoId:**
   - *Description:* Update a photo owned by the currently logged-in user. User must provide the current password for verification 
   - *Request Body:*
     ```json
     {
       "new_title": "Updated Photo Title",
       "new_caption": "Updated photo description",
       "password": "admin123" // required
     }
     ```
   - *Response:* 200 OK
   - *Response Body:*
     ```json
     {
       "message": "your photo successfully changed", 
     }
     ```

4. **DELETE - /photos/photoId:**
   - *Description:* Delete a photo owned by the currently logged-in user.
   - *Response:* 204 No Content
   - *Response Body:*
     ```json
     {
       "message": "delete photo success", 
     }
     ```

## Authentication
- User authentication is required for all endpoints in the Photo Management section. Upon successful login, a token is provided, which can be used for managing user data.

## Tools
[![My Skills](https://skillicons.dev/icons?i=go,postman,git,github,mysql,vscode)](https://skillicons.dev)