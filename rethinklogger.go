package rethinklogger

/**
 * Created by M on 19/08/17. With ❤
 * Go package to persist all logs from stdio and stderr to RethinkDB.
 */

/*
 Author - Mustansir Zia.
*/

import (
	"strings"
	"time"

	"github.com/segmentio/go-shipit"
	r "gopkg.in/dancannon/gorethink.v2"
)

const (
	dbName     = "rethinkLogs"
	dbTable    = "logs"
	dbUser     = "rethinkLogger"
	dbPassword = "rethinkLogger"
)

// Start - Clients call this to start the persisting of logs.
func Start(dbAddress, adminPassword string) error {
	return StartWithBuffer(dbAddress, adminPassword, 1)
}

// StartWithBuffer - Clients call this to start the persisting of logs
// with a customizable buffer size.
func StartWithBuffer(dbAddress, adminPassword string, bufferSize int) error {
	session, err := createSession(dbAddress, adminPassword)
	if err != nil {
		return err
	}
	writer := &rethinkWriter{
		session: session,
		db:      r.DB(dbName).Table(dbTable),
		// Let's persist when there is at least a single log in the buffer.
		bufferSize: bufferSize,
		// Keep persisting logs at 5 second intervals.
		flushInterval: 5 * time.Second,
		buffer:        make([]RethinkLog, 0),
	}
	go writer.start()
	// Start the pipeline of logs to our custom writer.
	return shipit.To(writer)
}

func createSession(dbAddress string, adminPassword string) (*r.Session, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:    dbAddress,
		Database:   dbName,
		Username:   dbUser,
		Password:   dbPassword,
		MaxOpen:    10,
		InitialCap: 10,
	})
	if err != nil {
		// Really hoped if there was an explicit error type here.
		if strings.Contains(err.Error(), "Unknown user") {
			// This is the first run of the server. Database is not initialized.
			// Let's create the database and insert the user.
			if err = createDB(dbAddress, adminPassword); err != nil {
				return nil, err
			}
			// Create the session again now that the DB's ready.
			return createSession(dbAddress, adminPassword)
		}
		return nil, err
	}
	return session, nil
}

func createDB(address, adminPassword string) error {
	session, err := r.Connect(r.ConnectOpts{
		Address:  address,
		Username: "admin",
		Password: adminPassword,
	})
	if err != nil {
		return err
	}
	defer session.Close()
	// Create database.
	if _, err := r.DBCreate(dbName).RunWrite(session); err != nil {
		return err
	}
	// Create database user.
	if _, err := r.DB("rethinkdb").Table("users").Insert(map[string]interface{}{"id": dbUser, "password": dbPassword}).RunWrite(session); err != nil {
		return err
	}
	// Grant all permissions of database to user.
	if _, err := r.DB(dbName).Grant(dbUser, map[string]interface{}{"read": true, "write": true, "config": true}).RunWrite(session); err != nil {
		return err
	}
	// Create all tables of database.
	return createTables(session)
}

func createTables(session *r.Session) error {
	if _, err := r.Do(
		r.DB(dbName).TableCreate(dbTable),
	).RunWrite(session); err != nil {
		return err
	}
	return nil
}
