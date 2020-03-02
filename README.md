# jotfly-server
Run anywhere server for Jotfly notes.

Jotfly-server can be self-hosted or run anywhere as a container for storing and retrieving notes and lists.
Currently notes saving is entirely in JSON on individual files saved in user-created directories.

This server is setup with clear text files, for simplicity, portability, and migration of notes.
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