# go-vyos

go-vyos is a Go client library for accessing the [Vyos API](https://docs.vyos.io/en/latest/automation/vyos-api.html#vyosapi)

## Usage

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

### Configure, then Set

```go
    out, resp, err := c.Conf.Set(ctx, "interfaces ethernet eth0 address 192.168.1.1/24")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
```

### Show a Single Object Value

```go

    out, resp, err := c.Show.Do(ctx, "interfaces dummy dum1 address")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)
```

### Configure, then Show Object

```go

    out, resp, err := c.Conf.Get(ctx, "interfaces dummy dum1", nil)
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)
```

### Configure, then Show Multivalue Object

```go

    options := RetrieveOptions{
        Multivalue: true,
    }

    out, resp, err := c.Conf.Get(ctx, "interfaces dummy dum1", options)
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
```

### Configure, then Delete Object

```go

    out, resp, err := c.Conf.Delete(ctx, "interfaces dummy dum1")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
```

### Configure, then Save

```go

    out, resp, err := c.Conf.Save(ctx, "")

    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
```

### Configure, then Save File

```go

    out, resp, err := c.Conf.Save(ctx, "/config/test300.config")

    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
```

### Show Object

```go

    out, resp, err := c.Show.Do(ctx, "system image")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)
```

### Generate Object

```go

    out, resp, err := c.Generate.Do(ctx, "pki wireguard key-pair")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)
```

### Reset Object

```go

    out, resp, err := c.Reset.Do(ctx, "ip bgp 192.0.2.11")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)
    
```

### Configure, then Load File

```go

    out, resp, err := c.ConfigFile.Load(ctx, "/config/test300.config")
```

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
