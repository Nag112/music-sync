package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
	"testing"
)

func TestAddLoggingFlags(t *testing.T) {
	names := []string{"comm-logging", "play-logging", "shed-logging", "ssh-logging", "time-logging", "logging"}
	f := AddLoggingFlags([]cli.Flag{})
	require.Equal(t, len(names), len(f), "AddLoggingFlags did not add the right number of flags")
	for i := range names {
		assert.Equal(t, names[i], f[i].GetName(), "AddLoggingFlags has the wrong flag at index %d", i)
	}
}

func TestFlagKey(t *testing.T) {
	cases := [][2]string{{"long-name, alias", "long-name"}, {"without-alias", "without-alias"}}
	for _, c := range cases {
		actual := FlagKey(cli.StringFlag{Name: c[0]})
		assert.Equal(t, c[1], actual, "wrong FlagKey for flag with name %s", c[0])
	}
}
