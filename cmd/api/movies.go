package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alomia/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Pretty Woman",
		Runtime:   119,
		Genres:    []string{"Comedy", "Romance"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movies": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
