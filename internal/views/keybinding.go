package views

import "github.com/jroimartin/gocui"

func initKeyBinding(g *gocui.Gui) error {
	if err := kbGlobal(g); err != nil {
		return err
	}
	return nil
}

func kbGlobal(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
