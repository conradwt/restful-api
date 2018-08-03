package grifts

import (
	"fmt"
	"log"

	"github.com/conradwt/restful_api/models"
	"github.com/gobuffalo/pop"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		// reset DB table(s)
		DB, error := pop.Connect("development")
		if error != nil {
			log.Panic(error)
		}

		if error = PlayerDestroyAll(); error != nil {
			log.Panic(error)
		}

		// create players
		john := models.Player{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Age:       25,
			Position:  "forward"}
		_, error = DB.ValidateAndSave(&john)
		if error != nil {
			log.Panic(error)
		}

		jane := models.Player{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@example.com",
			Age:       23,
			Position:  "goalkeeper"}
		_, error = DB.ValidateAndSave(&jane)
		if error != nil {
			log.Panic(error)
		}

		return nil
	})

})

// PlayerDestroyAll removes all the player records from the DB.
func PlayerDestroyAll() error {
	DB, error := pop.Connect("development")
	if error != nil {
		log.Panic(error)
	}

	players := models.Players{}
	error = DB.All(&players)
	if error != nil {
		fmt.Print("PlayerDestroyAll ERROR!\n")
		fmt.Printf("%v\n", error)
		return error
	}

	for i := 0; i < len(players); i++ {
		player := players[i]
		if error = DB.Destroy(&player); error != nil {
			log.Panic(error)
		}
	}

	return nil
}
