package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsDesignPossible(t *testing.T) {
	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	testCases := []struct {
		input    string
		expected bool
	}{
		{"brwrr", true},
		{"bggr", true},
		{"gbbr", true},
		{"rrbgbr", true},
		{"ubwu", false},
		{"bwurrg", true},
		{"brgr", true},
		{"bbrgwb", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsDesignPossible(tc.input, patterns)
			if result != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, result)
			}
		})
	}

}

func TestIsDesignPossible2(t *testing.T) {
	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	patternDict := patternToDict(patterns)
	testCases := []struct {
		input    string
		expected bool
	}{
		{"brwrr", true},
		{"bggr", true},
		{"gbbr", true},
		{"rrbgbr", true},
		{"ubwu", false},
		{"bwurrg", true},
		{"brgr", true},
		{"bbrgwb", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsDesignPossible2(tc.input, patternDict)
			if result != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, result)
			}
		})
	}
}

func TestIsDesignPossible2Long(t *testing.T) {
	patterns := strings.Split("grbb, burb, wrwbrwg, uwwb, bwbbw, ubgrbu, gguu, uru, gwr, wrw, gubwb, g, gwbu, rbw, bbuu, rwgbr, urrr, rwww, wrrb, ug, rubbwuuu, gbrbr, brb, wrubb, gwrgbbgu, wggwbrww, rwwb, br, buuwr, rgrbb, wgubb, gbb, gbrb, rubw, ubr, guu, wrugbg, gubwru, bww, rbu, burr, ugu, bbggguw, bguw, ubu, uuuw, uugww, urugr, uwgur, gugrgggw, b, gugu, wgb, rwgwuu, guw, brg, ubur, wbwwb, ggbg, wr, urw, wur, brwgrw, uuur, rwwg, rrbrbggw, gwrurwg, rurru, r, bug, ruwrww, brgwu, bwgurbu, rwuu, guub, rrbwb, rbgwbr, grbrugbu, bwugwugb, guwww, ugwggg, buggrug, urb, bbbwuur, ururgur, wu, bbbru, ruuw, grr, buw, bub, gbwwug, rwrbw, wrbg, wubr, u, wgrb, wbuub, ugrugwg, wggbrr, gr, uur, ugb, wbg, uub, bbbrww, uwbr, buwguww, wbwgw, ubuuwr, wuuwgr, wggw, gwurgub, rwwu, bruwg, bgbrbwb, rbgg, rbg, bubggwu, gubgrur, gwb, rgu, wrr, uug, bgw, ruuugu, uwbg, rgurgwwb, wrwgb, brbbr, wgrgu, brgbwb, gub, uugw, wrrbb, wwruuru, grrw, urgg, gbwg, ubbur, bggwwgw, ubggw, uwgw, ruu, wwuu, bwuu, uggw, gwu, urg, rwbb, uwwub, ruug, bwbbu, wbb, bwwwu, rbwrruu, gg, ugrbbwgb, wbbwwr, bbwrrrb, rgwwu, bbr, wbuwb, grub, ubwwgu, ubb, brrbuu, grw, bwugg, wurbu, uwrb, rwgg, rurbuu, gbrug, brwgr, gru, bgrub, uuwr, gbwww, bbrw, rbr, bwrg, ruur, gbbgr, brgwug, ugg, bwgubu, rrg, ub, gbuwu, ubuwuw, ggw, wrgww, uuwb, bwr, bb, ruwu, bwugrw, buu, wug, gggwbwu, brurg, rrur, gb, wg, ggruwbu, gbg, bruur, gug, rgr, gwuuu, uwwurr, ubw, uugbuw, ruugwu, bu, ubrbw, ggu, rug, wgurwurr, ugrrbr, brgrb, rrbrgwww, rrgg, grb, brbw, bruwgw, uwbw, rgg, wrg, rrruwr, gguw, bugbbrw, bbg, rwg, wbuwg, uuu, wruwur, ugwb, bur, urgb, bg, uwubrwu, rwbwwg, rub, wwwwbw, rbbugbr, wgg, bwbb, brw, ggugb, guggw, bgg, wwr, gwrbb, rbb, bgb, brrubwb, wwu, bggrrwu, bugb, gwugw, grg, wrb, wuub, bru, gggrr, wuw, uwg, gbgurg, wguwug, gbwgw, bgrb, wguuw, rwr, ubuww, wruuw, wurur, wuu, bbw, brbgubb, bubg, wgw, wwuwrw, bbwu, bbrgubw, uwu, bw, wbr, bbu, bgr, ubg, rwrgg, uggu, wugggb, wwbgguur, wrwuwb, rgrr, wgbu, ggb, ubuw, rgugg, gwuru, guuubbw, bgbrrw, rrgwg, gbw, wbu, wuwg, gugb, uww, ruugb, rubu, ugwbw, bbrgrg, grbrr, guwgruwr, wwbr, wgruw, grwg, wgr, ururwr, gbbg, gbuub, uwr, grwgrwr, ruw, wru, ggr, guwuu, uwuuwu, gbu, rwbbgg, ruruww, wwwur, rgrug, rbbwg, rgwugwg, wbw, bbru, ggg, bbuugrrg, uugb, rw, gburwww, rububugg, gguubbg, rrrurwwg, buruwrb, bbwrggwb, gbrbwr, uruwwu, wwwurbu, rrwb, bbb, bwg, wurggubb, wgu, gwwg, wururruu, bbbu, ur, rwb, uwww, wgwr, rr, rru, uuw, uu, bwgw, brr, wub, bwgbrrw, gwrgrw, uw, ggguwbg, bubbur, ubbggu, ugr, rrrw, ru, rb, gwwbgw, rrr, rrw, gruw, rwrrbb, rugr, uubw, ruwb, uubbbu, gur, wb, rrrrwww, bbub, bgurwg, ubbugw, gwbw, gggbru, ugru, brru, ggwg, wwg, ugubr, urggbu, ggwuww, gbr, guwubu, rwu, gwg, rrbur, gwrg, wwubr, uwb, gbrr, rurburb, brrg, ggur, rurbw, uwbu, grbbu, wgur, urr, rgggg, rrb, ww, rww, bwwwr, wubwubgw, ggubw, wwb, uurb, wbrggwu, ugw, rggwwr, urwu, gwwggg, uwuubwr, gwuubww, uwbwu, www, gwgwr, bwu, gu, rur, wwwg, uwruub", ", ")
	patternDict := patternToDict(patterns)
	design := "rgruurwubbgggwwuwwgurrwuugggbrbuwgwrubrgw"

	DebugLog = true
	result := IsDesignPossible2(design, patternDict)
	fmt.Println("result: ", result)
}

func TestCountPossibilities(t *testing.T) {
	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	patternDict := patternToDict(patterns)

	testCases := []struct {
		input    string
		expected int
	}{
		{"brwrr", 2},
		{"bggr", 1},
		{"gbbr", 4},
		{"rrbgbr", 6},
		{"ubwu", 0},
		{"bwurrg", 1},
		{"brgr", 2},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := CountDesignPosibilities(tc.input, patternDict)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
