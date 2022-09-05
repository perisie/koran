FROM golang:1.18.3
WORKDIR /app
COPY . .
EXPOSE 8080
RUN ./gobuild.sh
CMD ./gorun.sh
