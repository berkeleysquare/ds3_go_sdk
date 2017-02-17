package models

import (
    "strings"
    "net/http"
    "ds3/networking"
)

type PutPartResponse struct {
    ETag string
}

func NewPutPartResponse(response networking.Response) (*PutPartResponse, error) {
    if err := checkStatusCode(response, http.StatusOK); err != nil {
        return nil, err
    } else {
        etags := (*response.Header())["etag"]
        var etag string
        if len(etags) > 0 {
            etag = strings.Trim(etags[0], "\"")
        }
        return &PutPartResponse{etag}, nil
    }
}

