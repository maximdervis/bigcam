FROM golang:1.24
WORKDIR /source
COPY . .
RUN go build -o gen_server ./src
CMD ["./gen_server"]

