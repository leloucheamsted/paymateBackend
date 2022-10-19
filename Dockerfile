# ##
# ## Build
# ##

# FROM golang:1.16-buster AS build

# WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# COPY *.go ./

# RUN go build -o /paymate

# ##
# ## Deploy
# ##

# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /paymate /paymate

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/paymate"]

FROM golang:1.16-alpine
WORKDIR /go/src/fm-api
COPY . .

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o  paymate .


EXPOSE 8080

CMD [ "/paymate" ]