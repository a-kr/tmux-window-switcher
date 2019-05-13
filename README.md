# tmux-window-switcher
A simple window selection menu to replace `Ctrl-B, w` functionality.

Recent versions of tmux (>= 2.6) changed the window switching hotkeys for window numbers 10 and above (see https://github.com/tmux/tmux/issues/1132): previously you could press `Ctrl-B, w, a` to switch to window number 10, but now you must press `Ctrl-B, w, M-a`. This is a major inconvenience for some people, me included.

This program provides a window switching menu with hotkeys working as before tmux 2.6 was released.

## Installation
```
go get github.com/a-kr/tmux-window-switcher/tmux_window_switcher
```
Copy the compiled binary from `$GOPATH/bin/tmux_window_switcher` to somewhere on your $PATH.

Also copy `$GOPATH/src/github.com/a-kr/tmux-window-switcher/tmux_window_selector_tmux` to somewhere on your $PATH.

Add the following line to your `~/.tmux.conf`:
```
bind-key -r -T prefix w run "tmux list-windows -F \"##I:##W\" | tmux_window_selector_tmux | cut -d \":\" -f 1 | xargs tmux select-window -t"
```
