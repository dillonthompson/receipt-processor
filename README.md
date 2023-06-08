# receipt-processor

## Development
In order to develop on this applicationa and utilize hot-reloading, I used a tool called [Air](https://github.com/cosmtrek/air). In order to use this tool and run this application with hot reloading available, run `go install github.com/cosmtrek/air@latest` and then run `air` in the root directory of the project.

## Testing
I wrote some small unit tests for the processing functionality. In order to run these tests, simply run `go test -v ./...`

## Building and running with Docker
I also wrote a multi-stage Dockerfile for building, linting, testing, and running the application. In order to build this Dockerfile, run the command `docker build . -t receipt-processor` in the root directory of the project. To then run the project using Docker, run `docker run -p 8080:8080 receipt-processor`