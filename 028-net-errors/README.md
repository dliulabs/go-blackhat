# Net Errors

[net.Error](https://pkg.go.dev/net#Error)
[net.OpError](https://golang.org/pkg/net/#OpError/)

```
if nErr, ok := err.(net.Error); ok && !nErr.Temporary() { return err }
```
