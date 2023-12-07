FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .
RUN go mod download -x

RUN go get ./... 
RUN go build -o /bin/worker

EXPOSE 8888

CMD [ "/bin/worker" ]