package bundle

import (
	"fmt"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle/apierrors"
	"github.com/erda-project/erda/pkg/httputil"
)

// FetchPublisher 获取 publisher 详情
func (b *Bundle) FetchPublisher(publisherID uint64) (*apistructs.PublisherDTO, error) {
	host, err := b.urls.CMDB()
	if err != nil {
		return nil, err
	}
	hc := b.hc
	var publisherResp apistructs.PublisherDetailResponse
	resp, err := hc.Get(host).Path(fmt.Sprintf("/api/publishers/%d", publisherID)).
		Header("Internal-Client", "bundle").
		Do().JSON(&publisherResp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !resp.IsOK() || !publisherResp.Success {
		return nil, toAPIError(resp.StatusCode(), publisherResp.Error)
	}
	return &publisherResp.Data, nil
}

func (b *Bundle) GetUserRelationPublisher(userID string, orgID string) (*apistructs.PagingPublisherDTO, error) {
	host, err := b.urls.CMDB()
	if err != nil {
		return nil, err
	}
	hc := b.hc
	var publisherResp apistructs.PublisherListResponse
	resp, err := hc.Get(host).Path(fmt.Sprintf("/api/publishers/actions/list-my-publishers")).
		Header("Internal-Client", "bundle").
		Header("USER-ID", userID).
		Header(httputil.OrgHeader, orgID).
		Do().JSON(&publisherResp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !resp.IsOK() || !publisherResp.Success {
		return nil, toAPIError(resp.StatusCode(), publisherResp.Error)
	}
	return &publisherResp.Data, nil
}
