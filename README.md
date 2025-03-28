# todo_app_in_golang

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe 
Tri will help you get more done in less time.
It's designed to be as simple as possible to help you accomplish your goals.

Usage:
  tri [command]

Available Commands:
  add         Add a new todo
  completion  Generate the autocompletion script for the specified shell
  done        Mark Item as Done
  help        Help about any command
  list        Display all the tasks

Flags:
      --config string   config file (default is $HOME/.tri.yaml)
  -h, --help            help for tri
  -t, --toggle          Help message for toggle

Use "tri [command] --help" for more information about a command.
PS C:\Users\gusta\go\src> 
```

## Adding a task

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe add "Study Golang"
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
2025/03/23 21:17:45 unexpected end of JSON input
PS C:\Users\gusta\go\src> 
```

## Adding another task

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe add "Wash the dishes"
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
PS C:\Users\gusta\go\src> 
```

## List tasks

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe list                 
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
1.       Study Golang
2.       Wash the dishes
PS C:\Users\gusta\go\src> 
```

## Mark a tasks as done

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe done 1
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
"Study Golang" marked done
PS C:\Users\gusta\go\src> 
```

## Lists completed tasks

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe list --done
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
2. X     Study Golang
PS C:\Users\gusta\go\src> 
```

## List all tasks

```
PS C:\Users\gusta\go\src> .\github.com\gustavos86\tri\tri.exe list --all 
Using config file: C:\Users\gusta\.tri.yaml
Storage JSON file: C:\Users\gusta\.tridos.json
1.       Wash the dishes
2. X     Study Golang
PS C:\Users\gusta\go\src> 
```

# Example of YAML configuation file

.tri.yaml

```
datafile: C:\Users\gusta\.tridos.json
```

Data is stored in a JSON file.