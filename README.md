# fishistozsh
## Description
`fishistozsh` is a small tool for reading [Fish Shell](https://fishshell.com/)'s
history file and displaying it in format of [ZSH](http://zsh.sourceforge.net/) history.

## Quickstart
- Download release or `go build` the source to obtain the binary file
- To parse fish history and display it run `fishistozsh`
``` sh
./fishistozsh
```

- To also add it to `Z Shell`'s history file just redirect the output
  to your history file, for example (__THIS WILL OVERWRITE CURRENT HISTORY__):
``` sh
./fishistozsh > ~/.zsh_history
```
