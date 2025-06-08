[![ci](https://github.com/lu-sd/learn-cicd-starter/actions/workflows/ci.yml/badge.svg?branch=main&event=pull_request)](https://github.com/lu-sd/learn-cicd-starter/actions/workflows/ci.yml)
# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

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

