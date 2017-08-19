# go-rethinklogger

#### go-rethinklogger persists all your logs from `stdio` and `stderr` to [RethinkDB](http://rethinkdb.com/).

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
    // Start takes a rethinkDB address and the admin password
    // as arguments. Don't worry though, admin user is only
    // used to create the database and a user.
    rethinklogger.Start("localhost:28015", "adminPassword")
    fmt.Fprintln(os.Stdout, "Sample stdio log!")
    fmt.Fprintln(os.Stderr, "Sample stderr log!")

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
                Log = "Sample stdio log!"
                CreatedAtHuman = "17 Aug 17 15:04 +530"
                CreatedAt = 1503162020
            }
            log {
                Log = "Sample stderr log!"
                CreatedAtHuman = "17 Aug 17 15:04 +530"
                CreatedAt = 1503162020
            }
        */

    }



```


## Sidenotes.