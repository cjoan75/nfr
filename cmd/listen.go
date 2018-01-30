package cmd

import (
	"github.com/alphasoc/nfr/client"
	"github.com/alphasoc/nfr/config"
	"github.com/alphasoc/nfr/executor"
	"github.com/spf13/cobra"
)

func newListenCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "listen",
		Short: "Start the DNS sniffer and score live events",
		Long: `Captures DNS traffic and provides analysis of them.
API key must be set before calling this mode.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, c, err := createConfigAndClient(true)
			if err != nil {
				return err
			}
			return listen(c, cfg)
		},
	}
	return cmd
}

func listen(c client.Client, cfg *config.Config) error {
	e, err := executor.New(c, cfg)
	if err != nil {
		return err
	}
	return e.Start()
}
