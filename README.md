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

If you are using [kube-tmuxp](https://github.com/arunvelsriram/kube-tmuxp) to work with multiple clusters then `KUBECONFIG` environment variable will be set at tmux session level. In this case use:
```
set -g status-right '#[fg=black] #(echo "$(tmux show-environment KUBECONFIG) kube-tmux" | sh)'
```

Consider reducing the status-interval for more current information:
```
set -g status-interval 5
```


