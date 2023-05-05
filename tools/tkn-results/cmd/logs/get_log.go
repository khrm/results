package logs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	pb "github.com/tektoncd/results/proto/v1alpha2/results_go_proto"
	"github.com/tektoncd/results/tools/tkn-results/internal/flags"
	"github.com/tektoncd/results/tools/tkn-results/internal/format"
)

func GetRecordCommand(params *flags.Params) *cobra.Command {
	opts := &flags.GetOptions{}

	cmd := &cobra.Command{
		Use: `get [flags] <record>

  <record parent>: Record parent name to query. This is typically "<namespace>/results/<result name>", but may vary depending on the API Server. "-" may be used as <result name> to query all Results for a given parent.`,
		Short: "Get Record",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := params.LogsClient.GetLog(cmd.Context(), &pb.GetLogRequest{
				Name: args[0],
			})
			if err != nil {
				fmt.Printf("GetLog: %v\n", err)
				return err
			}
			data, err := resp.Recv()
			if err != nil {
				fmt.Printf("Get Log Client Resp: %v\n", err)
				return err
			}
			return format.PrintProto(os.Stdout, data, opts.Format)
		},
		Args: cobra.ExactArgs(1),
	}

	flags.AddGetFlags(opts, cmd)

	return cmd
}
