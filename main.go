package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/irevenko/go-nyaa/nyaa"
	h "github.com/irevenko/koneko/helpers"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var nyaaDownload = "https://nyaa.si/download/"

var query = ""
var category = 0
var filter = 0
var sortBy = 0
var markedTorrents []MarkedTorrent

var nyaaCategories = []string{
	"All",
	"Anime",
	"- Anime Music Video",
	"- English-translated",
	"- Non-English-translated",
	"- Raw",
	"Audio",
	"- Lossless",
	"- Lossy",
	"Literature",
	"- English-translated",
	"- Non-English-translated",
	"- Raw",
	"Live Action",
	"- English-translated",
	"- Idol/Promotional Video",
	"- Non-English-translated",
	"- Raw",
	"Pictures",
	"- Graphics",
	"- Photos",
	"Software",
	"- Applications",
	"- Games",
}

var filters = []string{
	"No filter",
	"No remakes",
	"Trusted only",
}

var sortOptions = []string{
	"Date",
	"Downloads",
	"Size",
	"Seeders",
	"Leechers",
	"Comments",
}

func main() {
	table := tview.NewTable().
		SetSelectable(true, false)
	table.SetBorder(true).SetTitle("nyaa.si")

	form := tview.NewForm().
		AddInputField("Query", "", 24, nil, func(text string) {
			query = text
		}).
		AddDropDown("Category", nyaaCategories, 0, func(option string, optionIndex int) {
			category = optionIndex
		}).
		AddDropDown("Filter", filters, 0, func(option string, optionIndex int) {
			filter = optionIndex
		}).
		AddDropDown("Sort By", sortOptions, 0, func(option string, optionIndex int) {
			sortBy = optionIndex
		}).
		AddButton("Search", func() {
			UnmarkAll(table)
			c := h.ConvertCategory(category)
			s := h.ConvertSort(sortBy)
			f := h.ConvertFilter(filter)
			torrents := fetchTorrents("nyaa", query, c, s, f)
			setTableData(table, torrents[:len(torrents)-1]) // remove last \n
			app.SetFocus(table)
			table.ScrollToBeginning()
		})

	form.SetBorder(true)
	form.SetHorizontal(true)

	form.SetCancelFunc(func() {
		app.SetFocus(table)
	})

	table.SetDoneFunc(func(key tcell.Key) {
		app.SetFocus(form)
	})

	table.SetSelectedFunc(func(row int, column int) {
		torrent := table.GetCell(row, 5)
		link := table.GetCell(row, 6)
		curColor := torrent.Color

		MarkTorrent(torrent, link, curColor, row)
		UnmarkTorrent(torrent, link, curColor)
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlD {
			_, err := downloadTorrents(markedTorrents)
			if err != nil {
				log.Fatal(err)
			}
			UnmarkAll(table)
		}
		if event.Key() == tcell.KeyCtrlO {
			err := openTorrents()
			if err != nil {
				log.Fatal(err)
			}
			UnmarkAll(table)
		}
		return event
	})

	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).SetDirection(tview.FlexRow).
		AddItem(table, 0, 5, true).SetDirection(tview.FlexRow)

	app.SetRoot(flex, true).EnableMouse(true)
	app.Run()
}

func setTableData(table *tview.Table, data string) {
	for row, line := range strings.Split(data, "\n") {
		for column, cell := range strings.Split(line, "{}") {
			textColor := tcell.ColorWhite
			bgColor := tcell.ColorBlack
			align := tview.AlignLeft

			if column == 0 {
				textColor = tcell.ColorYellow
			} else if column == 1 {
				textColor = tcell.ColorGreen
			} else if column == 2 {
				textColor = tcell.ColorRed
			} else if column == 3 {
				textColor = tcell.ColorPurple
			} else if column == 4 {
				textColor = tcell.ColorLightCyan
			}

			if strings.Contains(line, "(trusted torrent)") && column == 5 {
				textColor = tcell.ColorGreen
			}

			tableCell := tview.NewTableCell(cell).
				SetTextColor(textColor).
				SetBackgroundColor(bgColor).
				SetAlign(align).
				SetSelectable(true)
			table.SetCell(row, column, tableCell)
		}
	}
}

