FROM golang:1.22-alpine3.18

# Install Deps...
RUN apk add nodejs npm

WORKDIR /app

COPY . .

WORKDIR /app/frontend

RUN npm i

RUN npm run build

WORKDIR /app

RUN mkdir bin

RUN go build -o ./bin/ main.go

RUN ls

ENTRYPOINT ./bin/main
