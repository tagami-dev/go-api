# Build stage
FROM golang:1.22.3-alpine3.18 AS build

# Install git, required for fetching Go dependencies.
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /go/src/project/

# Copy go module files and download dependencies
COPY go.mod go.sum /go/src/project/
RUN go mod download

# Copy the rest of the source code
COPY . /go/src/project/

# Build the Go app
RUN go build -o /bin/project

# Final stage
FROM scratch

# Copy the built binary from the build stage to the final image
COPY --from=build /bin/project /bin/project

# Set the binary as the entrypoint of the container
ENTRYPOINT ["/bin/project"]