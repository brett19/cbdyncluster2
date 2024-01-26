package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var toolsCaoInstallAdmissionCmd = &cobra.Command{
	Use:   "install-admission",
	Short: "Automatically installs the cao admission controller.",
	Run: func(cmd *cobra.Command, args []string) {
		helper := CmdHelper{}
		logger := helper.GetLogger()
		ctx := helper.GetContext()
		caoDeployer := helper.GetCaoDeployer(ctx)

		namespace, _ := cmd.Flags().GetString("namespace")

		err := caoDeployer.GetClient().InstallGlobalAdmissionController(ctx, namespace)
		if err != nil {
			logger.Fatal("failed to install admission controller", zap.Error(err))
		}
	},
}

func init() {
	toolsCaoCmd.AddCommand(toolsCaoInstallAdmissionCmd)

	toolsCaoInstallAdmissionCmd.Flags().String("namespace", "", "Which namespace to install in.")
}
