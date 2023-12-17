FROM golang:latest

LABEL chipi="chipi"
LABEL chapa="chapa"
LABEL dubi="dubi"
LABEL daba="daba"
LABEL magico="mi"
LABEL dubi="dubi"
LABEL boom="boom"
LABEL boom="boom"

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o app

EXPOSE 8000

CMD ["/app/app"]
