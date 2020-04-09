FROM golang:latest

# Make work directory
RUN mkdir /app
WORKDIR /app

# Copy app
COPY ./server .

# Build the Go app
RUN go build -o main .

# Command to run the executable
EXPOSE 8080
CMD ["./main"]