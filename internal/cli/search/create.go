// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20241023002/admin"
)

var errInvalidIndex = errors.New("invalid index")

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	IndexOpts
	clusterName string
	store       store.SearchIndexCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Index {{.Name}} created.\n"

func (opts *CreateOpts) Run() error {
	i, err := opts.CreateSearchIndex()
	if err != nil {
		return err
	}

	switch index := i.(type) {
	case *admin.SearchIndexCreateRequest:
		telemetry.AppendOption(telemetry.WithSearchIndexType(index.GetType()))
		r, err := opts.store.CreateSearchIndexes(opts.ConfigProjectID(), opts.clusterName, index)
		if err != nil {
			return err
		}

		return opts.Print(r)
	case *admin.ClusterSearchIndex:
		_, _ = log.Warningln("you're using an old search index definition")
		telemetry.AppendOption(telemetry.WithSearchIndexType(index.GetType()))
		r, err := opts.store.CreateSearchIndexesDeprecated(opts.ConfigProjectID(), opts.clusterName, index)
		if err != nil {
			return err
		}

		return opts.Print(r)
	default:
		return errInvalidIndex
	}
}

// CreateBuilder
// Create an online archive for a cluster.
//
// Usage:
//
//	atlas clusters search index create [indexName] [flags]
//
// Flags:
//
//	    --analyzer string         Analyzer to use when creating the index (default "lucene.standard")
//	    --clusterName string      Name of the cluster.
//	    --collection string       Collection name.
//	    --db string               Database name.
//	    --dynamic                 Indicates whether the index uses dynamic or static mappings.
//	    --field strings           Static field specifications.
//	-h, --help                    help for create
//	    --projectId string        Project ID to use. Overrides configuration file or environment variable settings.
//	    --searchAnalyzer string   Analyzer to use when searching the index. (default "lucene.standard")
//	-f, --file string             JSON file to use in order to create the index
//
// Global Flags:
//
//	-P, --profile string   Profile to use from your configuration file.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	opts.Fs = afero.NewOsFs()

	cmd := &cobra.Command{
		Use:   "create [indexName]",
		Short: "Create a search index for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"indexNameDesc": "Name of the index.",
			"output":        createTemplate,
		},
		Example: `  # Create a search index for the cluster named myCluster using a JSON index configuration file named search-config.json:
  atlas clusters search indexes create --clusterName myCluster --file search-config.json --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Filename == "" {
				_ = cmd.MarkFlagRequired(flag.Database)
				_ = cmd.MarkFlagRequired(flag.Collection)

				if err := require.ExactArgs(1)(cmd, args); err != nil {
					return err
				}
			} else if len(args) > 0 {
				return errors.New("when passing --file no indexName should be provided")
			}

			return opts.PreRunE(
				opts.validateOpts,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.Name = args[0]
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVarP(&opts.Filename, flag.File, flag.FileShort, "", usage.SearchFilename)

	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagRequired(flag.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	// deprecated flags
	cmd.Flags().StringVar(&opts.DBName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.Collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.Analyzer, flag.Analyzer, DefaultAnalyzer, usage.Analyzer)
	cmd.Flags().StringVar(&opts.SearchAnalyzer, flag.SearchAnalyzer, DefaultAnalyzer, usage.SearchAnalyzer)
	cmd.Flags().BoolVar(&opts.Dynamic, flag.Dynamic, false, usage.Dynamic)
	cmd.Flags().StringSliceVar(&opts.fields, flag.Field, nil, usage.SearchFields)

	_ = cmd.Flags().MarkDeprecated(flag.Database, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Collection, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Analyzer, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.SearchAnalyzer, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Dynamic, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Field, deprecatedFlagMessage)

	return cmd
}
