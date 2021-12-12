package database

import (
	"fmt"
	"sync"
	"time"
)

type DBConnection struct {
	ShakeKey string
}

var singleton *DBClient
var once sync.Once

type DBClient struct {
	connections sync.Pool
}

func GetDBClient() *DBClient {
	once.Do(func() {
		singleton = new(DBClient)
		singleton.AppendConnection(5)
	})
	return singleton
}

func (d *DBClient) Query(query string) {
	conn := d.GetConnection()
	fmt.Printf("Running query %s => from client %p with connection %p\n", query, d, conn)
	time.Sleep(1)
	d.CloseConnection(conn)
}

func (d *DBClient) GetConnection() *DBConnection {
	get := d.connections.Get().(*DBConnection)
	return get
}

func (d *DBClient) AppendConnection(n int) {
	for {
		if n == 0 {
			break
		}
		d.connections.Put(&DBConnection{ShakeKey: fmt.Sprintf("%d", time.Now().Unix())})
		n--
	}
}

func (d *DBClient) CloseConnection(conn *DBConnection) {
	d.connections.Put(conn)
}
