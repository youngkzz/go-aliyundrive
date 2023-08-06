/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package aliyundrive

import (
	"context"
	"net/http"
)

// GetShareLink 获取分享链接
func (r *ShareLinkService) GetShareLink(ctx context.Context, request *GetShareLinkReq) (*GetShareLinkResp, error) {
	request.SyncToHomepage = false

	req := &RawRequestReq{
		Scope:  "ShareLink",
		API:    "GetShareLink",
		Method: http.MethodPost,
		URL:    "https://api.aliyundrive.com/adrive/v2/share_link/create",
		Body:   request,
	}
	resp := new(getShareLinkResp)

	if _, err := r.cli.RawRequest(ctx, req, resp); err != nil {
		return nil, err
	}
	return &resp.GetShareLinkResp, nil
}

type GetShareLinkReq struct {
	DriveID        string   `json:"drive_id"`
	SharePwd       string   `json:"share_pwd"`
	Expiration     string   `json:"expiration"`
	FileIdList     []string `json:"file_id_list"`
	SyncToHomepage bool     `json:"sync_to_homepage"`
}

type GetShareLinkResp struct {
	DriveID    string `json:"drive_id"`
	Category   string `json:"category"`
	ShareID    string `json:"share_id"`
	ShareUrl   string `json:"share_url"`
	ShareName  string `json:"share_name"`
	Expired    bool   `json:"expired"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Status     string `json:"status"`
	ShareTitle string `json:"share_title"`
	Expiration string `json:"expiration"`
	SharePwd   string `json:"share_pwd"`
}

type getShareLinkResp struct {
	Message string `json:"message"`
	GetShareLinkResp
}
