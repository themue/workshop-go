// Go Workshop - Language - Goroutines - Key/Value Database Service

package kvdb

import "context"

// Iterator describes functions which are performed for all
// key/value pairs of the database in the Do method.
type Iterator func(key string, value interface{})

// KVDB is an in-memory key/value database usable by multiple
// goroutines simultaneously.
type KVDB struct {
	ctx     context.Context
	data    map[string]interface{}
	actionc chan func()
}

// New returns a new database instance.
func New(ctx context.Context) *KVDB {
	db := &KVDB{
		ctx:     ctx,
		data:    make(map[string]interface{}),
		actionc: make(chan func()),
	}

	go db.backend()

	return db
}

// Put adds a key/value pair to the database and returns a
// possible already set value.
func (db *KVDB) Put(key string, value interface{}) interface{} {
	var current interface{}
	db.actionc <- func() {
		current = db.data[key]
		db.data[key] = value
	}
	return current
}

// Get returns the value for a key.
func (db *KVDB) Get(key string) interface{} {
	var current interface{}
	db.actionc <- func() {
		current = db.data[key]
	}
	return current
}

// Do executes the given iterator function for each
// key/value pair.
func (db *KVDB) Do(do Iterator) {
	db.actionc <- func() {
		for key, value := range db.data {
			do(key, value)
		}
	}
}

// backend is the groutine running the database and
// ensuring only one access per time.
func (db *KVDB) backend() {
	for {
		select {
		case <-db.ctx.Done():
			// The signal to end.
			return
		case action := <-db.actionc:
			// Execute the action.
			action()
		}
	}
}
