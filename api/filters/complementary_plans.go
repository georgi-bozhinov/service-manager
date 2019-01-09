/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filters

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tidwall/gjson"

	"github.com/Peripli/service-manager/pkg/util"

	"github.com/Peripli/service-manager/pkg/types"
	"github.com/gofrs/uuid"

	"github.com/Peripli/service-manager/pkg/log"

	"github.com/Peripli/service-manager/storage"

	"github.com/Peripli/service-manager/pkg/web"
)

// ComplementaryServicePlansFilter reconciles the state of the complementary plans offered by all service brokers registered in SM. The
// filter makes sure that a public visibility exists for each complementary plan present in SM DB.
// Complementary plans are marked as complementary is the catalog service plan metadata.
type ComplementaryServicePlansFilter struct {
	Repository storage.Repository
}

func (csp *ComplementaryServicePlansFilter) Name() string {
	return "ComplementaryPlansFilter"
}

func (csp *ComplementaryServicePlansFilter) Run(req *web.Request, next web.Handler) (*web.Response, error) {
	response, err := next.Handle(req)
	if err != nil {
		return nil, err
	}
	ctx := req.Context()
	brokerID := gjson.GetBytes(response.Body, "id").String()
	log.C(ctx).Debugf("Reconciling complementary plans for broker with id: %s", brokerID)
	if err := csp.Repository.InTransaction(ctx, func(ctx context.Context, storage storage.Warehouse) error {
		soRepository := storage.ServiceOffering()
		vRepository := storage.Visibility()

		catalog, err := soRepository.ListWithServicePlansByBrokerID(ctx, brokerID)
		if err != nil {
			return err
		}
		for _, serviceOffering := range catalog {
			for _, servicePlan := range serviceOffering.Plans {
				planID := servicePlan.ID
				isComplementary := gjson.GetBytes(servicePlan.Metadata, "complementary").Bool()
				hasPublicVisibility := false
				visibilitiesForPlan, err := vRepository.ListByServicePlanID(ctx, planID)
				if err != nil {
					return err
				}
				for _, visibility := range visibilitiesForPlan {
					if isComplementary {
						if visibility.PlatformID == "" {
							hasPublicVisibility = true
							continue
						} else {
							if err := vRepository.Delete(ctx, visibility.ID); err != nil {
								return err
							}
						}
					} else {
						if visibility.PlatformID == "" {
							if err := vRepository.Delete(ctx, visibility.ID); err != nil {
								return err
							}
						} else {
							continue
						}
					}
				}

				if isComplementary && !hasPublicVisibility {
					UUID, err := uuid.NewV4()
					if err != nil {
						return fmt.Errorf("could not generate GUID for visibility: %s", err)
					}

					currentTime := time.Now().UTC()
					planID, err := vRepository.Create(ctx, &types.Visibility{
						ID:            UUID.String(),
						ServicePlanID: servicePlan.ID,
						CreatedAt:     currentTime,
						UpdatedAt:     currentTime,
					})
					if err != nil {
						return util.HandleStorageError(err, "visibility", UUID.String())
					}

					log.C(ctx).Debugf("Created new public visibility for broker with id %s and plan with id %s", brokerID, planID)
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	log.C(ctx).Debugf("Successfully finished reconciling complementary plans for broker with id %s", brokerID)
	return response, nil
}

func (csp *ComplementaryServicePlansFilter) FilterMatchers() []web.FilterMatcher {
	return []web.FilterMatcher{
		{
			Matchers: []web.Matcher{
				web.Path(web.BrokersURL + "/**"),
				web.Methods(http.MethodPost, http.MethodPatch),
			},
		},
	}
}
