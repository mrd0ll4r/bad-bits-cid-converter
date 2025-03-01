FROM golang:1.24-bullseye AS builder

WORKDIR /usr/src/bad-bits-converter/
# Download all dependencies first, this should be cached.
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build ./cmd/bad-bits-cid-converter/

FROM debian:bullseye-slim AS runner

# Create a system user to drop into.
RUN groupadd -r ipfs \
  && useradd --no-log-init -r -g ipfs ipfs

# Enter our working directory.
WORKDIR bad-bits-cid-converter

# Copy compiled binaries from builder.
COPY --from=builder /usr/src/bad-bits-converter/bad-bits-cid-converter ./bad-bits-cid-converter

# Set ownership.
RUN chown -R ipfs:ipfs ./bad-bits-cid-converter

# Drop root.
USER ipfs

# Run the binary.
ENTRYPOINT ["./bad-bits-cid-converter"]

