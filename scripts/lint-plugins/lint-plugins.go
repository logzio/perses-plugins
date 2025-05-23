// Copyright 2025 The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/perses/common/async"
	"github.com/perses/plugins/scripts/command"
	"github.com/perses/plugins/scripts/npm"
	"github.com/sirupsen/logrus"
)

func main() {
	workspaces, err := npm.GetWorkspaces()
	if err != nil {
		logrus.WithError(err).Fatal("unable to get the list of the workspaces")
	}

	pluginsToLint := make([]async.Future[string], 0, len(workspaces))

	for _, workspace := range workspaces {
		logrus.Infof("Linting plugin %s", workspace)
		pluginsToLint = append(pluginsToLint, async.Async(func() (string, error) {
			return workspace, command.Run("percli", "plugin", "lint", fmt.Sprintf("--plugin.path=%s", workspace))
		}))
	}
	isErr := false
	for _, pluginToLint := range pluginsToLint {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		workspace, buildErr := pluginToLint.AwaitWithContext(ctx)
		if buildErr != nil {
			isErr = true
			logrus.WithError(buildErr).Errorf("failed to lint plugin %s", workspace)
		}
		cancel()
	}
	if isErr {
		logrus.Fatal("some plugins have not been linted successfully")
	}
}
