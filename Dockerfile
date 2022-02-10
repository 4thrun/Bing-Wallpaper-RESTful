# build 
FROM golang:alpine AS build-env 
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
WORKDIR /
COPY . .
RUN go build -o bw .

# run 
FROM alpine 
WORKDIR / 
RUN mkdir /image
COPY --from=build-env /bw /
ADD ./image /image/
# server port
EXPOSE 9002 
CMD ["/bw", "run"]
 