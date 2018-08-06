FROM gcr.io/cloud-builders/go

ENV GOPATH=/go
ENV GOBIN=/go/bin

ADD main.go /go/src
RUN go install /go/src/main.go

ENTRYPOINT ["/go/bin/main"]