package id_transfer_api

import (
	"context"
	"github.com/lixinio/weixin/utils"
)

const (
	apiUnionID2ExternalUserID = "/cgi-bin/idconvert/unionid_to_external_userid"
)

type idTransferApi struct {
	*utils.Client
}

func NewApi(client *utils.Client) *idTransferApi {
	return &idTransferApi{Client: client}
}

type SubjectType int

const (
	// 0表示主体名称是企业的 (默认)
	SubjectTypeEnterprise = SubjectType(0)
	// 1表示主体名称是服务商的
	SubjectTypeProvider = SubjectType(1)
)

type UnionID2ExternalUserIDParam struct {
	UnionID     string      `json:"unionid"`
	OpenID      string      `json:"openid"`
	SubjectType SubjectType `json:"subject_type"` // 小程序或公众号的主体类型
}

type UnionID2ExternalUserIDResponse struct {
	utils.WeixinError
	ExternalUserID string `json:"external_userid"`
	PendingID      string `json:"pending_userid"`
}

// https://developer.work.weixin.qq.com/document/path/95900
// UnionID转换成企业第三方的external_userid
func (api *idTransferApi) UnionID2ExternalUserID(
	ctx context.Context,
	unionID, openID string,
	subjectType SubjectType,
) (*UnionID2ExternalUserIDResponse, error) {
	result := &UnionID2ExternalUserIDResponse{}

	if err := api.Client.HTTPPostJson(
		ctx,
		apiUnionID2ExternalUserID,
		&UnionID2ExternalUserIDParam{
			UnionID:     unionID,
			OpenID:      openID,
			SubjectType: subjectType,
		},
		result,
	); err != nil {
		return nil, err
	}

	return result, nil
}
