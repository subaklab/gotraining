// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/yOn3nZU5rf

// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// result is what is sent back from each operation.
type result struct {
	id  int
	op  string
	err error
}

// init called before main.
func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	performInserts()
}

// performInserts coordinates the possible inserts that need to take place.
func performInserts() {
	log.Println("Inserts Started")

	// Set the number of routines and inserts.
	const routines = 10
	const inserts = routines * 2

	// Buffered channel to receive information about any possible insert.
	ch := make(chan result, inserts)

	// Number of responses we need to handle.
	waitInserts := inserts

	// Perform all the inserts.
	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)

			// We don't need to wait to start the second insert
			// thanks to the buffered channel. The first send
			// will happen immediately.
			ch <- insertTrans(id)
		}(i)
	}

	// Process the insert results as they complete.
	for waitInserts > 0 {
		// Wait for a response from a goroutine.
		r := <-ch

		// Display the result.
		log.Printf("ID: %d OP: %s ERR: %v", r.id, r.op, r.err)

		// Decrement the wait count and determine if we are done.
		waitInserts--
	}

	log.Println("Inserts Complete")
}

// insertUser simulates a database operation.
func insertUser(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}

	return r
}

// insertTrans simulates a database operation.
func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert TRANS value (%d)", id),
	}

	// Randomize if the insert fails or not.
	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}

	return r
}
