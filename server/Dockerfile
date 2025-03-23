FROM golang:1.24
WORKDIR /source
COPY . .
RUN go build -o server ./src
CMD ["./server"]

