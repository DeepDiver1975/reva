// Copyright 2018-2024 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package appctx

import "context"

// ResoucePathCtx is the key used in the opaque id for passing the resource path.
const ResoucePathCtx = "resource_path"

// ContextGetResourcePath returns the resource path if set in the given context.
func ContextGetResourcePath(ctx context.Context) (string, bool) {
	p, ok := ctx.Value(pathKey).(string)
	return p, ok
}

// ContextGetResourcePath stores the resource path in the context.
func ContextSetResourcePath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, pathKey, path)
}
