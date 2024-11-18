# go-vyos

go-vyos is a Go client library for accessing the [Vyos API](https://docs.vyos.io/en/latest/automation/vyos-api.html#vyosapi)

# Usage

```go
import "github.com/ganawaj/go-vyos/vyos"
```

Construct a new Vyos client, then use the various services on the client to access different parts of the Vyos API. For example:

```go
c := vyos.NewClient(nil).WithToken("AUTH_KEY").WithURL("https://192.168.0.1")
```

If you're using self-signed certificates or don't want to verify certificates, you can disable TLS verification by adding .Insecure():

```go
c := vyos.NewClient(nil).WithToken("AUTH_KEY").WithURL("https://192.168.0.1").Insecure()
```