// Copyright © 2023 Horizoncd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package role

import (
	"github.com/gin-gonic/gin"
	"github.com/horizoncd/horizon/core/controller/role"
	"github.com/horizoncd/horizon/pkg/server/response"
)

type API struct {
	roleCtrl role.Controller
}

func NewAPI(controller role.Controller) *API {
	return &API{roleCtrl: controller}
}

func (a *API) ListRole(c *gin.Context) {
	roles, err := a.roleCtrl.ListRole(c)
	if err != nil {
		response.AbortWithInternalError(c, err.Error())
		return
	}
	response.SuccessWithData(c, roles)
}
