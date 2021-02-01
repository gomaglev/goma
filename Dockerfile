FROM golang:latest 
COPY . /go/src/app_name__
WORKDIR /go/src/app_name__
RUN go get ./cmd/app_name__
RUN go build -ldflags "-w -s" -o ./cmd/app_name__/app_name__ ./cmd/app_name__
RUN touch .env
CMD ["app_name__", "start"]