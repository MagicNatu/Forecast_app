# use an existing docker image as a base
FROM golang:1.11.1 as builder
#Installing additional tools/dependencies for the container
WORKDIR '/go/src/app'
COPY . .
RUN go get -d -v github.com/tools/godep
RUN go test -c -o /out/tests 
RUN go build -o /out/runme
CMD ["/out/runme"]


#CMD COMMANDS------------------------------------------

#FROM golang:alpine
#WORKDIR '/go/src/app'
#COPY --from=builtapp /go/src/app /go/src/app
#WORKDIR '/go/src/app'
#CMD ["app"]

# commands: docker run -p 8080:8080 -v "//c/Users/Public/Weather_app_development:/go/src/app" ae75d2ef6e30
