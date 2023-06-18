FROM golang:1.18.3
WORKDIR /app
COPY . .
EXPOSE 8080
RUN ./x_build.sh
CMD ./x_run.sh
