FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN rm Dockerfile

RUN go build -o /cosmos-notifier

#EXPOSE 50051

CMD [ "/cosmos-notifier" ]
