package models

import "time"

type CreateSubUIDRequest struct {
	Username   string `json:"username"`         // Give a username of the new sub user id. 6-16 characters, must include both numbers and letters. cannot be the same as the exist or deleted one.
	MemberType int64  `json:"memberType"`       // 1: normal sub account, 6: custodial sub account
	Switch     int64  `json:"switch,omitempty"` // [optional]0: turn off quick login (default) 1: turn on quick login.
	Note       string `json:"note,omitempty"`   // [optional]Set a remark
}

type CreateSubUIDResponse struct {
	Uid        string `json:"uid"`        // Sub user Id
	Username   string `json:"username"`   // Give a username of the new sub user id. 6-16 characters, must include both numbers and letters. cannot be the same as the exist or deleted one.
	MemberType int64  `json:"memberType"` // 1: normal sub account, 6: custodial sub account
	Status     int64  `json:"status"`     // The status of the user account 1: normal 2: login banned 4: frozen
	Remark     string `json:"remark"`     // The remark
}

type GetSubUIDListResponse struct {
	SubMembers []struct {
		Uid        string `json:"uid"`        // Sub user Id
		Username   string `json:"username"`   // Username
		MemberType int64  `json:"memberType"` // 1: normal sub account, 6: custodial sub account
		Status     int64  `json:"status"`     // The status of the user account 1: normal 2: login banned 4: frozen
		Remark     string `json:"remark"`     // The remark
	} `json:"subMembers"` // Object
}

type FrozeSubUIDRequest struct {
	Subuid int64 `json:"subuid"` // Sub user Id
	Frozen int64 `json:"frozen"` // 0：unfreeze, 1：freeze
}

type GetAPIKeyInformationResponse struct {
	Id          string `json:"id"`       // Unique id. Internal used
	Note        string `json:"note"`     // The remark
	ApiKey      string `json:"apiKey"`   // Api key
	ReadOnly    int64  `json:"readOnly"` // 0：Read and Write. 1：Read only
	Secret      string `json:"secret"`   // Always ""
	Permissions struct {
		ContractTrade []string `json:"ContractTrade"` // Permission of contract trade
		Spot          []string `json:"Spot"`          // Permission of spot
		Wallet        []string `json:"Wallet"`        // Permission of wallet
		Options       []string `json:"Options"`       // Permission of USDC Contract. It supports trade option and usdc perpetual.
		Derivatives   []string `json:"Derivatives"`   // Permission of derivatives
		CopyTrading   []string `json:"CopyTrading"`   // Permission of copytrade. Not applicable to sub account, always []
		BlockTrade    []string `json:"BlockTrade"`    // Permission of blocktrade. Not applicable to sub account, always []
		Exchange      []string `json:"Exchange"`      // Permission of exchange
		NFT           []string `json:"NFT"`           // Permission of NFT. Not applicable to sub account, always []
	} `json:"permissions"` // The types of permission
	Ips           []string  `json:"ips"`           // IP bound
	Type          int64     `json:"type"`          // The type of api key. 1：personal, 2：connected to the third-party app
	DeadlineDay   int64     `json:"deadlineDay"`   // The remaining valid days of api key. Only for those api key with no IP bound or the password has been changed
	ExpiredAt     time.Time `json:"expiredAt"`     // The expiry day of the api key. Only for those api key with no IP bound or the password has been changed
	CreatedAt     time.Time `json:"createdAt"`     // The create day of the api key
	Unified       int64     `json:"unified"`       // Whether the account to which the api key belongs is a unified margin account. 0：regular account; 1：unified margin account
	Uta           int64     `json:"uta"`           // Whether the account to which the account upgrade to unified trade account. 0：regular account; 1：unified trade account
	UserID        int64     `json:"userID"`        // User ID
	InviterID     int64     `json:"inviterID"`     // Inviter ID (the UID of the account which invited this account to the platform)
	VipLevel      string    `json:"vipLevel"`      // VIP Level
	MktMakerLevel string    `json:"mktMakerLevel"` // Market maker level
	AffiliateID   int64     `json:"affiliateID"`   // Affiliate Id. 0 represents that there is no binding relationship.
	RsaPublicKey  string    `json:"rsaPublicKey"`  // Rsa public key
}
