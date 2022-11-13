FROM golang:1.19-alpine

WORKDIR /app

ADD . .

RUN go mod tidy

RUN go build -o /new_smthp_server

EXPOSE 5555

CMD [ "/new_smthp_server" ]