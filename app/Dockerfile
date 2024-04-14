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

# If this was a real deployment we'd create a second image including just the binary and static files

ENTRYPOINT ./bin/main
