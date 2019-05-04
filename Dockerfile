FROM scratch
LABEL maintainer "@flemay"
COPY ./bin/envvars /
ENTRYPOINT [ "/envvars" ]
