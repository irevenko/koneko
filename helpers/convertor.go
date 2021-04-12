package helpers

func ConvertNyaaCategory(idx int) string {
	switch idx {
	case 0:
		return "all"
	case 1:
		return "anime"
	case 2:
		return "anime-amv"
	case 3:
		return "anime-eng"
	case 4:
		return "anime-non-eng"
	case 5:
		return "anime-raw"
	case 6:
		return "audio"
	case 7:
		return "audio-lossless"
	case 8:
		return "audio-lossy"
	case 9:
		return "literature"
	case 10:
		return "literature-eng"
	case 11:
		return "literature-non-eng"
	case 12:
		return "literature-raw"
	case 13:
		return "live-action"
	case 14:
		return "live-action-eng"
	case 15:
		return "live-action-idol-prom"
	case 16:
		return "live-action-non-eng"
	case 17:
		return "live-action-raw	"
	case 18:
		return "pictures"
	case 19:
		return "pictures-graphics"
	case 20:
		return "pictures-photos"
	case 21:
		return "software"
	case 22:
		return "software-apps"
	case 23:
		return "software-games"
	default:
		return ""
	}
}

func ConvertSukebeiCategory(idx int) string {
	switch idx {
	case 0:
		return "all"
	case 1:
		return "art"
	case 2:
		return "art-anime"
	case 3:
		return "art-doujinshi"
	case 4:
		return "art-games"
	case 5:
		return "art-manga"
	case 6:
		return "art-pictures"
	case 7:
		return "real-life"
	case 8:
		return "real-life-photos"
	case 9:
		return "real-life-videos"
	default:
		return ""
	}
}

func ConvertSort(idx int) string {
	switch idx {
	case 0:
		return "date"
	case 1:
		return "downloads"
	case 2:
		return "size"
	case 3:
		return "seeders"
	case 4:
		return "leechers"
	case 5:
		return "comments"
	default:
		return ""
	}
}

func ConvertFilter(idx int) string {
	switch idx {
	case 0:
		return "no-filter"
	case 1:
		return "no-remakes"
	case 2:
		return "trusted-only"
	default:
		return ""
	}
}

func ConvertTableNyaa(category string) string {
	switch category {
	case "Anime - Anime Music Video":
		return "Anime-AMV"
	case "Anime - English-translated":
		return "Anime-ENG"
	case "Anime - Non-English-translated":
		return "Anime-Non-ENG"
	case "Anime - Raw":
		return "Anime-Raw"
	case "Audio - Lossless":
		return "Audio-Lossless"
	case "Audio - Lossy":
		return "Audio-Lossy"
	case "Literature - English-translated":
		return "Literature-ENG"
	case "Literature - Non-English-translated":
		return "Literature-Non-ENG"
	case "Literature - Raw":
		return "Literature-Raw"
	case "Live Action - English-translated":
		return "Live-Act-ENG"
	case "Live Action - Idol/Promotional Video":
		return "Live-Act-Idol/Prom"
	case "Live Action - Non-English-translated":
		return "Live-Act-Non-ENG"
	case "Live Action - Raw":
		return "Live-Act-Raw"
	case "Pictures - Graphics":
		return "Pics-Graphics"
	case "Pictures - Photos":
		return "Pics-Photos"
	case "Software - Applications":
		return "Applications"
	case "Software - Games":
		return "Games"
	default:
		return ""
	}
}

func ConvertTableSukebei(category string) string {
	switch category {
	case "Art - Anime":
		return "Anime"
	case "Art - Doujinshi":
		return "Doujinshi"
	case "Art - Games":
		return "Games"
	case "Art - Manga":
		return "Manga"
	case "Art - Pictures":
		return "Art-Pictures"
	case "Real Life - Photobooks and Pictures":
		return "Real-Life-Pics"
	case "Real Life - Videos":
		return "Real-Life-Vids"
	default:
		return ""
	}
}