func fetchTorrents(p string, q string, c string, s string, f string) string {
	var torrents string

	opt := nyaa.SearchOptions{
		Provider: p,
		Query:    q,
		Category: c,
		SortBy:   s,
		Filter:   f,
	}

	res, err := nyaa.Search(opt)
	if err != nil {
		log.Fatal(err)
	}

	if len(res) == 0 {
		return "No results found"
	}

	initialLayout := "Mon, 02 Jan 2006 15:04:05 -0700"
	dateLayout := "2006-01-02"

	for _, v := range res {
		t, _ := time.Parse(initialLayout, v.Date)
		date := t.Format(dateLayout)

		nameSlice := strings.Split(v.Name, "")
		for i, v := range nameSlice { // chars like: "[" & "]" conflict with tcell or tview rendering process
			if v == "[" || v == "<" { // windows file name handling
				nameSlice[i] = "("
			} else if v == "]" || v == ">" {
				nameSlice[i] = ")"
			} else if v == "/" || v == "\\" || v == "|" {
				nameSlice[i] = "#"
			} else if v == ":" || v == "*" || v == "?" || v == `"` {
				nameSlice[i] = " "
			}
		}
		nameStr := strings.Join(nameSlice, "")
		name := nameStr

		isTrusted := ""
		if v.IsTrusted == "Yes" {
			isTrusted = v.IsTrusted
		}

		torrents += v.Downloads + "{}" + v.Seeders + "{}" + v.Leechers + "{}" + v.Size + "{}" + date + "{}" + name

		if isTrusted == "Yes" {
			torrents += " (trusted torrent)"
		}

		link := strings.Split(v.Link, "download/")
		torrentID := strings.Split(link[1], ".torrent")
		torrents += "{}" + torrentID[0]
		torrents += "\n"
	}

	return torrents
}

func downloadTorrents(torrents []MarkedTorrent) ([]string, error) {
	var names []string

	for _, v := range torrents {
		res, err := http.Get(nyaaDownload + v.LinkCell.Text + ".torrent")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		var fileName string
		if len(v.TorrentCell.Text) > 200 {
			fileName = v.TorrentCell.Text[:200]
		} else {
			fileName = v.TorrentCell.Text
		}

		out, err := os.Create(fileName + ".torrent")
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, out.Name())
		defer out.Close()

		_, err = io.Copy(out, res.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	return names, nil
}

func openTorrents() error {
	torrents, err := downloadTorrents(markedTorrents)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range torrents {
		if runtime.GOOS == "windows" {
			exec.Command("start", v).Start()
		} else if runtime.GOOS == "linux" {
			exec.Command("xdg-open", v).Start()
		} else if runtime.GOOS == "darwin" {
			exec.Command("open", v).Start()
		}
	}
	return nil
}

func MarkTorrent(torrent *tview.TableCell, link *tview.TableCell, curColor tcell.Color, row int) {
	if curColor == tcell.ColorGreen || curColor == tcell.ColorWhite {
		torrent.Color = tcell.ColorBlue

		if !sliceHas(markedTorrents, link.Text) {
			markedTorrents = append(markedTorrents, MarkedTorrent{
				TorrentCell: torrent,
				LinkCell:    link,
				Row:         row,
				Color:       curColor,
			})
		}
	}
}

func UnmarkTorrent(torrent *tview.TableCell, link *tview.TableCell, curColor tcell.Color) {
	if curColor == tcell.ColorBlue {
		if strings.Contains(torrent.Text, "trusted torrent") {
			torrent.Color = tcell.ColorGreen
		} else {
			torrent.Color = tcell.ColorWhite
		}

		markedTorrents = remove(markedTorrents, link.Text)
	}
}

func UnmarkAll(table *tview.Table) {
	rows := table.GetRowCount()
	for i := 0; i < rows; i++ {
		torrent := table.GetCell(i, 5)
		link := table.GetCell(i, 6)

		if torrent.Color == tcell.ColorBlue {
			if strings.Contains(torrent.Text, "trusted torrent") {
				torrent.Color = tcell.ColorGreen
			} else {
				torrent.Color = tcell.ColorWhite
			}

			markedTorrents = remove(markedTorrents, link.Text)
		}
	}
}

func sliceHas(s []MarkedTorrent, item string) bool {
	for _, v := range s {
		if v.LinkCell.Text == item {
			return true
		}
	}
	return false
}

func remove(s []MarkedTorrent, item string) []MarkedTorrent {
	for i, v := range s {
		if v.LinkCell.Text == item {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

type MarkedTorrent struct {
	Row         int
	TorrentCell *tview.TableCell
	LinkCell    *tview.TableCell
	Color       tcell.Color
}
