// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"

	"code.gitea.io/gitea/modules/structs"
)

// PullRequest is equal to structs.PullRequest
type PullRequest = structs.PullRequest

// ListRepoPullRequests list PRs of one repository
func (c *Client) ListRepoPullRequests(owner, repo string, opt structs.ListPullRequestsOptions) ([]*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	prs := make([]*PullRequest, 0, 10)
	return prs, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/pulls", owner, repo), jsonHeader, bytes.NewReader(body), &prs)
}

// GetPullRequest get information of one PR
func (c *Client) GetPullRequest(owner, repo string, index int64) (*PullRequest, error) {
	pr := new(PullRequest)
	return pr, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, index), nil, nil, pr)
}

// CreatePullRequest create pull request with options
func (c *Client) CreatePullRequest(owner, repo string, opt structs.CreatePullRequestOption) (*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	pr := new(PullRequest)
	return pr, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/pulls", owner, repo),
		jsonHeader, bytes.NewReader(body), pr)
}

// EditPullRequest modify pull request with PR id and options
func (c *Client) EditPullRequest(owner, repo string, index int64, opt structs.EditPullRequestOption) (*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	pr := new(PullRequest)
	return pr, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, index),
		jsonHeader, bytes.NewReader(body), pr)
}

// MergePullRequest merge a PR to repository by PR id
func (c *Client) MergePullRequest(owner, repo string, index int64) error {
	_, err := c.getResponse("POST", fmt.Sprintf("/repos/%s/%s/pulls/%d/merge", owner, repo, index), nil, nil)
	return err
}

// IsPullRequestMerged test if one PR is merged to one repository
func (c *Client) IsPullRequestMerged(owner, repo string, index int64) (bool, error) {
	statusCode, err := c.getStatusCode("GET", fmt.Sprintf("/repos/%s/%s/pulls/%d/merge", owner, repo, index), nil, nil)

	if err != nil {
		return false, err
	}

	return statusCode == 204, nil
}
