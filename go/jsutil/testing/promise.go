//go:build js

// Copyright 2022 Google LLC
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

package jsutil

import (
	"syscall/js"

	"github.com/google/chrome-ssh-agent/go/jsutil"
)

// DoSync runs f, and then blocks until it is complete. Since this blocks,
// it is not suitable to run outside of tests.
func DoSync(f func(ctx jsutil.AsyncContext)) {
	done := make(chan struct{})

	jsutil.Async(func(ctx jsutil.AsyncContext) (js.Value, error) {
		f(ctx)
		return js.Undefined(), nil
	}).Then(
		func(val js.Value) { close(done) },
		func(err error) { close(done) },
	)

	<-done
}
