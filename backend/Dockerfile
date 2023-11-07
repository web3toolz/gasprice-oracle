FROM --platform=linux/amd64 golang:1.21.3-alpine as builder

WORKDIR /home/app

COPY . /home/app

RUN go mod download
RUN go build -o /gasprice-oracle ./cmd/cli


FROM --platform=linux/amd64 alpine:3.18.0

RUN mkdir /home/app
WORKDIR /home/app

RUN apk --no-cache add gcompat tini
COPY --from=builder /gasprice-oracle /home/app/gasprice-oracle
# copy default settings
COPY config.example.yaml /home/app/config.example.yaml

# expose default port
EXPOSE 8000

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/home/app/gasprice-oracle", "server" ]