# tart by building the application.
# Build em crud-dao com distroless
FROM golang:1.15 as builder
WORKDIR /go/src/crud.user.singleton
COPY . .
ENV GO111MODULE=on
#CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --trimpath -a -tags netgo -ldflags '-w -extldflags "-static"' -o crud-dao main.go
CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags '-s -w' -o crud.user.singleton main.go
RUN cp crud.user.singleton /go/bin/crud.user.singleton
RUN ls -lh
# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/crud.user.singleton /
CMD ["/crud.user.singleton"]