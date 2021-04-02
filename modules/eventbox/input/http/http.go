package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/erda-project/erda/modules/eventbox/input"
	"github.com/erda-project/erda/modules/eventbox/monitor"
	stypes "github.com/erda-project/erda/modules/eventbox/server/types"
	"github.com/erda-project/erda/modules/eventbox/types"

	"github.com/sirupsen/logrus"
)

type HttpInput struct {
	handler input.Handler
	stopch  chan struct{}
}

func New() (*HttpInput, error) {
	return &HttpInput{
		stopch: make(chan struct{}),
	}, nil
}

func (h *HttpInput) Start(handler input.Handler) error {
	h.handler = handler
	<-h.stopch
	return nil
}

func (h *HttpInput) Stop() error {
	h.stopch <- struct{}{}
	return nil
}

func (h *HttpInput) Name() string {
	return "HTTP"
}

func (h *HttpInput) GetHTTPEndPoints() []stypes.Endpoint {
	return []stypes.Endpoint{
		{"/message/create", http.MethodPost, h.createMessage},
	}
}

func (h *HttpInput) createMessage(ctx context.Context, req *http.Request, vars map[string]string) (stypes.Responser, error) {
	var m types.Message
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		return stypes.HTTPResponse{Status: http.StatusBadRequest, Content: "unmarshal message failed!"}, err
	}
	logrus.Debugf("%s input message timestamp:%d", h.Name(), m.Time)

	monitor.Notify(monitor.MonitorInfo{Tp: monitor.HTTPInput})
	e := h.handler(&m)
	resp := genResponse(e)
	return resp, nil
}
