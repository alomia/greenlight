# Greenlight API

Greenlight is a JSON API for retrieving and managing movie information. Similar to [Open Movie Database API](https://www.omdbapi.com/).

## Endpoint

| Method    | URL Pattern     | Handler            | Action                                 |
|-----------|-----------------|--------------------|----------------------------------------|
| GET       | /v1/healthcheck | healthcheckHandler | Show application information           |
| GET       | /v1/movies      | listMoviesHandler  | Show the details of all movies         |
| POST      | /v1/movies      | createMovieHandler | Create a new movie                     |
| GET       | /v1/movies/:id  | showMovieHandler   | Show the details of a specific movie   |
| PUT       | /v1/movies/:id  | updateMovieHandler | Update the details of a specific movie |
| DELETE    | /v1/movies/:id  | deleteMovieHandler | Delete a specific movie                |

## Tech Stack

**Server:** [Go](https://go.dev/), [httprouter](https://github.com/julienschmidt/httprouter)
