# GoProjectCreator

**GoProjectCreator** is a CLI TUI tool created with [Cobra](https://github.com/spf13/cobra) to generate standard Go projects. Inspired by [go-blueprint](https://github.com/Melkeydev/go-blueprint), it provides an easy way to create projects with predefined structures and configurations. The TUI (Text User Interface) is built using [pterm](https://github.com/pterm/pterm).

## Installation
`go install github.com/A-Oez/GoProjectCreator@main`

## Example Command
To open the TUI:
- GoProjectCreator tui

To create a project with a Postgres database structure:
- GoProjectCreator create `--p` PROJECTNAME `--db` Postgres `--code` 

To create a basic project:
- GoProjectCreator create `--p` PROJECTNAME `--code`

## Flags 
- `--p`: Specifies the name of the project `(mandatory)`
- `--db`: Creates a database structure `(optional)`
          Currently available: Postgres
- `--code`: Open project in code editor `(optional)`