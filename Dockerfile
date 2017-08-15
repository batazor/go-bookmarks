FROM golang:1.8-alpine as builder

ENV GLIDE_VERSION v0.12.3

# Install glide
RUN apk add --update ca-certificates wget git && \
    update-ca-certificates && \
    wget https://github.com/Masterminds/glide/releases/download/${GLIDE_VERSION}/glide-${GLIDE_VERSION}-linux-amd64.tar.gz && \
    tar -zxf glide-${GLIDE_VERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/glide /usr/local/bin/

# Build project
WORKDIR /go/src/app
COPY . .
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/config.yaml .
CMD ["./app"]