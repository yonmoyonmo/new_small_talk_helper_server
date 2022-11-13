FROM golang:1.19-alpine

WORKDIR /app

ADD . .

RUN go mod download

RUN go build -o /new_smthp_server

EXPOSE 5555

CMD [ "/new_smthp_server" ]