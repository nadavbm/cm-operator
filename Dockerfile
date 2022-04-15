# Build the manager binary
FROM golang:1.17-stretch as builder

# Copy the go source
COPY . /cm-operator
WORKDIR /cm-operator

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o cm-operator main.go


FROM alpine:latest
WORKDIR /
COPY --from=builder /cm-operator/cm-operator /cm-operator

CMD /cm-operator
