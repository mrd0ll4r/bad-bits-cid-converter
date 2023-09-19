# bad-bits-cid-converter

Converts CIDs into the format expected by the [bad bits denylist](https://badbits.dwebops.pub/).

## Building

On recent versions of go, build via
```
go build ./cmd/bad-bits-cid-converter/
```

Alternatively, build within docker via
```
./build-in-docker.sh
```
The output will be placed in `out/`

## Running

The too reads lines from STDIN and writes converted results to STDOUT.
Errors are written to STDERR and can be safely redirected.

Input lines should be either
- blank CIDs, e.g. `Qm.....`
- IPFS URIs, e.g. `ipfs://Qm.../<optional path>`
- CIDs with subdirectory paths, e.g. `Qm.../<some path>/<more path>`
- Absolute IPFS paths, e.g. `/ipfs/Qm.../<some path>/<more path>`

They will be converted to `<base32 CIDv1>/<optional path>`, SHA256 hashed, and hex-encoded.
The results can then be searched for in the [deny list](https://badbits.dwebops.pub/)
