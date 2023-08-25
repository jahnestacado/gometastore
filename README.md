# GO Hive Metastore Client

This is an adjusted version of the [Hive metastore client library](https://github.com/akolb1/gometastore/tree/master/hmsclient) for Golang

## Installation

Standard `go get`:

```
$ go get github.com/jahnestacado/gometastore/hmsclient
```

## Example usage:

    import	(
        "log"
        "github.com/jahnestacado/gometastore/hmsclient"
        "time"
    )
    
    func printDatabases(accessToken string, connectionTimeout time.Duration) {
        options = hmsclient.Options{AuthToken: accessToken, ConnectTimeout: &connectionTimeout}
        client, err := hmsclient.Open("localhost", 9083, &options)
        if err != nil {
            log.Fatal(err)
        }
        defer client.Close()
        databases, err := client.GetAllDatabases()
        if err != nil {
            log.Fatal(err)
        }
        for _, d := range databases {
            fmt.Println(d)
        }
    }
