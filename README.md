# go-rethinklogger

go-rethinklogger persists all your logs from `stdio` and `stderr` to [RethinkDB](http://rethinkdb.com/).

• Can be used to monitor logs and analytics of your GO application remotely.
<br />

• It works and is compatible with literally all types of loggers. The API is exceptionally simple and you just need to do normal logging, 
it automatically reads from `stdio` and `stderr`, persists logs for later use and echoes the same back to their respective streams.

## Installation.

• Using Go.
```
go get github.com/mustansirzia/go-rethinklogger
```

• Using Glide.
```
glide get github.com/mustansirzia/go-rethinklogger
glide up
```

• Using go-vendor.
```
govendor fetch github.com/mustansirzia/go-rethinklogger
```

## Usage.

```go

    // Start persisting logs. Call preferably from init.
    rethinklogger.Start("localhost:28015", "")
    ...
    ...
    // Query logs from anywhere in your application.
    previousDay := time.Now().Add(0, 0, -1)
    now := time.Now()
    logs, err := rethinklogger.QueryLogs("localhost:28015", previousDay, now)
    if err != nil {
        fmt.Println(err.Error())
    }
    for _, log := range logs {

        /*
            log {
                Log = "Sample log"
                CreatedAtHuman = "17 Aug 17 15:04 +530"
                CreatedAt = 1503162020
            }
        */

    }


```


## Sidenotes.