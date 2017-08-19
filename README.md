# go-rethinklogger

Rethink logger persists all your logs from `stdio` and `stderr` to [RethinkDB](http://rethinkdb.com/). Can be used to monitor 
logs and analytics of your GO application remotely.
<br />
Works and is compatible with literally all types of loggers. The API is exceptionally simple and you just need to do normal logging, 
it automatically reads from `stdio` and `stderr`, persists logs for later use and echoes the same back to their respective streams.

### Installation 

• Using Go.
`go get github.com/mustansirzia/go-rethinklogger` 

• Using Glide.
`glide get github.com/mustansirzia/go-rethinklogger` 

• Using go-vendor.
`govendor fetch github.com/mustansirzia/go-rethinklogger` 