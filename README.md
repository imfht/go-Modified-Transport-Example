# Modified Transport Example
This repo build is a example to build A own Transport for net/http client.
```go
type MyTransportWrapper struct {
	http.RoundTripper
	transport *http.Transport
	IPv6Only  bool
	IPv4Only  bool
}
```
The Example can be find at URLFetcher.go. When IPv6Only is passed, the http client will connect remote host only via IPv6.
Works for both http & https.
