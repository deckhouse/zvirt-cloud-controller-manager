package zvirtapi

import (
	"errors"

	ovirtclientlog "github.com/ovirt/go-ovirt-client-log/v3"
	ovirtclient "github.com/ovirt/go-ovirt-client/v3"
)

var ErrNotFound = errors.New("not found")

type ZvirtCloudAPI struct {
	ComputeSvc ComputeService
}

func NewZvirtCloudAPI(apiURL, username, password string, insecure bool) (*ZvirtCloudAPI, error) {
	logger := ovirtclientlog.NewGoLogger()

	tls := ovirtclient.TLS()
	tls.CACertsFromSystem()

	if insecure {
		tls.Insecure()
	}

	client, err := ovirtclient.New(
		apiURL,
		username,
		password,
		tls,
		logger,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &ZvirtCloudAPI{
		ComputeSvc: *NewComputeService(client),
	}, nil
}