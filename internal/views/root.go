package views

import (
	"errors"
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	viewArr = []string{"v1", "v2", "v3", "v4"}
	active  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]
	out, err := g.View("v2")

	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func Root() {
	// init GUI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.InputEsc = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// test
	g.SetManagerFunc(layout)
	//real
	//if err := layout(g); err != nil {
	//	log.Fatalln(err)
	//}
	err = initKeyBinding(g)
	if err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Println("Error")
		log.Fatalln(err)
	}
}

// DEF
//func layout(g *gocui.Gui) error {
//	g.SetManagerFunc(func(g *gocui.Gui) error {
//		//if err := navigationView(g); err != nil {
//		//	return err
//		//}
//		return nil
//	})
//
//	return nil
//}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, maxX/2-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("V1")
			return err
		}
		v.Title = "v1 (editable)"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v2", maxX/2-1, 0, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("V2")
			return err
		}
		v.Title = "v2"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v3", 0, maxY/2-1, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("V3")
			return err
		}
		v.Title = "v3"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("v4", maxX/2, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("V4")
			return err
		}
		v.Title = "v4 (editable)"
		v.Editable = true
	}
	return nil
}
