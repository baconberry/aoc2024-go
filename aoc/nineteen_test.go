package aoc

import (
	"strings"
	"testing"
)

func TestNineteen(t *testing.T) {
	text := "r, wr, b, g, bwu, rb, gb, br\n\nbrwrr\nbggr\ngbbr\nrrbgbr\nubwu\nbwurrg\nbrgr\nbbrgwb\n"
	result := Nineteen(strings.Split(text, "\n"))

	if result != 6 {
		t.Fail()
	}
}
func TestNineteenHeavy(t *testing.T) {
	text := "wwubu, guw, uguggg, bgr, rgbrg, bbrb, rwgru, bwgr, uwwbuw, rwrbrgb, ggbrbuur, bb, ugu, grurrr, gbwbwb, rgbrb, bgrwrbu, bgb, wrwu, rwwr, wubr, bug, wu, brwu, w, grbw, uwbgg, gggwrggb, rurgr, uuw, bbw, wbg, bgw, ugw, brb, uww, buww, uuwgr, buu, rrbugw, ggggbgug, wwwrbbr, rgbbr, ugrr, wgwrg, bwuubub, uugg, bbrgg, bbugw, wuuwb, uwr, bugww, rwrrbbr, ugwrrrb, ubbr, rrwbgg, wb, wgr, gbrb, gwgg, rbg, bub, brg, bubr, ugggw, uuugbbb, grbug, rbww, rg, wgwwwwg, wgb, brugbbb, wgrww, bbrgrurg, ruu, brr, grgwg, wurbb, bww, ubwgr, bburwurb, urrrrbb, rgg, wwww, gbgw, rbrb, buw, wur, ruugugb, uwugg, rgrrugub, ggg, ggr, wrurr, bgg, uuru, brbg, rrbb, bbr, gbbu, grug, rwrrug, ggubub, ggrgg, uu, grb, rrrwwbg, ubbg, wgwwrgbr, rbug, rwrgwbrb, wubwgu, rrw, ruw, rbwubwgr, wwugg, bbbu, wgwr, rwgb, ruwg, rrrb, ubw, rwgw, wbr, uwgrb, rbwrrrru, guwbwg, rgbwr, wbu, ubgu, uruug, urbuuruw, rug, wg, gugbw, brw, wgrb, rbwu, wrw, bgggr, rubu, ur, uggrur, ggwbbuwu, gggr, gwu, bbwb, rrru, rubgb, wub, rgb, rbur, wuwbb, rbb, bwb, wwr, wubrwb, bwwwg, wwbwww, ugrrurw, grg, gwuw, uwwww, ubgb, gwr, gwuurgg, grwuu, gbru, rbbrrg, r, rgbrrg, ugrrubr, bw, wwbwbg, brbwbwr, rrrbg, rggb, bgwu, rrgrwwb, wrgrrb, grw, ubu, wgrugw, bbrugug, wubuwrb, gbrr, rrbg, bbg, bgrwg, urwu, uug, gruruwb, bur, rbwrguw, gbbuw, wubbr, gwurrw, rurr, gwrubb, wwwrw, wrwrbbr, ubbru, gbr, rwg, rrgwgr, ugrwwg, wwbgbr, wwubuwu, uwbb, gggbb, bru, uwb, bgru, wgu, gbrgr, urrwb, ub, u, bgrug, rrwr, bguug, gwbb, uuu, brubwg, bwrwbu, ruggur, gggww, rgu, ug, gur, wgurgwb, uuwwb, wbwgr, gbrg, rru, uuwurrr, rbgbw, ubgwgw, wgwbgu, brbwgg, uwwgw, rgww, buuu, uwu, wrgug, gug, rr, gggwb, bgu, grwg, wrwr, uub, gbuwwrwr, bgug, bwg, ubg, urrb, bbgggw, uuurb, ggw, wgrbwrg, bguw, brrrb, gwrbr, brbw, ggrrbu, uru, wrrbg, ubr, ggrr, ww, wwbgwwr, bgrr, wugbrwg, rrrrg, rww, wrbg, ubur, gwb, rgw, bubgbw, gbb, buuuuur, wwu, uubwww, buwg, ugbg, wgrwbub, wurwuu, bbu, wuru, ggu, wguwgr, gbw, gub, grr, rubgw, urrurb, wbubrr, grgwwb, gbwwru, urb, ruuw, urr, gu, bwrbwg, wrr, wru, wug, wbuwwwwg, guugu, brwurrb, brruuw, wuuuu, rbu, gwg, bgwg, rgr, bbrw, bu, gbbg, bbub, bwr, buur, wrgg, wwuurbug, bgbu, ruburubr, bwurbr, rbr, wgw, gwwug, rruu, wubgg, bwwg, wgwbwu, gwrr, wbbrrb, rwbgbbbr, gr, rwuw, rggbur, gbu, uggb, gwgr, guggbuu, gugw, brgww, wuw, rbubgbu, wgwbrb, g, urbb, gb, gbwu, uuug, ggrbbw, wbrrww, wbw, ugb, gbg, ugr, ggwww, rwr, rbrbw, wrg, wwbr, uggu, rrg, gwguuuuu, bbwuur, rgru, bburww, bbbg, wbb, uuugww, bwgg, bgggw, rugbr, wugrbwr, urw, bwu, uur, gggwrw, buwr, rrr, wwb, wgbw, urg, uwgbw, gru, bgbgurw, rwrrgbub, rwrw, buuug, wggbrgw, guu, bruuubw, gw, wuu, rrb, uuuwg, uwbugu, ugbbrr, wr, uw, ruwbw, wwg, ubru, guub, wrrwg, www, bbuuuw, guruub, wwbrwu, rwgbuu, rguruww, bgggrwr, bg, uwuwr, ugwub, ubgbu, ubbur, rub, gww, rw, urrbbbbu, urrbb, wuwu, ubuggww, bbbwg, ubugg, grgb, rwu, rugbug, ugg, wgg, gg, ru, wruw, wrrwuw, rubg, rbw, rwgwr, uuruu, ubb, ggb\n\nuwgwguuguruwggwbrbwurruwgwwwurgbgwggwwbbubgugguuurgugugwu\ngwwggrgbwbbguwgbguubbwrbgurgbburrgbwruwgurgwgguwbwgbgrr\nrgggbwgwgwuubwbwrrbrrrurwuwurgubuwrbwrbbbbr\nrguruururrwbbrgrbgrbwggrubbggbrbrugrgugggurubrbuuw\nrruggbrruruggubguwrubwbwuwgrwrguuurbrrbbbggbrwg\nwbwurbbrggubugbrgrrbuuwubwurrwururgbgbwurubbuuwbuw\nbrbwrrruwrrrubrwuugrbuuwuuwrwrbrrgububwurugbwwrb"
	result := Nineteen(strings.Split(text, "\n"))

	if result != 6 {
		t.Fail()
	}
}
