# Development Principles

## Architecture & Principles
### SOLID
We are using SOLID concept here

## Variables
### Naming
  - b: body
  - cl: client
  - col: collection
  - con: connexion
  - cr: create
  - ctx: context
  - h: handler(s)
  - l: log / logger
  - nw: new
  - mg: mongo
  - r: router
  - re: reader
  - req: requestion
  - res: response
  - s: serve / server
  - sig: signal
  - str: string
  - v: valid / validate
  - w: writer

## Packages
### Imports naming
  - database: DB
  - handlers: H
  - models: M
  - router: R
  - services: S
  - utils: U

## Folders
### _v1
Routes and Handler of the first version of the service
Http related files

### build
All deployement related files

### cmd/service-users-data
Keep `main.go` file

### internals
All internal logic files (workers, db-services, etc.)

### testing
All tests related files
