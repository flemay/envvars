FROM flemay/golang:1.10.0-stretch AS builder
WORKDIR /go/src/github.com/flemay
RUN git clone https://github.com/flemay/envvars.git
WORKDIR /go/src/github.com/flemay/envvars
RUN make _deps _test _buildForScratch

FROM scratch
LABEL maintainer="Frederic Lemay"
COPY --from=builder /go/src/github.com/flemay/envvars/bin/envvars /
ENTRYPOINT [ "/envvars" ]