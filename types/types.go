package types

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MarkedTorrent struct {
	Row         int
	TorrentCell *tview.TableCell
	LinkCell    *tview.TableCell
	Color       tcell.Color
}
