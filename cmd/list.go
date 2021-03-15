package cmd

import (
	"github.com/spf13/afero"
	"github.com/yiranzai/github-starred/cmd/option"
	"github.com/yiranzai/github-starred/usecase"

	"github.com/spf13/cobra"
)

func newListCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Print Github Starred List",
		Long:    ``,
		Example: "github-starred list",
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := option.NewListCmdConfigFromViper()
			if err != nil {
				return err
			}
			return usecase.GetList(conf.Username, conf.Output, conf.Write, conf.All)
		},
	}

	if err := registerListCommandFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func registerListCommandFlags(cmd *cobra.Command) error {
	flags := []option.Flag{
		&option.StringFlag{
			BaseFlag: &option.BaseFlag{
				Name:       "username",
				Shorthand:  "u",
				Usage:      "your github username",
				IsRequired: true,
			},
		},
		&option.StringFlag{
			BaseFlag: &option.BaseFlag{
				Name:      "output",
				Shorthand: "o",
				Usage:     "output dir",
			},
			Value:     "docs/",
			IsDirName: true,
		},
		&option.BoolFlag{
			BaseFlag: &option.BaseFlag{
				Name:      "write",
				Shorthand: "w",
				Usage:     "write to output_dir/username.md",
			},
			Value: false,
		},
		&option.BoolFlag{
			BaseFlag: &option.BaseFlag{
				Name:      "all",
				Shorthand: "a",
				Usage:     "all body, no filters, format json",
			},
			Value: false,
		},
	}
	return option.RegisterFlags(cmd, flags)
}

func init() {
	cmdGenerators = append(cmdGenerators, newListCmd)
}
