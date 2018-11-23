/*
 *    Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Package visibility contains logic for service visibility API
package visibility

import (
	"net/http"

	"github.com/Peripli/service-manager/pkg/log"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/pkg/web"
)

// Controller broker controller
type Controller struct{}

func (c *Controller) getAllVisibilities(r *web.Request) (*web.Response, error) {
	ctx := r.Context()
	log.C(ctx).Debug("Getting all brokers")

	v := []*types.Visibility{
		&types.Visibility{
			CatalogPlanID: "1f400825-1434-5278-9913-dfcf63fcd647",
			Labels: []*types.Label{
				&types.Label{
					Key:    "organization_guid",
					Values: []string{"8c09d1a9-6acb-48a4-9b7f-bfbdb8373165", "8bc3cef7-7500-4af6-8689-38994632647e", "577ab247-805d-4f72-9c96-97e9a20d4dec"},
				},
				&types.Label{
					Key:    "expiration",
					Values: []string{"1h"},
				},
			},
		},
		&types.Visibility{
			CatalogPlanID: "2f400825-1234-5278-9013-dfcf63fcd646",
			Labels: []*types.Label{
				&types.Label{
					Key:    "organization_guid",
					Values: []string{"8c09d1a9-6acb-48a4-9b7f-bfbdb8373165"},
				},
			},
		},
		&types.Visibility{
			CatalogPlanID: "3f400825-1234-5278-2913-dfcf63fcd647",
			Labels: []*types.Label{
				&types.Label{
					Key:    "organization_guid",
					Values: []string{"8bc3cef7-7500-4af6-8689-38994632647e"},
				},
			},
		},
		&types.Visibility{
			CatalogPlanID: "4f400825-1234-5278-9513-dfcf63fcd646",
			Labels: []*types.Label{
				&types.Label{
					Key:    "organization_guid",
					Values: []string{"8bc3cef7-7500-4af6-8689-38994632647e", "577ab247-805d-4f72-9c96-97e9a20d4dec"},
				},
				&types.Label{
					Key:    "expiration",
					Values: []string{"1h"},
				},
			},
		},
	}

	return util.NewJSONResponse(http.StatusOK, &types.Visibilities{
		Visibilities: v,
	})
}
