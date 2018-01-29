package twitter

import (
	"fmt"
	"testing"
)

func TestScrapeNissy(t *testing.T) {
	actual, err := ScrapeNissy()
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("data : %+v", actual[0])
}

func TestScrapeAAA(t *testing.T) {
	actual, err := ScrapeAAA()
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("data : %+v", actual[0])
}
func TestScrapeAtae(t *testing.T) {
	actual, err := ScrapeAtae()
	if err != nil {
		t.Errorf("err : %v", err)
	}
	fmt.Printf("data : %+v", actual[0])
}
