# GoProjectCreator

CLI TUI Tool with Cobra, to create standard Go Projects inspired by https://github.com/Melkeydev/go-blueprint
TUI made with https://github.com/pterm/pterm 

## Installation
go install github.com/A-Oez/GoProjectCreator@latest

## Example Command
Open TUI: 
- GoProjectCreator tui

With Postgres database structure: 
- GoProjectCreator create `--p` PROJECTNAME `--db` Postgres `--code` 

Basic project structure: 
- GoProjectCreator create `--p` PROJECTNAME `--code`

## Flags 
- `--p`: Specifies the name of the project `(mandatory)`
- `--db`: Creates a database structure `(optional)`
          Currently available: Postgres
- `--code`: Open project in code editor `(optional)`