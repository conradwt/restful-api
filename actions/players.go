package actions

import (
	"github.com/conradwt/restful_api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// PlayersResource is the resource for the Player model
type PlayersResource struct {
	buffalo.Resource
}

// List gets all Players. This function is mapped to the path
// GET /players
func (v PlayersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	players := &models.Players{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Players from the DB
	if err := q.All(players); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, players))
}

// Show gets the data for one Player. This function is mapped to
// the path GET /players/{player_id}
func (v PlayersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Player
	player := &models.Player{}

	// To find the Player the parameter player_id is used.
	if err := tx.Find(player, c.Param("player_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, player))
}

// New renders the form for creating a new Player.
// This function is mapped to the path GET /players/new
func (v PlayersResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Player{}))
}

// Create adds a Player to the DB. This function is mapped to the
// path POST /players
func (v PlayersResource) Create(c buffalo.Context) error {
	// Allocate an empty Player
	player := &models.Player{}

	// Bind player to the html form elements
	if err := c.Bind(player); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(player)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, player))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Player was created successfully")

	// and redirect to the players index page
	return c.Render(201, r.Auto(c, player))
}

// Edit renders a edit form for a Player. This function is
// mapped to the path GET /players/{player_id}/edit
func (v PlayersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Player
	player := &models.Player{}

	if err := tx.Find(player, c.Param("player_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, player))
}

// Update changes a Player in the DB. This function is mapped to
// the path PUT /players/{player_id}
func (v PlayersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Player
	player := &models.Player{}

	if err := tx.Find(player, c.Param("player_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Player to the html form elements
	if err := c.Bind(player); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(player)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, player))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Player was updated successfully")

	// and redirect to the players index page
	return c.Render(200, r.Auto(c, player))
}

// Destroy deletes a Player from the DB. This function is mapped
// to the path DELETE /players/{player_id}
func (v PlayersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Player
	player := &models.Player{}

	// To find the Player the parameter player_id is used.
	if err := tx.Find(player, c.Param("player_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(player); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Player was destroyed successfully")

	// Redirect to the players index page
	return c.Render(200, r.Auto(c, player))
}
