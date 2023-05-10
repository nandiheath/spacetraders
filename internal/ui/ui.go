package ui

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"time"
//
//	ui "github.com/gizak/termui/v3"
//	"github.com/gizak/termui/v3/widgets"
//	"github.com/nandiheath/spacetraders/internal/api"
//	"github.com/nandiheath/spacetraders/internal/utils"
//)
//
//type FocusMode int
//
//const (
//	FocusInput         FocusMode = 1
//	FocusMainPanelList FocusMode = 2
//)
//
//type ListMode int
//
//const (
//	ListNone      ListMode = 0
//	ListContracts ListMode = 1
//)
//
//type Dashboard struct {
//	width     int
//	height    int
//	input     string
//	client    *api.APIClient
//	focusMode FocusMode
//	list      *widgets.List
//	listMode  ListMode
//}
//
//func NewDashboard() *Dashboard {
//	l := widgets.NewList()
//	l.SelectedRowStyle.Bg = ui.ColorCyan
//	return &Dashboard{
//		client:    utils.NewAPIClient(),
//		list:      l,
//		focusMode: FocusInput,
//		listMode:  ListNone,
//	}
//}
//
//func (d *Dashboard) Run() {
//	if err := ui.Init(); err != nil {
//		fmt.Errorf("failed to initialize termui: %v", err)
//		panic(err)
//	}
//	defer ui.Close()
//
//	uiEvents := ui.PollEvents()
//	ticker := time.NewTicker(time.Second).C
//	for {
//		select {
//		case e := <-uiEvents:
//			switch e.ID { // event string/identifier
//			case "<C-c>": // press 'q' or 'C-c' to quit
//				return
//			//case "<MouseLeft>":
//			//	payload := e.Payload.(ui.Mouse)
//			//	x, y := payload.X, payload.Y
//			case "<Resize>":
//				payload := e.Payload.(ui.Resize)
//				d.width, d.height = payload.Width, payload.Height
//			}
//			switch e.Type {
//			case ui.KeyboardEvent: // handle all key presses
//				d.processKeyboardInput(e.ID)
//				d.drawInput()
//			}
//		// use Go's built-in tickers for updating and drawing data
//		case <-ticker:
//			d.drawInput()
//		}
//	}
//}
//
//func getMainPanelRect() (int, int, int, int) {
//	return 0, 3, 80, 20
//}
//
//func (d *Dashboard) drawInput() {
//	p := widgets.NewParagraph()
//	p.Title = "Input"
//	if d.input == "" {
//		p.TextStyle = ui.NewStyle(ui.ColorWhite)
//		p.Text = ".."
//
//	} else {
//		p.TextStyle = ui.NewStyle(ui.ColorClear)
//		p.Text = d.input
//	}
//
//	p.SetRect(0, 0, 30, 3)
//	if d.focusMode == FocusInput {
//		p.BorderStyle.Fg = ui.ColorYellow
//	} else {
//		p.BorderStyle.Fg = ui.ColorWhite
//	}
//
//	ui.Render(p)
//}
//func (d *Dashboard) updateMainListPanel(focused bool) {
//	if focused {
//		d.focusMode = FocusMainPanelList
//		d.list.BorderStyle.Fg = ui.ColorYellow
//	} else {
//		d.focusMode = FocusInput
//		d.list.BorderStyle.Fg = ui.ColorWhite
//	}
//	d.drawInput()
//	ui.Render(d.list)
//}
//func (d *Dashboard) processKeyboardInput(input string) {
//	switch input {
//	case "<Down>":
//		d.list.ScrollDown()
//	case "<Up>":
//		d.list.ScrollUp()
//		d.updateMainListPanel(true)
//		break
//	}
//	switch input {
//	case "<Backspace>":
//		d.input = d.input[0 : len(d.input)-1]
//	case "<Escape>":
//		d.input = ""
//	case "<Enter>":
//		d.processCommand()
//		break
//
//	default:
//		if len(input) == 1 {
//			d.input = d.input + input
//			//d.updateMainListPanel(false)
//		}
//	}
//}
//
//func (d *Dashboard) processCommand() {
//	// handles input text only if input box is focused
//	if d.focusMode == FocusInput {
//		switch d.input {
//		case ":info":
//			d.showInfo()
//		case ":contract":
//			d.showContracts()
//			d.input = ""
//			break
//		default:
//			d.input = "invalid input"
//		}
//	} else if d.focusMode == FocusMainPanelList {
//		switch d.listMode {
//		case ListNone:
//			d.showError(errors.New("invalid list mode"), "List")
//			break
//		case ListContracts:
//			break
//		}
//	}
//
//}
//
//func (d *Dashboard) showError(e error, title string) {
//	p := widgets.NewParagraph()
//	p.Title = title
//	x1, y1, x2, y2 := getMainPanelRect()
//	p.SetRect(x1, y1, x2, y2)
//	p.Text = e.Error()
//	p.BorderStyle.Fg = ui.ColorRed
//}
//
//func (d *Dashboard) showInfo() {
//	ctx := context.Background()
//	req := d.client.AgentsApi.GetMyAgent(ctx)
//	r, _, err := req.Execute()
//	if err != nil {
//		d.showError(err, "Info")
//		return
//	}
//	table := widgets.NewTable()
//	table.Rows = [][]string{
//		[]string{"Account", r.Data.AccountId},
//		[]string{"Symbol", r.Data.Symbol},
//		[]string{"Headquarters", r.Data.Headquarters},
//		[]string{"Credit", fmt.Sprintf("%d", r.Data.Credits)},
//	}
//	table.RowSeparator = true
//	table.FillRow = true
//	table.TextStyle = ui.NewStyle(ui.ColorWhite)
//
//	x1, y1, x2, y2 := getMainPanelRect()
//	table.SetRect(x1, y1, x2, y2)
//
//	ui.Render(table)
//}
//
//func (d *Dashboard) showContracts() {
//	d.listMode = ListContracts
//
//	ctx := context.Background()
//	req := d.client.ContractsApi.GetContracts(ctx)
//	r, _, err := req.Execute()
//	if err != nil {
//		d.showError(err, "Info")
//		return
//	}
//
//	d.list.Rows = []string{}
//	for _, contract := range r.Data {
//		d.list.Rows = append(d.list.Rows, fmt.Sprintf("[%s] %s - %s", contract.Id, contract.Type, contract.Expiration))
//	}
//
//	x1, y1, x2, y2 := getMainPanelRect()
//	d.list.SetRect(x1, y1, x2, y2)
//
//	ui.Render(d.list)
//}
