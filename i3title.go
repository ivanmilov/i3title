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
	if f := get_focused(t.Root); f != nil {
		print_win_title(f.WindowProperties.Title, *truncate)
	}

	recv := i3.Subscribe(i3.WindowEventType, i3.WorkspaceEventType)
	for recv.Next() {
		switch e := recv.Event().(type) {
		case *i3.WindowEvent:
			title := e.Container.WindowProperties.Title
			print_win_title(title, *truncate)
		case *i3.WorkspaceEvent:
			if get_focused(&e.Current) == nil {
				if e.Change == "rename" || e.Change == "focus" {
					print_win_title("", *truncate)
				}
			}
		}
	}
}

func print_win_title(title string, truncate int) {
	if truncate > 0 {
		if len(title) > truncate {
			title = fmt.Sprintf("%s...", title[:truncate])
		}
	}
	fmt.Printf("%s\n", title)
}

func get_focused(n *i3.Node) *i3.Node {
	return n.FindFocused(func(n *i3.Node) bool { return n.Type == i3.Con && len(n.Nodes) == 0 })
}
