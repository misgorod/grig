FROM golang:1.13
COPY . /src
WORKDIR /src
RUN go build -o /main . && \
    chmod ugo+x /main
CMD /main