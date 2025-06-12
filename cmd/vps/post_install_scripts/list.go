package post_install_scripts

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get post-install script list",
	Long:  `This endpoint retrieves a list of post-install scripts associated with your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetPostInstallScriptListV1WithResponse(context.TODO(), listRequestParameters(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 1, "Page number")
}

func listRequestParameters(cmd *cobra.Command) *client.VPSGetPostInstallScriptListV1Params {
	pageId, _ := cmd.Flags().GetInt("page")

	return &client.VPSGetPostInstallScriptListV1Params{
		Page: utils.IntPtrOrNil(pageId),
	}
}
