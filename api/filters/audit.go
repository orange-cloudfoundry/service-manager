package filters

import (
	"fmt"

	"github.com/Peripli/service-manager/pkg/audit"
	"github.com/Peripli/service-manager/pkg/web"
)

type AuditFilter struct {
	Backend audit.Backend
}

func (*AuditFilter) Name() string {
	return "AuditFilter"
}

func (f *AuditFilter) Run(req *web.Request, next web.Handler) (*web.Response, error) {
	req, event, e := audit.RequestWithEvent(req)
	if e != nil {
		return nil, fmt.Errorf("failed to create audit event: %v", e)
	}

	event.State = audit.StatePreRequest
	f.Backend.Process(event)

	resp, err := next.Handle(req)
	if err == nil {
		event.State = audit.StatePostRequest
		event.ResponseStatus = resp.StatusCode
		event.ResponseObject = resp.Body
	}
	if user, exists := web.UserFromContext(req.Context()); exists {
		event.User = user.Name
	}
	f.Backend.Process(event)
	return resp, err
}

func (*AuditFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path("/**"),
			},
		},
	}
}