FROM golang:1.17 AS build
COPY . /src
RUN cd /src && BUILD_PREFIX=/ make client

FROM scratch
COPY --from=build /client /
ENTRYPOINT ["/client"]
