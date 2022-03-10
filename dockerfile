FROM golang:1.17-alpine
ENV GOPROXY=https://goproxy.cn
RUN mkdir /app 
ADD . /app
EXPOSE 8080

WORKDIR /app
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17

RUN go build -o /app main.go
CMD ["/app/main"]
