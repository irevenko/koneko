package tui

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
	t "github.com/irevenko/koneko/types"
	"github.com/rivo/tview"
)

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
			if v == "[" || v == "<" { // file name handling
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

		category := h.ConvertTableCategory(v.Category)

		torrents += v.Downloads + "{}" + v.Seeders + "{}" + v.Leechers + "{}" + v.Size + "{}" + date + "{}" + category + "{}" + name

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

func downloadTorrents(torrents []t.MarkedTorrent) ([]string, error) {
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

		if !h.SliceHas(markedTorrents, link.Text) {
			markedTorrents = append(markedTorrents, t.MarkedTorrent{
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

		markedTorrents = h.Remove(markedTorrents, link.Text)
	}
}

func UnmarkAll(table *tview.Table) {
	rows := table.GetRowCount()
	for i := 0; i < rows; i++ {
		torrent := table.GetCell(i, 6)
		link := table.GetCell(i, 7)

		if torrent.Color == tcell.ColorBlue {
			if strings.Contains(torrent.Text, "trusted torrent") {
				torrent.Color = tcell.ColorGreen
			} else {
				torrent.Color = tcell.ColorWhite
			}

			markedTorrents = h.Remove(markedTorrents, link.Text)
		}
	}
}
