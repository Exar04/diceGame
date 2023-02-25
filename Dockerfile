FROM golang:alpine
RUN mkdir /DICEGAME
COPY . /DICEGAME
WORKDIR /DICEGAME
RUN go build -o main .
CMD ["/DICEGAME/main"]