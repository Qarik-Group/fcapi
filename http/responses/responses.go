package responses

import (
	cf "github.com/cloudfoundry-community/go-cfclient"
)

type Pageref struct {
	Href string `json:"href"`
}

type Pagination struct {
	First        *Pageref `json:"first,omitempty"`
	Last         *Pageref `json:"last,omitempty"`
	Next         *Pageref `json:"next,omitempty"`
	Previous     *Pageref `json:"previous,omitempty"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type RoleCollection struct {
	Pagination *Pagination `json:"pagination"`
	Resources  []cf.V3Role `json:"resources"`
}

type OrgCollection struct {
	Pagination *Pagination         `json:"pagination"`
	Resources  []cf.V3Organization `json:"resources"`
}

type SpaceCollection struct {
	Pagination *Pagination  `json:"pagination"`
	Resources  []cf.V3Space `json:"resources"`
}

type RootResp struct {
	Links *RootLinkCollection `json:"links"`
}

type RootLinkCollection struct {
	Self              *Rootref `json:"self"`
	BitsService       *Rootref `json:"bits_service"`
	CloudControllerV2 *Rootref `json:"cloud_controller_v2"`
	CloudControllerV3 *Rootref `json:"cloud_controller_v3"`
	NetworkPolicyV0   *Rootref `json:"network_policy_v0"`
	NetworkPolicyV1   *Rootref `json:"network_policy_v1"`
	Login             *Rootref `json:"login"`
	UAA               *Rootref `json:"uaa"`
	Credhub           *Rootref `json:"credhub"`
	Routing           *Rootref `json:"routing"`
	Logging           *Rootref `json:"loggin"`
	LogCache          *Rootref `json:"log_cache"`
	LogStream         *Rootref `json:"log_stream"`
	AppSSH            *Rootref `json:"app_ssh"`
}

type Rootref struct {
	Href string            `json:"href"`
	Meta map[string]string `json:"meta,omitempty"`
}

type UserResp struct {
	Meta   cf.Meta `json:"metadata"`
	Entity cf.User `json:"entity"`
}

type TaskReq struct {
	Command string `json:"command"`
	Disk    int    `json:"disk_in_mb"`
	Mem     int    `json:"memory_in_mb"`
}
