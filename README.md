# 🔒 Coauth
**_A Go Server Wrapped in a Docker Container To Implement Authentication / Authorization To Your Project Fast_**
[Coauth Docker Image](https://hub.docker.com/r/ralflopez/coauth)

# Features
1. Session Authentication (Cookie)
2. JWT Authentication (Access and Refresh Tokens)
3. Role Based Authenticaion (Admin, Member and Guest)

# ➕  Installation
## Docker
1. Start coauth + postgres container
docker-compose.yml example:
```yml
version: '3.1'
services:
  api:
    image: ralflopez/coauth
    ports:
      - "9000:9000"
    depends_on:
      - db
      - adminer
      - cache
    environment:
      - PORT=9000
      - DATABASE_HOST=db
      - DATABASE_PORT=5432
      - DATABASE_USER=user
      - DATABASE_PASSWORD=password
      - DATABASE_NAME=coauth
      - URL=http://localhost:3000
    networks:
      - coauthapi
  db:
    image: postgres
    restart: always
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: password
      POSTGRES_USER: user
    networks:
      - coauthapi
  adminer:
    image: adminer
    restart: always
    ports:
      - "8080:8080"
    networks:
      - coauthapi
networks:
  coauthapi:
    driver: bridge
```

# 📙 Guide
To integrate *coauth* with your app. You can send an http request to the container from the available routes below.

## 🧪 Status
To test if the server is running the test route should return a 200 status
```
Route: /test
```

## 🔐 Session
### Login (session)
Stores a user in the session store.
Body: email, password
Return: user object and sets a session id as cookie
```
Route: /session/login
Body:
{
    "email": "EMAIL_HERE",
    "password": "PASSWORD_HERE"
}
```
### Signup (session)
Creates and stores user in the session store.
Body: name, email, password
Return: user object and sets a session id as cookie
```
Route: /session/signup
Body:
{
    "name": "NAME_HERE",
    "email": "EMAIL_HERE",
    "password": "PASSWORD_HERE"
}
```
### Get Logged In User (session)
Gets a user object based on the session id cookie
Return: a user object of the current logged in user
```
Route: /session/user
```

## 🔐 JSON Web Token
### Login (JWT)
Generates new access and refresh JWT tokens
Body: email, password
Return: object containing access and refresh token, refresh token as cookie
```
Route: /jwt/login
Body:
{
    "email": "EMAIL_HERE",
    "password": "PASSWORD_PASSWORD"
}
```
### Signup (JWT)
Creates a user and generates access and refresh JWT tokens
Body: email, password
Return: user object and sets a session id as cookie
```
Route: /jwt/login
Body:
{
    "email": "EMAIL_HERE",
    "password": "PASSWORD_HERE"
}
```
### Get Logged In User (JWT)
Gets a user object based on the Bearer token
Return: a user object of the current logged in user, sets a new refresh token as cookie
```
Route: /session/user
Header.Authorization: "Bearer JWT_TOKEN_HERE"
```