package cmd

import (
	"context"
	"strings"
	"testing"

	"github.com/gabe565/ransom/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{"word", []string{"hello"}, ":h::e::l::l::o:", require.NoError},
		{"space", []string{"hello world"}, ":h::e::l::l::o::blank::w::o::r::l::d:", require.NoError},
		{"soon son on", []string{"soon son on spoon"}, ":soon::blank::s::o::n::blank::on::blank::s::p::o::o::n:", require.NoError},
		{"stop top", []string{"stop top"}, ":s::t::o::p::blank::top:", require.NoError},
		{"back", []string{"baby got back"}, ":b::a::b::y::blank::g::o::t::blank::back:", require.NoError},
		{"numbers", []string{"0123456789"}, ":zero::one::two::three::four::five::six::seven::eight::nine:", require.NoError},
		{"punctuation", []string{"!! !? ?! !!!."}, ":bangbang::blank::interrobang::blank::question::exclamation::blank::bangbang::exclamation::hole:", require.NoError},
		{"symbols", []string{"#*+-$=^"}, ":hash::keycap_star::heavy_plus_sign::heavy_minus_sign::heavy_dollar_sign::heavy_equals_sign::this:", require.NoError},
		{"multiple args", []string{"these", "are", "words"}, ":t::h::e::s::e::blank::a::r::e::blank::w::o::r::d::s:", require.NoError},
		{"unknown", []string{"/@()"}, "/@()", require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			cmd := &cobra.Command{}
			cmd.SetContext(config.NewContext(context.Background(), config.New()))
			cmd.SetOut(&buf)
			tt.wantErr(t, run(cmd, tt.args))
			got := strings.TrimSuffix(buf.String(), "\n")
			assert.Equal(t, tt.want, got)
		})
	}
}
