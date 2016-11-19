# golyrics-api
A dead simple HTTP API for [golyrics](https://github.com/mamal72/golyrics).


## Installation

```bash
go get github.com/mamal72/golyrics-api
```


## Usage

```bash
golyrics-api
```

*You can also set ENV variables or run the app like this to change host or port:*
```bash
HOST="localhost" PORT="5000" golyrics-api
```

Now you can send simple http requests to `http://HOST:PORT/search/:query`.

Replace your search query with `:query` part. The recommended style for query part is `artist:name` of track.


## Development

1. Clone the repository.

1. Edit the example config ENV config in `.env.example` file and add `PORT` and `HOST`.

1. Rename `.env.example` to `.env`.

1. Build or run the app.


## Ideas || Issues
Just fill an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. :heart:


## License
> MIT