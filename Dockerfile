FROM golang:1.17-alpine

ARG USER=default
ENV HOME /home/$USER

RUN apk add --update sudo
RUN adduser -D $USER \
        && echo \
     "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY assets assets
COPY *.go ./

RUN go build -o $HOME/bird-app

EXPOSE 8080
CMD ["/home/default/bird-app"]