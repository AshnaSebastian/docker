package node

import (
	"fmt"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	apiclient "github.com/docker/engine-api/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

// NewNodeCommand returns a cobra command for `node` subcommands
func NewNodeCommand(dockerCli *client.DockerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Manage Docker Swarm nodes",
		Args:  cli.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(dockerCli.Err(), "\n"+cmd.UsageString())
		},
	}
	cmd.AddCommand(
		newDemoteCommand(dockerCli),
		newInspectCommand(dockerCli),
		newListCommand(dockerCli),
		newPromoteCommand(dockerCli),
		newRemoveCommand(dockerCli),
		newPsCommand(dockerCli),
		newUpdateCommand(dockerCli),
	)
	return cmd
}

// Reference returns the reference of a node. The special value "self" for a node
// reference is mapped to the current node, hence the node ID is retrieved using
// the `/info` endpoint.
func Reference(ctx context.Context, client apiclient.APIClient, ref string) (string, error) {
	if ref == "self" {
		info, err := client.Info(ctx)
		if err != nil {
			return "", err
		}
		return info.Swarm.NodeID, nil
	}
	return ref, nil
}
