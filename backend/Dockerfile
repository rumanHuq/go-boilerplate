# Compile stage
FROM golang:1.15.2 AS build-env
RUN go get github.com/pilu/fresh
EXPOSE 3000
WORKDIR /go/src/webapp/backend
CMD [ "fresh" ]

# # Final stage
# FROM debian:buster
# EXPOSE 3000
# WORKDIR /
# COPY --from=build-env /server /
# CMD ["/server"]
