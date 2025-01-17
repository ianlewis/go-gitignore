// Copyright 2016 Denormal Limited
// Copyright 2023 Google LLC
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

//go:build !windows

package gitignore_test

import (
	"testing"

	"github.com/ianlewis/go-gitignore"
)

// TODO(#19): Re-enable TestInvalidRepository on Windows.
func TestInvalidRepository(t *testing.T) {
	_test := &repositorytest{}
	_test.instance = func(path string) (gitignore.GitIgnore, error) {
		return gitignore.NewRepository(path)
	}

	// perform the invalid repository tests
	invalid(t, _test)
} // TestInvalidRepository()
