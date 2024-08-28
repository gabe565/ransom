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
		config  *config.Config
		args    args
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			"word",
			nil,
			args{args: []string{"hello"}},
			":h::e::l::l::o:\n",
			require.NoError,
		},
		{
			"space",
			nil,
			args{args: []string{"hello world"}},
			":h::e::l::l::o::blank::w::o::r::l::d:\n",
			require.NoError,
		},
		{
			"soon son on",
			nil,
			args{args: []string{"soon son on spoon"}},
			":soon::blank::s::o::n::blank::on::blank::s::p::o::o::n:\n",
			require.NoError,
		},
		{
			"stop top",
			nil,
			args{args: []string{"stop top"}},
			":s::t::o::p::blank::top:\n",
			require.NoError,
		},
		{
			"back",
			nil,
			args{args: []string{"baby got back"}},
			":b::a::b::y::blank::g::o::t::blank::back:\n",
			require.NoError,
		},
		{
			"numbers",
			nil,
			args{args: []string{"0123456789"}},
			":zero::one::two::three::four::five::six::seven::eight::nine:\n",
			require.NoError,
		},
		{
			"punctuation",
			nil,
			args{args: []string{"!! !? ?! !!!."}},
			":bangbang::blank::interrobang::blank::question::exclamation::blank::bangbang::exclamation::hole:\n",
			require.NoError,
		},
		{
			"symbols",
			nil,
			args{args: []string{"#*+-$=^"}},
			":hash::keycap_star::heavy_plus_sign::heavy_minus_sign::heavy_dollar_sign::heavy_equals_sign::this:\n",
			require.NoError,
		},
		{
			"multiple args",
			nil,
			args{args: []string{"these", "are", "words"}},
			":t::h::e::s::e::blank::a::r::e::blank::w::o::r::d::s:\n",
			require.NoError,
		},
		{
			"unknown",
			nil,
			args{args: []string{"/@()"}},
			"/@()\n",
			require.NoError,
		},
		{
			"stdin",
			nil,
			args{stdin: strings.NewReader("hello world")},
			":h::e::l::l::o::blank::w::o::r::l::d:\n",
			require.NoError,
		},
		{
			"stdin multiline",
			nil,
			args{stdin: strings.NewReader("hello\nworld")},
			":h::e::l::l::o:\n:w::o::r::l::d:\n",
			require.NoError,
		},
		{
			"yellow pack",
			&config.Config{Prefix: "alphabet-white"},
			args{args: []string{"non-letters are left as-is."}},
			":alphabet-white-n::alphabet-white-o::alphabet-white-n::heavy_minus_sign::alphabet-white-l::alphabet-white-e::alphabet-white-t::alphabet-white-t::alphabet-white-e::alphabet-white-r::alphabet-white-s::blank::alphabet-white-a::alphabet-white-r::alphabet-white-e::blank::alphabet-white-l::alphabet-white-e::alphabet-white-f::alphabet-white-t::blank::alphabet-white-a::alphabet-white-s::heavy_minus_sign::alphabet-white-i::alphabet-white-s::hole:\n",
			require.NoError,
		},
		{
			"empty string",
			nil,
			args{args: []string{""}},
			"",
			require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			cmd := &cobra.Command{}
			if tt.config == nil {
				tt.config = config.New()
			}
			tt.config.NoCopy = true
			if tt.args.stdin != nil {
				cmd.SetIn(tt.args.stdin)
			}
			cmd.SetContext(config.NewContext(context.Background(), tt.config))
			cmd.SetOut(&buf)
			tt.wantErr(t, run(cmd, tt.args.args))
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
