# GoProjectCreator

CLI Tool with Cobra, to create standard Go Projects inspired by https://github.com/Melkeydev/go-blueprint


## Installation
go install github.com/A-Oez/GoProjectCreator@latest


## Example Command
With Postgres database structure: 
- GoProjectCreator create `--p` PROJECTNAME `--db` Postgres `--code` 

Basic project structure: 
- GoProjectCreator create `--p` PROJECTNAME `--code`

## Flags 
- `--p`: Specifies the name of the project `(mandatory)`
- `--db`: Creates a database structure `(optional)`
-         Currently available: Postgres
- `--code`: Open project in code editor `(optional)`


## Project structure
The project structure will look like this:

ProjectName
- `cmd`
- `config`
- `internal`
- `pkg`
- `go.mod`
- `main.go`
- `README.md`


## Template Files
In order for the structure to be generated correctly, the "gocreate_config" folder must be stored in the go/bin directory
- C:\Users\USER\go\bin