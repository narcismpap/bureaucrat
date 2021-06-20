# Bureaucr.at Coding Challenge
# Author: Narcis M. Pap - https://www.linkedin.com/in/narcismpap/
# London, Jun 2021
# github.com/narcismpap/bureaucrat

FROM golang:1.16.5-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p "$GOPATH/src/bureaucrat"
WORKDIR $GOPATH/src/bureaucrat
ADD /bureaucrat .
ADD /go.mod .

RUN go mod download

RUN go build \
    -installsuffix 'static' \
    -o /bureaucrat .

FROM scratch AS final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /bureaucrat /bureaucrat

EXPOSE 9000
USER nobody:nobody

COPY ./payload/GoT.json /GoT.json
ENTRYPOINT ["/bureaucrat"]
