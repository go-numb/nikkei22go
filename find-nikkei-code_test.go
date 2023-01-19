package nikke22go

import (
	"os"
	"testing"

	"github.com/gocarina/gocsv"
)

// 日経225の証券コードをカテゴリ別に取得
func TestFindNikkei225Code(t *testing.T) {
	client := New()

	// まとめサイト速報＋の場合、depth:3で1,300リンクほど
	var depth = 1
	client.SetTarget("https://indexes.nikkei.co.jp/nkave/index/component?idx=nk225", depth)

	results := client.FindNikkei225Code()

	f, _ := os.Create("./find-nikkei-code.csv")
	defer f.Close()

	if err := gocsv.Marshal(&results, f); err != nil {
		t.Fatal(err)
	}

}
