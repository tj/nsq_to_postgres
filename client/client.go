package client

import "github.com/jmoiron/sqlx"
import "github.com/lib/pq"
import "fmt"
import "log"

// Client connection.
type Client struct {
	*Config
	*sqlx.DB
	stmt *sqlx.Stmt
}

// New client with the given `config`.
func New(config *Config) (*Client, error) {
	c := &Client{
		Config: config,
	}

	if c.Verbose {
		log.Printf("connecting to %s (max connections: %d)", c.Connection, c.MaxOpenConns)
	}

	err := c.connect()
	return c, err
}

// Establish connection.
func (c *Client) connect() error {
	db, err := sqlx.Connect("postgres", c.Connection)
	if err != nil {
		return err
	}
	c.DB = db
	db.SetMaxOpenConns(c.MaxOpenConns)

	stmt, err := c.Preparex(fmt.Sprintf(`insert into %s values ($1)`, c.Table))
	if err != nil {
		return err
	}
	c.stmt = stmt

	return nil
}

// Bootstrap:
//
// - create "events" table
//
func (c *Client) Bootstrap() error {
	return c.CreateEventsTable()
}

// CreateEventsTable creates the "events" table and noop when it exists.
func (c *Client) CreateEventsTable() error {
	_, err := c.Exec(fmt.Sprintf(`create table %s (%s jsonb)`, c.Table, c.Column))

	if err, ok := err.(*pq.Error); ok {
		if err.Code == "42P07" {
			return nil
		}
	}

	return err
}

// Insert event json blob.
func (c *Client) Insert(e []byte) error {
	if c.Verbose {
		log.Printf("inserting %s", e)
	}

	_, err := c.stmt.Exec(e)
	return err
}
