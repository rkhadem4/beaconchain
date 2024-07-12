package types

type AccountDashboard struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
type ValidatorDashboard struct {
	Id             uint64        `json:"id"`
	Name           string        `json:"name"`
	PublicIds      []VDBPublicId `json:"public_ids,omitempty"`
	Archived       bool          `json:"archived"`
	ArchivedReason string        `json:"archived_reason,omitempty"` // dashboard_limit, validator_limit, group_limit, none
	ValidatorCount uint64        `json:"validator_count"`
	GroupCount     uint64        `json:"group_count"`
}

type UserDashboardsData struct {
	ValidatorDashboards []ValidatorDashboard `json:"validator_dashboards"`
	AccountDashboards   []AccountDashboard   `json:"account_dashboards"`
}

type GetUserDashboardsResponse ApiDataResponse[UserDashboardsData]
