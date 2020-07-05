package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"sloth/ent"
	"sloth/ent/migrate"
	"time"
)

func main() {
	client, err := ent.Open("sqlite3", "file:slodb?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	client = client.Debug()
	defer client.Close()
	// run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//generator := infra.NewULIDGenerator()

	owner := client.User.Create().
		SetUsername("admin").
		SetPassword("admin").
		SaveX(context.Background())


	client.User.Create().
		SetUsername("user").
		SetPassword("user").
		SaveX(context.Background())

	a := client.Activity.Create().
		SetEndedAt(time.Now()).
		SetStartedAt(time.Now()).
		SetOwner(owner).
		SaveX(context.Background())

	//log.Print("\n\n\n")
	//activities := client.Activity.Query().AllX(context.Background())
	//users := client.User.Query().AllX(context.Background())
	//log.Printf("%v", activities)
	//log.Printf("%v", users)

	log.Printf("\n\n\n %v \n\n\n ", a)

	x := client.Activity.
		Query().
		WithOwner().
		//Where(activity.ID(a.ID)).
		AllX(context.Background())

	log.Printf("%v", x)
}
