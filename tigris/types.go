// Copyright 2022-2023 Tigris Data, Inc.
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

package tigris

import (
	api "github.com/tigrisdata/tigris-client-go/api/server/v1"
	"github.com/tigrisdata/tigris-client-go/code"
	"github.com/tigrisdata/tigris-client-go/driver"
)

// InsertResponse includes metadata regarding just inserted documents.
// Returned by Insert-documents collection API.
type InsertResponse struct {
	// JSON documents, which contain autogenerated fields
	Keys [][]byte
}

// InsertOrReplaceResponse includes metadata regarding just inserted
// or replaced documents. Returned by InsertOrReplace-documents collection API.
type InsertOrReplaceResponse struct {
	// JSON documents, which contain autogenerated fields
	Keys [][]byte
}

// DeleteResponse includes metadata about just deleted documents.
// Returned by Delete-documents collection API.
type DeleteResponse struct{}

// UpdateResponse includes metadata about just updated documents.
// Returned by Update-documents collection API.
type UpdateResponse struct{}

// ExplainResponse includes the query plan
// Tigris would run for a query.
type ExplainResponse = driver.ExplainResponse

// Error contains Tigris server error.
type Error driver.Error

func (e *Error) AsTigrisError(de *driver.Error) bool {
	e.TigrisError = de.TigrisError

	return true
}

func NewError(c code.Code, format string, a ...interface{}) *Error {
	if c == code.OK {
		return nil
	}

	return &Error{TigrisError: api.Errorf(c, format, a...)}
}
