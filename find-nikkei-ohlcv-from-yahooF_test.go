package nikke22go

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gocarina/gocsv"
)

func TestGetOHLCvFromYahoo(t *testing.T) {
	var str = make([]NIkkei225Code, 0)
	f, _ := os.Open("find-nikkei-code.csv")
	if err := gocsv.Unmarshal(f, &str); err != nil {
		t.Fatal(err)
	}
	f.Close()

	for i, v := range str {
		code := v.Code
		fmt.Println("start: ", i, code, v.Name)
		results, metas, err := GetOHLCvFromYahoo(code)
		if err != nil {
			t.Fatal(err)
		}

		f, _ := os.Create(fmt.Sprintf("./yahooF/%s.csv", code))
		gocsv.Marshal(&results, f)
		f.Close()

		f, _ = os.Create(fmt.Sprintf("./yahooF/%s-meta.csv", code))
		gocsv.Marshal(&metas, f)
		f.Close()

		time.Sleep(2 * time.Second)
	}

}
