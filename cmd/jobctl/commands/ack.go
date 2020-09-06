package commands

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/jeffrom/job-manager/jobclient"
	jobv1 "github.com/jeffrom/job-manager/pkg/resource/job/v1"
)

type ackCmd struct {
	*cobra.Command
}

func newAckCmd(cfg *jobclient.Config) *ackCmd {
	c := &ackCmd{
		Command: &cobra.Command{
			Use: "ack",
		},
	}

	return c
}

func (c *ackCmd) Cmd() *cobra.Command { return c.Command }
func (c *ackCmd) Execute(ctx context.Context, cfg *jobclient.Config, cmd *cobra.Command, args []string) error {
	client := clientFromContext(ctx)
	status := jobv1.StatusComplete
	return client.AckJob(ctx, args[0], status)
}
