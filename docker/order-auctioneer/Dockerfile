###############################################################################
# STEP 1 build executable binary
FROM golang:1.17-alpine as builder
# Install SSL ca certificates
RUN apk add --no-cache git ca-certificates
# Build Go app
COPY ./go/pkg/zerologconfig $GOPATH/src/pkg/zerologconfig
COPY ./go/pkg/zerologr $GOPATH/src/pkg/zerologr
COPY ./go/pkg/apis $GOPATH/src/pkg/apis
COPY ./go/cmd/order-auctioneer $GOPATH/src/cmd/order-auctioneer
WORKDIR $GOPATH/src/cmd/order-auctioneer
# Run tests
RUN CGO_ENABLED=0 GOOS=linux go test -v
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/order-auctioneer
###############################################################################
# STEP 2 app image
FROM alpine:latest
# Install SSL ca certificates
RUN apk add --no-cache ca-certificates
# Copy Go static executable
COPY --from=builder /go/bin/order-auctioneer /cmd/order-auctioneer
RUN chmod +x /cmd/order-auctioneer

ENTRYPOINT ["/cmd/order-auctioneer"]

# Create & use appuser
RUN adduser -D appuser
USER appuser