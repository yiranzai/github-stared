package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yiranzai/github-starred/util"
)

func newSelfUpdateCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "selfupdate",
		Short: "Update github_starred",
		//Long: `Update github_starred`,
		Run: func(cmd *cobra.Command, args []string) {
			updated, err := util.Do()
			if err != nil {
				cmd.Println("Binary update failed:", err)
				return
			}
			if updated {
				cmd.Println("Current binary is the latest version", util.Version)
			} else {
				cmd.Println("Successfully updated to version", util.Version)
			}
		},
	}
	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newSelfUpdateCmd)
}
