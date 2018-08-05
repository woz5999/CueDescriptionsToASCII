FROM golang:1.9 as build
WORKDIR $GOPATH/src/github.com/woz5999/CueDescriptionsToASCII
COPY ./ ./
ARG VERSION
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /cuedescriptionstoascii

FROM alpine:3.6
WORKDIR /app
COPY --from=build /cuedescriptionstoascii /bin
ENTRYPOINT ["cuedescriptionstoascii"]
EXPOSE 80
