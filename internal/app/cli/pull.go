package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/web-of-things-open-source/tm-catalog-cli/internal/commands"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/model"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/remotes"
	"github.com/web-of-things-open-source/tm-catalog-cli/internal/utils"
)

const (
	PullOK = PullResultType(iota)
	PullErr
)

type PullResultType int

func (t PullResultType) String() string {
	switch t {
	case PullOK:
		return "OK"
	case PullErr:
		return "error"
	default:
		return "unknown"
	}
}

type PullResult struct {
	typ  PullResultType
	tmid string
	text string
}

func (r PullResult) String() string {
	return fmt.Sprintf("%v\t %s %s", r.typ, r.tmid, r.text)
}

type PullExecutor struct {
	rm remotes.RemoteManager
}

func NewPullExecutor(rm remotes.RemoteManager) *PullExecutor {
	return &PullExecutor{
		rm: rm,
	}
}

func (e *PullExecutor) Pull(remote remotes.RepoSpec, search *model.SearchParams, outputPath string) error {
	if len(outputPath) == 0 {
		Stderrf("requires output target folder --output")
		return errors.New("--output not provided")
	}

	f, err := os.Stat(outputPath)
	if f != nil && !f.IsDir() {
		Stderrf("output target folder --output is not a folder")
		return errors.New("output target folder --output is not a folder")
	}

	searchResult, err := commands.NewListCommand(e.rm).List(remote, search)
	if err != nil {
		Stderrf("Error listing: %v", err)
		return err
	}

	fmt.Printf("Pulling %d ThingModels ...\n", len(searchResult.Entries))

	fc := commands.NewFetchCommand(e.rm)
	var totalRes []PullResult
	for _, entry := range searchResult.Entries {
		for _, version := range entry.Versions {
			res, pErr := e.pullThingModel(fc, outputPath, version)
			totalRes = append(totalRes, res)
			if pErr != nil {
				err = pErr
			}
		}
	}

	for _, res := range totalRes {
		fmt.Println(res)
	}

	return err
}

func (e *PullExecutor) pullThingModel(fc *commands.FetchCommand, outputPath string, version model.FoundVersion) (PullResult, error) {
	spec := remotes.NewSpecFromFoundSource(version.FoundIn)
	id, thing, err := fc.FetchByTMID(spec, version.TMID)
	if err != nil {
		Stderrf("Error fetch %s: %v", version.TMID, err)
		return PullResult{PullErr, version.TMID, fmt.Sprintf("(cannot fetch from remote %s)", spec.ToFoundSource())}, err
	}
	thing = utils.ConvertToNativeLineEndings(thing)

	finalOutput := filepath.Join(outputPath, id)

	err = os.MkdirAll(filepath.Dir(finalOutput), 0770)
	if err != nil {
		Stderrf("Could not write ThingModel to file %s: %v", finalOutput, err)
		return PullResult{PullErr, version.TMID, fmt.Sprintf("(cannot write to ouput directory %s)", outputPath)}, err
	}

	err = os.WriteFile(finalOutput, thing, 0660)
	if err != nil {
		Stderrf("Could not write ThingModel to file %s: %v", finalOutput, err)
		return PullResult{PullErr, version.TMID, fmt.Sprintf("(cannot write to ouput directory %s)", outputPath)}, err
	}

	return PullResult{PullOK, version.TMID, ""}, err
}
