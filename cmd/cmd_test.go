package cmd

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"word", []string{"hello"}, ":h::e::l::l::o:"},
		{"space", []string{"hello world"}, ":h::e::l::l::o::blank::w::o::r::l::d:"},
		{"soon son on", []string{"soon son on spoon"}, ":soon::blank::s::o::n::blank::on::blank::s::p::o::o::n:"},
		{"stop top", []string{"stop top"}, ":s::t::o::p::blank::top:"},
		{"back", []string{"baby got back"}, ":b::a::b::y::blank::g::o::t::blank::back:"},
		{"numbers", []string{"0123456789"}, ":zero::one::two::three::four::five::six::seven::eight::nine:"},
		{"punctuation", []string{"!! !? ?! !!!."}, ":bangbang::blank::interrobang::blank::question::exclamation::blank::bangbang::exclamation::hole:"},
		{"symbols", []string{"#*+-$=^"}, ":hash::keycap_star::heavy_plus_sign::heavy_minus_sign::heavy_dollar_sign::heavy_equals_sign::this:"},
		{"multiple args", []string{"these", "are", "words"}, ":t::h::e::s::e::blank::a::r::e::blank::w::o::r::d::s:"},
		{"unknown", []string{"/@()"}, "/@()"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			cmd := &cobra.Command{}
			cmd.SetOut(&buf)
			run(cmd, tt.args)
			got := strings.TrimSuffix(buf.String(), "\n")
			assert.Equal(t, tt.want, got)
		})
	}
}
