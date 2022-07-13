// Copyright 2022 Cisco Systems, Inc. and its affiliates
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
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cisco-open/flame/cmd/flamectl/resources/code"
)

var createDesignCodeCmd = &cobra.Command{
	Use:   "code <zipped code file>",
	Short: "Create a new ML code for a design",
	Long:  "Command to create a new ML code for a design",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		checkInsecure(cmd)

		codePath := args[0]
		flags := cmd.Flags()

		designId, err := flags.GetString("design")
		if err != nil {
			return err
		}

		params := code.Params{}
		params.Endpoint = config.ApiServer.Endpoint
		params.User = config.User
		params.DesignId = designId
		params.CodePath = codePath
		params.CodeVer = ""

		return code.Create(params)
	},
}

func init() {
	createDesignCodeCmd.PersistentFlags().StringP("design", "d", "", "Design ID")
	createDesignCodeCmd.MarkPersistentFlagRequired("design")
	createCmd.AddCommand(createDesignCodeCmd)
}
