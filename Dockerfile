# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:latest

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image

RUN mkdir /go/src/github.com/claudioontheweb
RUN git clone https://github.com/claudioontheweb/url-shortener /go/src/github.com/claudioontheweb/url-shortener
RUN go get -u github.com/golang/dep/cmd/dep

# Add Maintainer Info
LABEL maintainer="Claudio Weckherlin <claudio.weckherlin@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/claudioontheweb/url-shortener

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN dep ensure -v

# Build the Go app
RUN go build cmd/*.go .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./app"]