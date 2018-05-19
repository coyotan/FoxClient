package FoxCore

import "github.com/jroimartin/gocui"

var (
  maxX int
  maxY int
)

func guiInit() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.InputEsc = true
	g.BgColor = gocui.ColorBlue
	g.SetManagerFunc(Layout)
}

func Layout(g *gocui.Gui) error {
 return nil
}