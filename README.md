# Adoption???

This tool is for tracking levels of adoption of adoptables,
practices or change that you want to track

The intent of this tools is to understand the impact
of the adoptables which support or enabling what you are pursuing
and also have awareness that the work doesn't end when the
implementation is up and running, but when the implementation
is been used by different people.

## Endpoints

* GET /teams
* POST /teams
* GET /teams/{id}/people
* POST /teams/{id}/people
* GET /people
* POST /people
* POST /people/{id}/adoptables

## Build the Binary

```sh
go build -o adoption main.go
```

## Run the server

```sh
./adoption server
```
