# jotfly-server
Run anywhere server for Jotfly notes.

Jotfly-server is intended to be a notes server that can be run anywhere for storing and retrieving notes and lists.
Currently notes saving is entirely done in JSON text on individual files saved in user-created directories.

### Building

Go:  
`go build -o jotfly cmd/server/main.go`

Docker:  
`docker build -f Dockerfile -t jotfly .`

### Future
- gRPC gateway for REST, swagger docs(generated)
- Greater test coverage
- Saving notes and users in SQL
- Storing into rclone
- Directory sharing for multi-users
- Moving directories and files
- Notes diff updating