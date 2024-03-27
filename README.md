# GoProjectCreator

CLI Tool with Cobra, to create standard Go Projects

## Installation
go install 

## Example Command
GoCreate create --dir PROJECTNAME --rmd --code 

## Flags 
- `--dir`: Specifies the name of the project (mandatory)
- `--rmd`: Generates a README.md file (optional)
- `--code`: Opens the new project in VS Code (optional)

## Project structure
The project structure will look like this:

ProjectName
├── cmd
├── config
├── internal
├── pkg
├── go.mod
├── main.go
├── README.md

## Template Files
Currently, the main.go file & the README.md can be filled with standard values.

To do this:
- Create a folder named "gocreate_config" in the go/bin directory
- Add `config_main.go.txt` & `config_README.md.txt` files with specific values