package bos

import "github.com/baidubce/bce-sdk-go/services/bos/api"

type objectPageStatus struct {
	delimiter string
	marker    string
	maxKeys   int
	prefix    string
}

func (i *objectPageStatus) ContinuationToken() string {
	return i.marker
}

type storagePageStatus struct {
	buckets []api.BucketSummaryType
}

func (i *storagePageStatus) ContinuationToken() string {
	return ""
}
