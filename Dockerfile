FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
RUN mkdir -p /home/www/bw
WORKDIR /home/www/bw

COPY . .

RUN go build -o bw .

# server port
EXPOSE 9002

# run
CMD ["/home/www/bw/bw", "run"]
 