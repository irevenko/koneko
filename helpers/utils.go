package helpers

import t "github.com/irevenko/koneko/types"

func SliceHas(s []t.MarkedTorrent, item string) bool {
	for _, v := range s {
		if v.LinkCell.Text == item {
			return true
		}
	}
	return false
}

func Remove(s []t.MarkedTorrent, item string) []t.MarkedTorrent {
	for i, v := range s {
		if v.LinkCell.Text == item {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
