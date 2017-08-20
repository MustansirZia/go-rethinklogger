# go-rethinklogger

![GitHub tag](https://github.com/MustansirZia/go-rethinklogger/releases)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)


#### go-rethinklogger persists all your logs from `stdio` and `stderr` to [RethinkDB](http://rethinkdb.com/).

![RethinkDB Logo](https://rethinkdb.com/assets/images/docs/api_illustrations/quickstart.png "RethinkDB thinker on the job.")

• Can be used to monitor logs and analytics of your GO application remotely.
<br />

• It works and is compatible with literally all types of loggers. The API is exceptionally simple and you just need to do normal logging, 
it automatically reads from `stdio` and `stderr`, persists logs for later use and echoes the same back to their respective streams.

<br />

## Installation.

• Using Go.
```
go get github.com/MustansirZia/go-rethinklogger
```

• Using Glide.
```
glide get github.com/MustansirZia/go-rethinklogger
glide up
```

• Using go-vendor.
```
govendor fetch github.com/MustansirZia/go-rethinklogger
```

<br />

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

<br />

### Sidenotes.

• This library is built over [go-shipit](https://github.com/segmentio/go-shipit) which dispatches std logs to a `io.Writer` interface. This library is basically a writer for shipit.
<br />
<br />
• To avoid our database from getting overwhelmed by logs, logs are first accumulated inside a buffer and then dispatched at an interval of 5 secs.
<br />
<br />
• Logging can also be started using  `StartWithBuffer` function which takes an additional buffer size argument. This is the minimum number of logs that must be accumulated inside the buffer before all the logs are dispatched to Rethink.
By default the value is `1`.

<br />

### Inspiration.
=> <b>go-loggly</b> - https://github.com/segmentio/go-loggly
<br />
<br />
=> <b>Winston for nodejs</b> - https://github.com/winstonjs/winston

<br />

### License.
See [License](https://github.com/MustansirZia/go-rethinklogger/blob/master/LICENSE.txt).