FROM golang:1.14
WORKDIR /app
COPY . .

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"
RUN go build -o go-figures .

ENTRYPOINT ["./go-figures"]