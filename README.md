# GO Hive Metastore Client

This is an adjusted version of the [Hive metastore client library](https://github.com/akolb1/gometastore/tree/master/hmsclient) for Golang.

The main differences from the original package are:
* use of golang context to timeout calls
* sets an AuthToken via the Thrift `SetUGI` method for supported OIDC. This is a dirty hack for in-house use only!
 


## Installation

Standard `go get`:

```
$ go get github.com/jahnestacado/gometastore
```

## Example usage:

    import	(
        "log"
        "github.com/jahnestacado/gometastore"
        "time"
    )
    
    func printDatabases(accessToken string, connectionTimeout time.Duration) {
        options = gometastore.Options{AuthToken: accessToken, ConnectTimeout: &connectionTimeout}
        client, err := gometastore.Open("localhost", 9083, &options)
        if err != nil {
            log.Fatal(err)
        }
        defer client.Close()
        databases, err := client.GetAllDatabases(context.Background())
        if err != nil {
            log.Fatal(err)
        }
        for _, d := range databases {
            fmt.Println(d)
        }
    }
