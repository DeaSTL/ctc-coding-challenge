### Requirements

The first step in building a new consumer-facing web application is the successful
storage of authentication credentials. We would like you to create an API that
allows users to sign up for an imaginary website and log in securely.

- [x] Create a Postgres or Mongo database to hold log-in credentials. Make
sure to include the schema and necessary table creation scripts for your
database in your github repo.

-  [x] Create the necessary REST endpoints allowing the user to sign up with an
email address and password, log-in with those credentials, and logout. Use
whatever language or framework you are comfortable with. At CTC you will
primarily be working with GoLang, Express, and PHP- so these are
certainly the preferred languages for this assignment.

- [x] Build a web page using HTML, CSS, Javascript, React, or PHP that allows
users to create an account, login, and logout.

- [x] After a successful login, display a successful login message and show the
logout option.

- [x] Employ error handling and data validation on the front and back end as
you see appropriate.

- [x] Include a Postman collection and/or sufficient documentation for testing the
routes and validating functionality.

- [x] Containerize your application using Docker and include the
Dockerfile in your repo as well and include the build run commands necessary to
deploy it in the readme file of the repo.

### Running the application

```bash
# After cloning project
cd ./ctc-coding-challenge
docker compose up
# pre compose v2
docker-compose up
```
In your browser enter `http://localhost:3000`

![2024-04-14_02-13](https://github.com/DeaSTL/ctc-coding-challenge/assets/19532324/15d08e13-d735-4f5e-9ffa-df8280528cdd)


### Stack

- Postgres
- Golang + Fiber
- Vanilla React

