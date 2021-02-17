FROM golang:alpine

WORKDIR /go/src

COPY . .

RUN GOOS=linux go build main.go

CMD [ "./main" ]