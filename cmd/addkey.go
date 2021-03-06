// Copyright © 2019 Karl Hepworth <Karl.Hepworth@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/fubarhouse/pygmy-go/service/library"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// addkeyCmd is the SSH key add command.
var addkeyCmd = &cobra.Command{
	Use:   "addkey",
	Example: "pygmy addkey ~/.ssh/id_rsa",
	Short: "Add/re-add an SSH key to the agent",
	Long: `Add or re-add an SSH key to Pygmy's SSH Agent by specifying the path to the private key.`,
	Run: func(cmd *cobra.Command, args []string) {

		Key, _ := cmd.Flags().GetString("key")

		library.SshKeyAdd(c, Key)

	},
}

func init() {

	homedir, _ := homedir.Dir()
	keypath := fmt.Sprintf("%v%v.ssh%vid_rsa", homedir, string(os.PathSeparator), string(os.PathSeparator))

	rootCmd.AddCommand(addkeyCmd)
	addkeyCmd.Flags().StringP("key", "", keypath, "Path of SSH key to add")

}
