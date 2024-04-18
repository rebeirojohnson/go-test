FROM golang:1.22.2

WORKDIR /.

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV TZ="Asia/Kolkata"


ENTRYPOINT ["go"] 

CMD ["run", "mongo.go"]
