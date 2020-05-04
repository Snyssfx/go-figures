FROM golang:1.14 as builder
WORKDIR /app
COPY . .

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"
RUN go build -o go-figures .

FROM alpine:latest
COPY --from=builder /app .

ENTRYPOINT ["./go-figures"]