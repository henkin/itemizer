FROM alpine
COPY ./artifacts .

# install bash, for fun
RUN apk add --no-cache bash

EXPOSE 8080
ENTRYPOINT ["./go-service"]