FROM golang:1.17-alpine as builder
ARG USER=default
ENV HOME /home/$USER
RUN apk add --update sudo
RUN adduser -D $USER \
        && echo \
     "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER
USER $USER
WORKDIR $HOME
COPY . .
RUN go mod download
RUN go build -o $HOME/bird-app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/default/bird-app .
COPY db db
COPY assets assets
EXPOSE 8080
CMD ["/root/bird-app"]