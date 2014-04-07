package models

import (
    "net/url"
    "ds3/net"
)

type PutObjectRequest struct {
    bucketName, objectName string
    content net.SizedReadCloser
}

func NewPutObjectRequest(bucketName, objectName string, content net.SizedReadCloser) *PutObjectRequest {
    return &PutObjectRequest{bucketName, objectName, content}
}

func (PutObjectRequest) Verb() net.HttpVerb {
    return net.PUT
}

func (self PutObjectRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (PutObjectRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (self PutObjectRequest) GetContentStream() net.SizedReadCloser {
    return self.content
}

