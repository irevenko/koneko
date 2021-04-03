package helpers

func ConvertCategory(idx int) string {
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
