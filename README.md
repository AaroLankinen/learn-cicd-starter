# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Build Status
![Build Status](https://github.com/AaroLankinen/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)

## Codecov Coverage Status
![Coverage Status](https://codecov.io/gh/AaroLankinen/learn-cicd-starter/branch/main/badge.svg?token=K85C9E3T1B)

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

## Aaro Lankinen's version of Boot.dev's Notely app