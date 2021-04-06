package cmdb

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
)

const serviceCMDB = "cmdb"

type provider struct{}

func init() { servicehub.RegisterProvider(serviceCMDB, &provider{}) }

func (p *provider) Service() []string                 { return []string{serviceCMDB} }
func (p *provider) Dependencies() []string            { return []string{} }
func (p *provider) Init(ctx servicehub.Context) error { return nil }
func (p *provider) Run(ctx context.Context) error     { return Initialize() }
func (p *provider) Creator() servicehub.Creator {
	return func() servicehub.Provider { return &provider{} }
}
