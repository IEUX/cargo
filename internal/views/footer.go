package views

import (
	"cargo/internal"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/shirou/gopsutil/v4/mem"
	"time"
)

func footer(g *gocui.Gui) error {
	//
	//
	maxX, maxY := g.Size()
	//Title view
	if v, err := g.SetView("Title", 0, maxY-3, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		go sysUsage(v)
	}
	//Made by view
	if v, err := g.SetView("madeBy", maxX-maxX/6, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(v, "> made by")
		fmt.Fprintf(v, internal.Purple, " I3UX_")
	}
	return nil
}

func sysUsage(v *gocui.View) {
	m, _ := mem.VirtualMemory()
	for true {
		v.Clear()
		time.Sleep(1 * time.Second)
		fmt.Fprintf(v, internal.WarningColor, "Cargo ")
		fmt.Fprint(v, " | ")
		fmt.Fprintf(v, internal.InfoColor, fmt.Sprintf(" Mem: %.2f%%", m.UsedPercent))
	}
}
