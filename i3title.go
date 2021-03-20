package main

import (
	"flag"
	"fmt"

	"go.i3wm.org/i3/v4"
)

func main() {
	var truncate = flag.Int("t", 0, "truncate symbols")
	flag.Parse()

	t, _ := i3.GetTree()
	if f := get_focused_win(t.Root); f != nil {
		print_win_title(f.WindowProperties.Title, *truncate)
	}

	recv := i3.Subscribe(i3.WindowEventType, i3.WorkspaceEventType)
	for recv.Next() {
		switch e := recv.Event().(type) {
		case *i3.WindowEvent:
			title := ""

			if e.Change == "close" {
				if t, _ := i3.GetTree(); t.Root != nil {
					if f := get_focused_win(t.Root); f != nil {
						title = f.WindowProperties.Title
					}
				}
			} else {
				title = e.Container.WindowProperties.Title
			}
			print_win_title(title, *truncate)
		case *i3.WorkspaceEvent:
			if e.Change == "focus" {
				if w := get_focused_win(&e.Current); w == nil {
					print_win_title("", *truncate)
				}
			}
		}
	}
}

func print_win_title(title string, truncate int) {
	if truncate > 0 {
		runes := []rune(title)
		if len(runes) > truncate {
			title = fmt.Sprintf("%s...", string(runes[:truncate]))
		}
	}
	fmt.Printf("%s\n", title)
}

func get_focused_win(n *i3.Node) *i3.Node {
	return n.FindFocused(func(n *i3.Node) bool { return n.Focused && n.Type == "con" })
}
