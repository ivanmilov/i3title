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
	if f := t.Root.FindFocused(func(n *i3.Node) bool { return n.Type == i3.Con && len(n.Nodes) == 0 }); f != nil {
		print_win_title(f.WindowProperties.Title, *truncate)
	}

	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WindowEvent)
		if ev.Change != "focus" && ev.Change != "title" {
			continue
		}
		title := ev.Container.WindowProperties.Title

		print_win_title(title, *truncate)
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
