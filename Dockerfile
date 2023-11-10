FROM golang:1.20
EXPOSE 9099

WORKDIR /prj

ENV GOPROXY=https://goproxy.io

COPY . /prj/
RUN apt-get install gcc
RUN go build main.go
CMD ./main