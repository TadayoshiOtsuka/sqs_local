FROM golang:1.18.3

ENV TZ=Asia/Tokyo
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV ROOTPATH=/go/app
ENV QUEUE_URL=http://queue:9324/queue/default
ENV AWS_REGION=ap-northeast-1
ENV AWS_ACCESS_KEY_ID=dummy
ENV AWS_SECRET_ACCESS_KEY=dummy

WORKDIR ${ROOTPATH}

RUN go install github.com/cosmtrek/air@v1.29.0
COPY go.mod go.sum .air.toml ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
