package domain

import "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/connectivity"

func NewCdnService(client *connectivity.CdnetworksClient) CdnService {
	return CdnService{client: client}
}

type CdnService struct {
	client *connectivity.CdnetworksClient
}
