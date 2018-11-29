# kube-tmux

## Installation

```
go get -u github.com/marcsauter/kube-tmux
```

## Configuration

Add `kube-tmux` to your status bar in `.tmux.conf`:
```
set -g status-right "#[fg=black] #(kube-tmux) #H %H:%M %d.%m.%Y"
```
or with a custom template:
```
set -g status-right "#[fg=black] #(kube-tmux '<{{.Context}}::{{.Namespace}}>') #H %H:%M %d.%m.%Y"
```

Consider reducing the status-interval for more current information:
```
set -g status-interval 5
```


