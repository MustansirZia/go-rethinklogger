package rethinklogger

/**
 * Created by M on 19/08/17. With ❤
 */

import (
	"strings"
	"sync"
	"time"

	r "gopkg.in/dancannon/gorethink.v2"
)

type rethinkWriter struct {
	// Has Rethink's connection pool.
	session *r.Session

	// Query for Rethink's table.
	db r.Term

	// Buffer size to be kept before dumping logs to DB. [100]
	bufferSize int

	// Bytes buffer.
	buffer []RethinkLog

	// Interval before buffer is dumped inside DB. [5 secs]
	flushInterval time.Duration

	// Locks for concurrent use.
	sync.Mutex
}

// RethinkLog - struct to a hold a single rethink log.
type RethinkLog struct {
	Log            string `gorethink:"Log,omitempty"`
	CreatedAt      int64  `gorethink:"CreatedAt,omitempty"`
	CreatedAtHuman string `gorethink:"CreatedAtHuman,omitempty"`
}

func (writer *rethinkWriter) Write(p []byte) (n int, err error) {
	buffer := RethinkLog{
		Log:       string(p),
		CreatedAt: time.Now().Unix(),
	}
	writer.Lock()
	writer.buffer = append(writer.buffer, buffer)
	writer.Unlock()
	return len(p), nil
}

func (writer *rethinkWriter) dump() {
	if len(writer.buffer) >= writer.bufferSize {
		writer.db.Insert(writer.buffer).RunWrite(writer.session)
		writer.Lock()
		writer.buffer = make([]RethinkLog, 0)
		writer.Unlock()
	}
}

func (writer *rethinkWriter) start() {
	for {
		time.Sleep(writer.flushInterval)
		writer.dump()
	}
}

// QueryLogs - Only exported function apart from Start() and StartWithBuffer(). Used to Query past logs.
func QueryLogs(dbAddress string, from, to time.Time) ([]RethinkLog, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  dbAddress,
		Database: dbName,
		Username: dbUser,
		Password: dbPassword,
	})
	if err != nil {
		// If rethinklogger.start() wasn't called,
		// just return an empty slice.
		if strings.Contains(err.Error(), "Unknown user") {
			return nil, nil
		}
		return nil, err
	}
	defer session.Close()
	filterQuery := r.Row.
		Field("CreatedAt").
		Ge(from.Unix()).
		And(r.Row.Field("CreatedAt")).
		Le(to.Unix())
	cursor, err := r.DB(dbName).
		Table(dbTable).
		Filter(filterQuery).
		OrderBy(r.Desc("CreatedAt")).
		Run(session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var logs []RethinkLog
	if err := cursor.All(&logs); err != nil {
		return nil, err
	}
	for i := range logs {
		logs[i].CreatedAtHuman = time.Unix(logs[i].CreatedAt, 0).Format(time.RFC822Z)
	}
	return logs, nil
}
