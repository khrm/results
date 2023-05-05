package records

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
			resp, err := params.ResultsClient.GetRecord(cmd.Context(), &pb.GetRecordRequest{
				Name: args[0],
			})
			if err != nil {
				fmt.Printf("GetRecord: %v\n", err)
				return err
			}
			return format.PrintProto(os.Stdout, resp, opts.Format)
		},
		Args: cobra.ExactArgs(1),
	}

	flags.AddGetFlags(opts, cmd)

	return cmd
}
