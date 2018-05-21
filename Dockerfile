FROM flemay/golang:1-stretch AS builder
COPY . /go/src/github.com/flemay/envvars/
WORKDIR /go/src/github.com/flemay/envvars
ENV IS_SCRATCH_IMAGE true
RUN make _deps _test _build _run

FROM scratch
LABEL maintainer "@flemay"
COPY --from=builder /go/src/github.com/flemay/envvars/bin/envvars /
ENTRYPOINT [ "/envvars" ]
