FROM golang:1.8-alpine

ENV GLIDE_VERSION v0.12.3

# Install glide
RUN apk add --update ca-certificates wget git && \
    update-ca-certificates && \
    wget https://github.com/Masterminds/glide/releases/download/${GLIDE_VERSION}/glide-${GLIDE_VERSION}-linux-amd64.tar.gz && \
    tar -zxf glide-${GLIDE_VERSION}-linux-amd64.tar.gz && \
    mv linux-amd64/glide /usr/local/bin/

# Build project
WORKDIR /go/src/ritfest-go-step-by-step
COPY . .
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ritfest-go-step-by-step .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/ritfest-go-step-by-step/ritfest-go-step-by-step .
COPY --from=0 /go/src/ritfest-go-step-by-step/config.yaml .
CMD ["./ritfest-go-step-by-step"]