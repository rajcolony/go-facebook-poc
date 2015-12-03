# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
#FROM golang

# Copy the local package files to the container's workspace.
#ADD . /src/poc

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
#RUN go install poc

# Run the outyet command by default when the container starts.
#ENTRYPOINT /bin/GoPOC

# Document that the service listens on port 8888.
#EXPOSE 8888

FROM golang:onbuild
EXPOSE 8888