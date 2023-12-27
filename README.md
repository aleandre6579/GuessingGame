# GuessingGame

A simple containerized web app made with React, Go and Docker.

<br/>

## Frontend
Made in React. Used [tsparticles](https://github.com/tsparticles/tsparticles) for the snow particles effect.

<br/>

## Backend
Made in Go. Used [chi](https://github.com/go-chi/chi) as a router, [gorm](https://github.com/go-gorm/gorm) for the database connection and [golang-jwt](https://github.com/golang-jwt/jwt) for jwt authentication.
Has 3 endpoints: '/login', '/register' and '/guess'.

<br/>

## Installation Guidelines
First clone the repository:
```
git clone git@github.com:aleandre6579/GuessingGame.git
```
> **_NOTE:_**  I kept the .env file in the repo for convenience sake but it's bad practice.

<br/>

Then go into the project's root folder and run the docker-compose:
```
cd GuessingGame && docker compose pull && docker compose up -d
```

<br/>

Finally visit <a href="http://127.0.0.1" target="_blank">localhost</a>
