# GoProjectCreator

**GoProjectCreator** is a CLI TUI tool created with [Cobra](https://github.com/spf13/cobra) to generate standard Go projects. Inspired by [go-blueprint](https://github.com/Melkeydev/go-blueprint), it provides an easy way to create projects with predefined structures and configurations. The TUI (Text User Interface) is built using [pterm](https://github.com/pterm/pterm).

![image](https://github.com/user-attachments/assets/93bf3a86-5618-4cbf-8ca5-1c1e7053fb2b)

## Installation
`go install github.com/A-Oez/gpc@main`

## Example Command
To open the TUI:
- gpc tui

To create a project with a Postgres database structure:
- gpc create `--p` PROJECTNAME `--db` Postgres `--code` 

To create a basic project:
- gpc create `--p` PROJECTNAME `--code`

## Flags 
- `--p`: Specifies the name of the project `(mandatory)`
- `--db`: Creates a database structure `(optional)`
          Currently available: Postgres
- `--code`: Open project in code editor `(optional)`
