// Go Workshop - Language - Goroutines - Key/Value Database Service

package kvdb

import "context"

// keyValue is internally used to transport the pairs of
// key and value. It also transports a reply channel for
// returned keys and values.
type keyValue struct {
	key    string
	value  interface{}
	replyc chan keyValue
}

// Iterator describes functions which are performed for all
// key/value pairs of the database in the Do method.
type Iterator func(key string, value interface{})

// KVDB is an in-memory key/value database usable by multiple
// goroutines simultaneously.
type KVDB struct {
	ctx  context.Context
	data map[string]interface{}
	putc chan keyValue
	getc chan keyValue
	doc  chan Iterator
}

// New returns a new database instance.
func New(ctx context.Context) *KVDB {
	db := &KVDB{
		ctx:  ctx,
		data: make(map[string]interface{}),
		putc: make(chan keyValue),
		getc: make(chan keyValue),
		doc:  make(chan Iterator),
	}

	go db.backend()

	return db
}

// Put adds a key/value pair to the database and returns a
// possible already set value.
func (db *KVDB) Put(key string, value interface{}) interface{} {
	// Create the transport struct including a channel for the reply.
	// Reply channel has buffer size 1 to make backend non-blocking.
	kvIn := keyValue{
		key:    key,
		value:  value,
		replyc: make(chan keyValue, 1),
	}

	// Send request, wait for reply and return it.
	db.putc <- kvIn

	kvOut := <-kvIn.replyc

	return kvOut.value
}

// Get returns the value for a key.
func (db *KVDB) Get(key string) interface{} {
	// Create the transport struct without a value.
	kvIn := keyValue{
		key:    key,
		replyc: make(chan keyValue),
	}

	// Send request, wait for reply and return it.
	db.getc <- kvIn

	kvOut := <-kvIn.replyc

	return kvOut.value
}

// Do executes the given iterator function for each
// key/value pair.
func (db *KVDB) Do(do Iterator) {
	db.doc <- do
}

// backend is the groutine running the database and
// ensuring only one access per time.
func (db *KVDB) backend() {
	// Endless loop for all the requests.
	for {
		select {
		case <-db.ctx.Done():
			// The signal to end.
			return
		case req := <-db.putc:
			// Put request.
			current := db.data[req.key]
			db.data[req.key] = req.value
			reply := keyValue{
				key:   req.key,
				value: current,
			}
			req.replyc <- reply
		case req := <-db.getc:
			// Get request.
			reply := keyValue{
				key:   req.key,
				value: db.data[req.key],
			}
			req.replyc <- reply
		case do := <-db.doc:
			for key, value := range db.data {
				do(key, value)
			}
		}
	}
}
