FROM resin/raspberrypi-golang:1.5

RUN apt-get update && \
    apt-get install -y libusb-1.0

RUN go get github.com/HackerLoop/rotonde && \
    go get github.com/HackerLoop/rotonde-client-go

ADD . $GOPATH/src/github.com/HackerLoop/rotonde-barcode-lv4500

RUN go install github.com/HackerLoop/rotonde-barcode-lv4500

RUN cp $GOPATH/bin/rotonde-barcode-lv4500 /usr/local/bin/

CMD ["/usr/local/bin/rotonde-barcode-lv4500"]
