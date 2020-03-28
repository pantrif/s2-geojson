FROM golang:1.13 as builder
RUN cd ..
RUN mkdir s2-geojson
WORKDIR s2-geojson
COPY . ./
RUN useradd -r -u 999 gopher
RUN CGO_ENABLED=0 go build -mod=vendor -a -installsuffix cgo -o s2-geojson ./cmd/s2-geojson

FROM scratch
COPY --from=builder /go/s2-geojson/website/ ./website
COPY --from=builder /go/s2-geojson/s2-geojson .
COPY --from=builder /etc/passwd /etc/passwd
USER gopher
CMD ["./s2-geojson"]
