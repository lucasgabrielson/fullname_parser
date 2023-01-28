package fullname_parser

import (
	"testing"
)

func TestParseFullname(t *testing.T) {
	tests := map[string]struct {
		fullname       string
		wantParsedName ParsedName
	}{
		"basic": {
			fullname: "Lucas Gabrielson",
			wantParsedName: ParsedName{
				First: "Lucas",
				Last:  "Gabrielson",
			},
		},

		"conjuctions": {
			fullname: "Lucas Gabrielson y Suarez",
			wantParsedName: ParsedName{
				First: "Lucas",
				Last:  "Gabrielson y Suarez",
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			gotParsedName := ParseFullname(tt.fullname)
			if gotParsedName.First != tt.wantParsedName.First {
				t.Errorf("got %v, want %v", gotParsedName.First, tt.wantParsedName.First)
			}
			if gotParsedName.Last != tt.wantParsedName.Last {
				t.Errorf("got %v, want %v", gotParsedName.Last, tt.wantParsedName.Last)
			}
		})
	}
}
