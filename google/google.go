package google

import (
	gce "github.com/scorelab/gocloud-v2/compute/gce"
	googlestorage "github.com/scorelab/gocloud-v2/storage/googlestorage"
)

//Google  struct represents google cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
}
