package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/app/cli"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/remotes"
)

var pFilterFlags = cli.FilterFlags{}

var pullCmd = &cobra.Command{
	Use:   "pull <NAME PATTERN>",
	Short: "Pull TMs from a catalog.",
	Long: `Pulls one or more TMs from a catalog by name pattern, filters or search. 
The pattern can be a full name or just a prefix the names shall start with. 
Name pattern, filters and search can be combined to narrow down the result.
Use --exact to force full-length matching of the name`,
	Args: cobra.MaximumNArgs(1),
	Run:  executePull,
}

func init() {
	RootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringP("remote", "r", "", "name of the remote to pull from")
	pullCmd.Flags().StringP("directory", "d", "", "TM repository directory to pull from")
	pullCmd.Flags().StringP("output", "o", "", "output directory where to save the pulled TMs")
	pullCmd.Flags().StringVar(&pFilterFlags.FilterAuthor, "filter.author", "", "filter TMs by one or more comma-separated authors")
	pullCmd.Flags().StringVar(&pFilterFlags.FilterManufacturer, "filter.manufacturer", "", "filter TMs by one or more comma-separated manufacturers")
	pullCmd.Flags().StringVar(&pFilterFlags.FilterMpn, "filter.mpn", "", "filter TMs by one or more comma-separated mpn (manufacturer part number)")
	pullCmd.Flags().StringVar(&pFilterFlags.FilterExternalID, "filter.externalID", "", "filter TMs by one or more comma-separated external ID")
	pullCmd.Flags().StringVarP(&pFilterFlags.Search, "search", "s", "", "search TMs by their content matching the search term")
	pullCmd.Flags().BoolP("exact", "e", false, "match the TM name exactly. overrides all other search filter flags")
	_ = pullCmd.MarkFlagRequired("output")
}

func executePull(cmd *cobra.Command, args []string) {
	remoteName := cmd.Flag("remote").Value.String()
	dirName := cmd.Flag("directory").Value.String()
	outputPath := cmd.Flag("output").Value.String()

	spec, err := remotes.NewSpec(remoteName, dirName)
	if errors.Is(err, remotes.ErrInvalidSpec) {
		cli.Stderrf("Invalid specification of target repository. --remote and --directory are mutually exclusive. Set at most one")
		os.Exit(1)
	}

	name := ""
	if len(args) > 0 {
		name = args[0]
	}
	exact, err := cmd.Flags().GetBool("exact")
	if err != nil {
		cli.Stderrf("invalid --exact flag")
		os.Exit(1)
	}
	search := cli.CreateSearchParamsFromCLI(pFilterFlags, name, exact)
	err = cli.NewPullExecutor(remotes.DefaultManager()).Pull(spec, search, outputPath)

	if err != nil {
		cli.Stderrf("pull failed")
		os.Exit(1)
	}
}
