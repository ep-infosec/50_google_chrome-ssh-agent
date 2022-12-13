//go:build js

// Copyright 2017 Google LLC
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

package message

import (
	"syscall/js"

	"github.com/google/chrome-ssh-agent/go/jsutil"
)

var (
	chromeObj = js.Global().Get("chrome")
	runtime   = func() js.Value {
		if chromeObj.IsUndefined() {
			return js.Undefined()
		}
		return chromeObj.Get("runtime")
	}()
)

// Sender specifies the interface for a type that sends messages.
type Sender interface {
	// Send sends a message. Response (or an error) are returned. See:
	//   https://developer.chrome.com/docs/extensions/reference/runtime/#method-sendMessage
	Send(ctx jsutil.AsyncContext, msg js.Value) (js.Value, error)
}

// ExtSender sends messages within our own extension.
//
// ExtSender implements the Sender interface.
type ExtSender struct{}

// NewLocalSender returns a ExtSender for sending messages within our own
// extension.
func NewLocalSender() *ExtSender {
	return &ExtSender{}
}

// Send implements Sender.Send().
func (e *ExtSender) Send(ctx jsutil.AsyncContext, msg js.Value) (js.Value, error) {
	return jsutil.AsPromise(runtime.Call("sendMessage", msg)).Await(ctx)
}
