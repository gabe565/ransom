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
			":h::e::l::l::o:\n",
			require.NoError,
		},
		{
			"space",
			args{args: []string{"hello world"}},
			":h::e::l::l::o::blank::w::o::r::l::d:\n",
			require.NoError,
		},
		{
			"soon son on",
			args{args: []string{"soon son on spoon"}},
			":soon::blank::s::o::n::blank::on::blank::s::p::o::o::n:\n",
			require.NoError,
		},
		{
			"stop top",
			args{args: []string{"stop top"}},
			":s::t::o::p::blank::top:\n",
			require.NoError,
		},
		{
			"back",
			args{args: []string{"baby got back"}},
			":b::a::b::y::blank::g::o::t::blank::back:\n",
			require.NoError,
		},
		{
			"numbers",
			args{args: []string{"0123456789"}},
			":zero::one::two::three::four::five::six::seven::eight::nine:\n",
			require.NoError,
		},
		{
			"punctuation",
			args{args: []string{"!! !? ?! !!!."}},
			":bangbang::blank::interrobang::blank::question::exclamation::blank::bangbang::exclamation::hole:\n",
			require.NoError,
		},
		{
			"symbols",
			args{args: []string{"#*+-$=^"}},
			":hash::keycap_star::heavy_plus_sign::heavy_minus_sign::heavy_dollar_sign::heavy_equals_sign::this:\n",
			require.NoError,
		},
		{
			"multiple args",
			args{args: []string{"these", "are", "words"}},
			":t::h::e::s::e::blank::a::r::e::blank::w::o::r::d::s:\n",
			require.NoError,
		},
		{
			"unknown",
			args{args: []string{"/@()"}},
			"/@()\n",
			require.NoError,
		},
		{
			"stdin",
			args{stdin: strings.NewReader("hello world")},
			":h::e::l::l::o::blank::w::o::r::l::d:\n",
			require.NoError,
		},
		{
			"empty string",
			args{args: []string{""}},
			"",
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
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
