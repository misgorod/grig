FROM golang:1.13
COPY . /src
WORKDIR /src
RUN go build -o /main
CMD /main