# GoProjectCreator

CLI Tool with Cobra, to create standard Go Projects inspired by https://github.com/Melkeydev/go-blueprint


## Installation
go install github.com/A-Oez/GoProjectCreator@latest


## Example Command
GoProjectCreator create `--dir` PROJECTNAME `--rmd` `--code` 


## Flags 
- `--dir`: Specifies the name of the project `(mandatory)`
- `--rmd`: Generates a README.md file `(optional)`
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
Currently, the main.go file & the README.md can be filled with standard values.

To do this:
- Create a folder named "gocreate_config" in the go/bin directory
- Add `config_main.go.txt` & `config_README.md.txt` files with specific values
- `(see the gocreate_config directory in the project as an example)`