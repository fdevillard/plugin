# Traefik plugin

This is a simple Traefik plugin for a middleware that wraps the `http.ResponseWriter` in order to investigate how
it deals with websocket connections (or any other that would require a `http.Hijacker`).

## Current status
The test is supposed to validate that such connection can be handled by the middleware. The test passes.

However, when run in real conditions, the middleware panics with the following messages:

```
{"level":"error","module":"github.com/fdevillard/plugin","msg":"plugins-local/src/github.com/fdevillard/plugin/main.go:21:8: panic","plugin":"plugin-audit","time":"2022-01-06T15:42:54Z"}
{"level":"error","middlewareName":"traefik-internal-recovery","middlewareType":"Recovery","msg":"Recovered from panic in HTTP handler [172.20.0.27:36380 - /graphql]: interface conversion: stdlib._net_http_ResponseWriter is not http.Hijacker: missing method Hijack","time":"2022-01-06T15:42:54Z"}
```
