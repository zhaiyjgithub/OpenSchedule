FROM golang:1.15-alpine

WORKDIR /app

#COPY go.mod .

#COPY go.sum .

#RUN go env -w GOPROXY=https://goproxy.io
#
#RUN go env
#
#RUN go mod download

COPY anyHealthApp .

RUN  chmod +x anyHealthApp

EXPOSE 8090

CMD ["/app/anyHealthApp"]

