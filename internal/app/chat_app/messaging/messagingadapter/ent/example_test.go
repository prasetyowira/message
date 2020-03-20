// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleMessage() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the message's edges.

	// create message vertex with its edges.
	m := client.Message.
		Create().
		SetUID("string").
		SetText("string").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SaveX(ctx)
	log.Println("message created:", m)

	// query edges.

	// Output:
}