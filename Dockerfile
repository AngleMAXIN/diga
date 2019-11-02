FROM golang:latest


ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . /app

RUN go build .

EXPOSE 3001
ENTRYPOINT ["./Analysis-statistics"]