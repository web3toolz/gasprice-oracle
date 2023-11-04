FROM golang:1.21.3-alpine as builder

WORKDIR /home/app

COPY . /home/app

RUN go mod download
RUN go build -o /app ./cmd/app


FROM scratch

COPY --from=builder /app /app

CMD [ "/app" ]