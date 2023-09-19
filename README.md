# bad-bits-cid-converter

Converts CIDs and IPFS URIs to the format expected by the [bad bits denylist](https://badbits.dwebops.pub/).

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

The tool reads lines from STDIN and writes converted results to STDOUT, together with the original.
Errors are written to STDERR and can be safely redirected.

Input:
```
<CID 1>
<CID 2>/path
...
```

Output:
```
<CID 1> <tab> <bad bits formatted version>
<CID 2> <tab> <bad bits formatted version>
...
```

You can easily pipe data into it, like so:
```
cat <some list of CIDs> | bad-bits-cid-converter 2>/dev/null | ...
```

Input lines can be either
- blank CIDs, e.g. `Qm.....`
- IPFS URIs, e.g. `ipfs://Qm.../<optional path>`
- CIDs with subdirectory paths, e.g. `Qm.../<some path>/<more path>`
- Absolute IPFS paths, e.g. `/ipfs/Qm.../<some path>/<more path>`

They will be converted to `<base32 CIDv1>/<optional path>`, SHA256 hashed, and hex-encoded.

The results can then be searched for in the [deny list](https://badbits.dwebops.pub/)

### Comparing to the deny list

You can download and prepare the deny list like so:
```
curl https://badbits.dwebops.pub/denylist.json | jq '.[].anchor' -r > deny-list.txt
```

And then compare output from this tool like so:
```
cat <some list of CIDs> | bad-bits-cid-converter 2>/dev/null | grep -F -f deny-list.txt
```

## License

MIT, see [LICENSE](LICENSE).