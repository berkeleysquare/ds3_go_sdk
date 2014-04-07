package models

import (
    "net/url"
    "ds3/net"
)

type PutBucketRequest struct {
    bucketName string
}

func NewPutBucketRequest(bucketName string) *PutBucketRequest {
    return &PutBucketRequest{
        bucketName: bucketName,
    }
}

func (PutBucketRequest) Verb() net.HttpVerb {
    return net.PUT
}

func (self PutBucketRequest) Path() string {
    return "/" + self.bucketName
}

func (self PutBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (PutBucketRequest) GetContentStream() net.SizedReadCloser {
    return nil
}


