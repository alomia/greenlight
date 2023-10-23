# Greenlight API

The Greenlight API is a JSON API that allows you to retrieve and manage information about movies. It provides functionality similar to the Open Movie Database API. This README will guide you through the usage of the API and its endpoints.

## Endpoints

| Method | URL Pattern          | Handler             | Action                             |
|--------|----------------------|---------------------|-----------------------------------|
| GET    | /v1/healthcheck      | healthcheckHandler  | Show application information      |
| POST   | /v1/movies           | createMovieHandler  | Create a new movie                |
| GET    | /v1/movies/:id       | showMovieHandler     | Show the details of a specific movie |
| PUT    | /v1/movies/:id       | updateMovieHandler   | Update the details of a specific movie |
| DELETE | /v1/movies/:id       | deleteMovieHandler   | Delete a specific movie           |

Use these endpoints to interact with the Greenlight API and perform various actions related to movie information.

## Getting Started

To get started with the Greenlight API, follow these steps:

1. Clone the repository to your local machine.
2. Install the necessary dependencies.
3. Configure the application settings.
4. Run the application using your preferred method.

## API Usage

Make HTTP requests to the defined endpoints to interact with the Greenlight API. You can use tools like `curl`, Postman, or any programming language to make requests and handle responses.


## Tech Stack

**Server:** Go, httprouter


## Authors

- [@alomia](https://www.github.com/alomia)

