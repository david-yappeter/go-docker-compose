FROM golang:1.14-alpine as builder
WORKDIR /src/
COPY . /src/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/myapp

FROM scratch
COPY --from=builder /bin/myapp /bin/myapp
ENTRYPOINT [ "/bin/myapp" ]
