FROM golang:1.17 AS build
COPY . /src
RUN cd /src && BUILD_PREFIX=/ make server

FROM scratch
COPY --from=build /server /
COPY quotes.txt /
ENTRYPOINT ["/server", "--quotes=/quotes.txt", "--listen=:9000"]
EXPOSE 9000
