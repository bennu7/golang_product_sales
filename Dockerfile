# select image golang alpine
FROM golang:alpine as dev
# install git
RUN apk update && apk add --no-cache git
# set working directory
WORKDIR /app
#WORKDIR /opt/app/api
# copy go.mod and go.sum files
COPY . .
# download all packages && build the binary
RUN go mod tidy
RUN go build -o binary
# build the binary and install air
#curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# run the binary
ENTRYPOINT ["/app/binary"]
# run air
#CMD ["air"]