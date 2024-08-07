package cmd

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/gabe565/ransom/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	type args struct {
		args  []string
		stdin io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			"word",
			args{args: []string{"hello"}},
			":h::e::l::l::o:",
			require.NoError,
		},
		{
			"space",
			args{args: []string{"hello world"}},
			":h::e::l::l::o::blank::w::o::r::l::d:",
			require.NoError,
		},
		{
			"soon son on",
			args{args: []string{"soon son on spoon"}},
			":soon::blank::s::o::n::blank::on::blank::s::p::o::o::n:",
			require.NoError,
		},
		{
			"stop top",
			args{args: []string{"stop top"}},
			":s::t::o::p::blank::top:",
			require.NoError,
		},
		{
			"back",
			args{args: []string{"baby got back"}},
			":b::a::b::y::blank::g::o::t::blank::back:",
			require.NoError,
		},
		{
			"numbers",
			args{args: []string{"0123456789"}},
			":zero::one::two::three::four::five::six::seven::eight::nine:",
			require.NoError,
		},
		{
			"punctuation",
			args{args: []string{"!! !? ?! !!!."}},
			":bangbang::blank::interrobang::blank::question::exclamation::blank::bangbang::exclamation::hole:",
			require.NoError,
		},
		{
			"symbols",
			args{args: []string{"#*+-$=^"}},
			":hash::keycap_star::heavy_plus_sign::heavy_minus_sign::heavy_dollar_sign::heavy_equals_sign::this:",
			require.NoError,
		},
		{
			"multiple args",
			args{args: []string{"these", "are", "words"}},
			":t::h::e::s::e::blank::a::r::e::blank::w::o::r::d::s:",
			require.NoError,
		},
		{
			"unknown",
			args{args: []string{"/@()"}},
			"/@()",
			require.NoError,
		},
		{
			"stdin",
			args{stdin: strings.NewReader("hello world")},
			":h::e::l::l::o::blank::w::o::r::l::d:",
			require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			cmd := &cobra.Command{}
			conf := config.New()
			conf.NoCopy = true
			if tt.args.stdin != nil {
				cmd.SetIn(tt.args.stdin)
			}
			cmd.SetContext(config.NewContext(context.Background(), conf))
			cmd.SetOut(&buf)
			tt.wantErr(t, run(cmd, tt.args.args))
			got := strings.TrimSuffix(buf.String(), "\n")
			assert.Equal(t, tt.want, got)
		})
	}
}
