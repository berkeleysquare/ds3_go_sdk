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
    "ds3/networking"
    "net/http"
)

type GetActiveJobsSpectraS3Response struct {
    ActiveJobList ActiveJobList
    Headers *http.Header
}

func (getActiveJobsSpectraS3Response *GetActiveJobsSpectraS3Response) parse(webResponse networking.WebResponse) error {
        return parseResponsePayload(webResponse, &getActiveJobsSpectraS3Response.ActiveJobList)
}

func NewGetActiveJobsSpectraS3Response(webResponse networking.WebResponse) (*GetActiveJobsSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body GetActiveJobsSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
