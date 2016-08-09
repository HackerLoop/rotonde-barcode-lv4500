FROM resin/raspberrypi-golang:1.5

RUN apt-get update && \
    apt-get install -y libusb-1.0

RUN go get github.com/HackerLoop/rotonde && \
    go get github.com/HackerLoop/rotonde-client-go

RUN go get github.com/HackerLoop/rotonde-barcode-lv4500 && \
    go install github.com/HackerLoop/rotonde-barcode-lv4500 && \
    mv /go/bin/rotonde-barcode-lv4500 /usr/local/bin/

RUN rm -rf $GOPATH $GOROOT

CMD ["/usr/local/bin/rotonde-barcode-lv4500"]
