# Project API using Golang

This repository contains a Golang-based API project that provides endpoints for managing users and their associated photos. The API follows a RESTful architecture and includes the following features:

## _User Management_

### User Attributes

-   ID (primary key, required)
-   Username (required)
-   Email (unique & required)
-   Password (required & minlength 6)
-   Relationship with the Photo model (using constraint cascade)
-   Created At (timestamp)
-   Updated At (timestamp)

### User Endpoints

1. **POST - /users/register:**

    - _Description:_ Register a new user.
    - _Request Body:_
        ```json
        {
            "username": "example_username", // required
            "email": "example@example.com", // required
            "password": "password123" // required minimum length 6
        }
        ```
    - _Response:_ 201 Created
    - _Response Body:_
        ```json
        {
            "message": "registration success"
        }
        ```

2. **POST - /users/login:**
    - _Description:_ Login an existing user and obtain an authentication token.
    - _Request Body:_
        ```json
        {
            "email": "example@example.com", // required
            "password": "password123" // required
        }
        ```
    - _Response:_ 200 OK
    - _Response Body:_
        ```json
        {
            "message": "login success"
        }
        ```
3. **GET - /users/logout:**

    - _Description:_ Logout the currently logged-in user.
    - _Response:_ 200 OK
    - _Response Body:_
        ```json
        {
            "message": "logout success"
        }
        ```

4. **PUT - /users/:userId:**

    - _Description:_ Update the currently logged-in user. User must provide the current password for verification.
    - _Request Body:_
        ```json
        {
            "new_username": "user123", // optional
            "new_email": "user123@example.com", // optional
            "new_password": "newpassword456", // optional
            "password": "admin123" // required
        }
        ```
    - _Response:_ 200 OK
    - _Response Body:_
        ```json
        {
            "message": "update success"
        }
        ```

5. **DELETE - /users/:userId:**
    - _Description:_ Delete the currently logged-in user.
    - _Response:_ 204 No Content
    - _Response Body:_
        ```json
        {
            "message": "delete user success"
        }
        ```

## _Photo Management_

### Photo Attributes

-   ID (primary key, required)
-   Title (required)
-   Caption (optional)
-   PhotoUrl (required)
-   UserID (foreign key, required)
-   Relationship with the User model
-   Created At (timestamp)
-   Updated At (timestamp)

### Photo Endpoints (User must be logged in)

1. **POST - /photos:**

    - _Description:_ Upload a new photo.
    - _Request Body:_
        ```json
        {
            "title": "Example Photo", // required
            "caption": "A description of the photo",
            "photoUrl": "https://example.com/photo.jpg" // required
        }
        ```
    - _Response:_ 201 Created
    - _Response Body:_
        ```json
        {
            "message": "photo uploaded"
        }
        ```

2. **GET - /photos:**

    - _Description:_ Get all photos uploaded by the currently logged-in user.
    - _Response:_ 200 OK
    - _Response Body:_
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

    - _Description:_ Update a photo owned by the currently logged-in user. User must provide the current password for verification
    - _Request Body:_
        ```json
        {
            "new_title": "Updated Photo Title",
            "new_caption": "Updated photo description",
            "password": "admin123" // required
        }
        ```
    - _Response:_ 200 OK
    - _Response Body:_
        ```json
        {
            "message": "your photo successfully changed"
        }
        ```

4. **DELETE - /photos/photoId:**
    - _Description:_ Delete a photo owned by the currently logged-in user.
    - _Response:_ 204 No Content
    - _Response Body:_
        ```json
        {
            "message": "delete photo success"
        }
        ```

## Authentication

-   User authentication is required for all endpoints in the Photo Management section. Upon successful login, a token is provided, which can be used for managing user data.

## Tools

[![My Skills](https://skillicons.dev/icons?i=go,postman,git,github,mysql,vscode)](https://skillicons.dev)