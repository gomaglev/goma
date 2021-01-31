FROM golang:latest 
COPY . /go/src/GOMA
WORKDIR /go/src/GOMA
RUN go get ./cmd/microshop
RUN go build -ldflags "-w -s" -o ./cmd/microshop/microshop ./cmd/microshop
RUN touch .env
CMD ["microshop", "start"]