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

type DeleteDataPersistenceRuleSpectraS3Request struct {
    dataPersistenceRuleId string
    queryParams *url.Values
}

func NewDeleteDataPersistenceRuleSpectraS3Request(dataPersistenceRuleId string) *DeleteDataPersistenceRuleSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteDataPersistenceRuleSpectraS3Request{
        dataPersistenceRuleId: dataPersistenceRuleId,
        queryParams: queryParams,
    }
}




func (DeleteDataPersistenceRuleSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteDataPersistenceRuleSpectraS3Request *DeleteDataPersistenceRuleSpectraS3Request) Path() string {
    return "/_rest_/data_persistence_rule/" + deleteDataPersistenceRuleSpectraS3Request.dataPersistenceRuleId
}

func (deleteDataPersistenceRuleSpectraS3Request *DeleteDataPersistenceRuleSpectraS3Request) QueryParams() *url.Values {
    return deleteDataPersistenceRuleSpectraS3Request.queryParams
}

func (DeleteDataPersistenceRuleSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteDataPersistenceRuleSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteDataPersistenceRuleSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}