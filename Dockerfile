FROM flemay/golang:1.9.3-stretch AS builder
RUN apt-get update
WORKDIR /go/src/github.com/flemay
RUN git clone https://github.com/flemay/envvars.git
WORKDIR /go/src/github.com/flemay/envvars
RUN make _deps _test _buildForScratch

FROM scratch
COPY --from=builder /go/src/github.com/flemay/envvars/bin/envvars /
ENTRYPOINT [ "/envvars" ]