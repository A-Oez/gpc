# GoProjectCreator

CLI Tool with Cobra, to create standard Go Projects

## Installation
go install 

## Example Command
GoCreate create --dir PROJECTNAME --rmd --code 

## Flags 
--dir | name of project (mandatory)
-- rmd | generate README.md file (optional)
-- code | open new project in vs.code (optional)

## Project structure
ProjectName
|__ cmd
|__ config
|__ internal
|__ pkg
|__ go.mod
|__ main.go
|__ README.md

## Template Files
Currently, the main.go file & the README.md can be filled with standard values.

To do this:
-> create folder "gocreate_config" in go/bin directory
-> add config_main.go.txt & config_README.md.txt files with specific value 