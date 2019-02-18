package stringsx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstring_NoLimits(t *testing.T) {
	testCases := map[string]struct {
		str            string
		expectedSubstr string
	}{
		"no limits, no unicode, should return whole string": {
			"How you doin'",
			"How you doin'",
		},
		"no limits, unicode, should return whole string": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
		},
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, tc.expectedSubstr, Substring(tc.str))
		})
	}
}

func TestSubstring_OnlyOffset(t *testing.T) {
	testCases := map[string]struct {
		str            string
		offset         int
		expectedSubstr string
	}{
		"no length, no unicode, offset 1 should remove 1st char": {
			"How you doin'",
			1,
			"ow you doin'",
		},
		"no length, unicode, offset 1 should remove 1st char": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			1,
			"èàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
		},
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, tc.expectedSubstr, Substring(tc.str, tc.offset))
		})
	}
}

func TestSubstring_OnlyBothLimits(t *testing.T) {
	testCases := map[string]struct {
		str            string
		offset         int
		length         int
		expectedSubstr string
	}{
		"both limits, no unicode, should return whole string": {
			"How you doin'",
			0,
			13,
			"How you doin'",
		},
		"both limits, unicode, should return whole string": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			0,
			21,
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
		},
		"both limits, no unicode, should remove first char": {
			"How you doin'",
			1,
			13,
			"ow you doin'",
		},
		"both limits, unicode, should remove first char": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			1,
			21,
			"èàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
		},
		"both limits, no unicode, should remove last char": {
			"How you doin'",
			0,
			12,
			"How you doin",
		},
		"both limits, unicode, should remove last char": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			0,
			20,
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º",
		},
		"both limits, no unicode, should remove first 3 and last 4 chars": {
			"How you doin'",
			3,
			9,
			" you d",
		},
		"both limits, unicode, should remove first 3 and last 4 chars": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			3,
			17,
			"ùêÊäÄöÖüÜçÇ¿ß¡",
		},
		"both limits, no unicode, length longer than string length": {
			"How you doin'",
			3,
			30,
			" you doin'",
		},
		"both limits, unicode, length longer than string length": {
			"ùèàùêÊäÄöÖüÜçÇ¿ß¡£§º©",
			3,
			30,
			"ùêÊäÄöÖüÜçÇ¿ß¡£§º©",
		},
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			assert.Equal(t, tc.expectedSubstr, Substring(tc.str, tc.offset, tc.length))
		})
	}
}
