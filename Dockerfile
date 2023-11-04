FROM --platform=linux/amd64 golang:1.21.3-alpine as builder

WORKDIR /home/app

COPY . /home/app

RUN go mod download
RUN go build -o /app ./cmd/app


FROM --platform=linux/amd64 alpine:3.18.0

RUN apk --no-cache add gcompat tini
COPY --from=builder /app /app

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app" ]