# kube-tmux

## build

```
GO111MODULE=on go build -o <some location in your $PATH>/kube-tmux
```

## tmux

Add `kube-tmux` to your status bar `.tmux.conf`:
```
set -g status-right "#[fg=black] #(kube-tmux) #H %H:%M %d.%m.%Y"
```

Consider reducing the status-interval for more current information:
```
set -g status-interval 5
```


