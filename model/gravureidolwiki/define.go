package gravureidolwiki

const (
	HomePageUrl = "https://idolwiki.web.fc2.com/"
)

type GuravureIdolWikiInfo struct {
	ThumbnailUrl string
}

type ArticleInfo = GuravureIdolWikiInfo

// bits演算によりフラグ分け。
type AlphabetFlg = uint

// アルファベット
const (
	// あ行
	Theあcolumn AlphabetFlg = 1 << iota
	// か行
	Theかcolumn
	// さ行
	Theさcolumn
	// た行
	Theたcolumn
	// な
	Theなcolumn
	// は
	Theはcolumn
	// ま
	Theまcolumn
	// や
	Theやcolumn
	// ら
	Theらcolumn
	// わ
	Theわcolumn

	// アルファベット 含むすべて
	ForAllColumnFlg = Theあcolumn |
		Theかcolumn |
		Theさcolumn |
		Theたcolumn |
		Theなcolumn |
		Theはcolumn |
		Theまcolumn |
		Theやcolumn |
		Theらcolumn |
		Theわcolumn
)
