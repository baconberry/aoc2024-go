package aoc

import (
	"strings"
	"testing"
)

func TestTwentythree(t *testing.T) {
	text := "kh-tc\nqp-kh\nde-cg\nka-co\nyn-aq\nqp-ub\ncg-tb\nvc-aq\ntb-ka\nwh-tc\nyn-cg\nkh-ub\nta-co\nde-co\ntc-td\ntb-wq\nwh-td\nta-ka\ntd-qp\naq-cg\nwq-ub\nub-vc\nde-ta\nwq-aq\nwq-vc\nwh-yn\nka-de\nkh-ta\nco-tc\nwh-qp\ntb-vc\ntd-yn\n"
	result := Twentythree(strings.Split(text, "\n"), 1)

	if result != 7 {
		t.Fail()
	}
	result = Twentythree(strings.Split(text, "\n"), 2)

	if result != 4 {
		t.Fail()
	}
}
