// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This code is auto-generated, do not modify

package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type CompactAllPoolsSpectraS3Request struct {
    priority Priority
    queryParams *url.Values
}

func NewCompactAllPoolsSpectraS3Request() *CompactAllPoolsSpectraS3Request {
    queryParams := &url.Values{}
    queryParams.Set("operation", "compact")

    return &CompactAllPoolsSpectraS3Request{
        queryParams: queryParams,
    }
}

func (compactAllPoolsSpectraS3Request *CompactAllPoolsSpectraS3Request) WithPriority(priority Priority) *CompactAllPoolsSpectraS3Request {
    compactAllPoolsSpectraS3Request.priority = priority
    compactAllPoolsSpectraS3Request.queryParams.Set("priority", priority.String())
    return compactAllPoolsSpectraS3Request
}



func (CompactAllPoolsSpectraS3Request) Verb() networking.HttpVerb {
    return networking.PUT
}

func (compactAllPoolsSpectraS3Request *CompactAllPoolsSpectraS3Request) Path() string {
    return "/_rest_/pool"
}

func (compactAllPoolsSpectraS3Request *CompactAllPoolsSpectraS3Request) QueryParams() *url.Values {
    return compactAllPoolsSpectraS3Request.queryParams
}

func (CompactAllPoolsSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (CompactAllPoolsSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (CompactAllPoolsSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}