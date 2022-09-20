FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN rm Dockerfile

RUN go build -o /dao-dao-notifier

#EXPOSE 50051

CMD [ "/dao-dao-notifier" ]