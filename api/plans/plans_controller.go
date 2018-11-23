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

// Package plans contains logic for service plans API
package plans

import (
	"net/http"

	"github.com/Peripli/service-manager/pkg/log"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/pkg/web"
)

// Controller broker controller
type Controller struct{}

func (c *Controller) getAllPlans(r *web.Request) (*web.Response, error) {
	ctx := r.Context()
	log.C(ctx).Debug("Getting all plans")

	p := []*types.Plan{
		&types.Plan{
			CatalogID: "1f400825-1434-5278-9913-dfcf63fcd647",
		},
		&types.Plan{
			CatalogID: "2f400825-1234-5278-9013-dfcf63fcd646",
		},
		&types.Plan{
			CatalogID: "3f400825-1234-5278-2913-dfcf63fcd647",
		},
		&types.Plan{
			CatalogID: "4f400825-1234-5278-9513-dfcf63fcd646",
		},
	}

	return util.NewJSONResponse(http.StatusOK, &types.Plans{
		Plans: p,
	})
}
