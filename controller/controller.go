/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"errors"
	"fmt"

	"github.com/eleanorrigby/openservicebroker_exper/client"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

type errNoSuchInstance struct {
	instanceID string
}

func (e errNoSuchInstance) Error() string {
	return fmt.Sprintf("no such instance with ID %s", e.instanceID)
}

type jenkinsServiceInstance struct {
	//Name       string
	//Credential *brokerapi.Credential
}

type jenkinsController struct {
	//rwMutex     sync.RWMutex
	//instanceMap map[string]*jenkinsServiceInstance
}

// CreateController creates an instance of a User Provided service broker controller.
func CreateController() controller.Controller {
	//var instanceMap = make(map[string]*jenkinsServiceInstance)
	return &jenkinsController{}
	//instanceMap: instanceMap,
	//}
}

func (c *jenkinsController) Catalog() (*brokerapi.Catalog, error) {
	return &brokerapi.Catalog{
		Services: []*brokerapi.Service{
			{
				Name:        "jenkins-service",
				ID:          "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468",
				Description: "Jenkins as a service",
				Plans: []brokerapi.ServicePlan{{
					Name:        "default",
					ID:          "86064792-7ea2-467b-af93-ac9694d96d52",
					Description: "Free Use Plan",
					Free:        true,
				},
				},
				Bindable: true,
			},
		},
	}, nil
}

func (c *jenkinsController) CreateServiceInstance(
	id string,
	req *brokerapi.CreateServiceInstanceRequest,
) (*brokerapi.CreateServiceInstanceResponse, error) {
	/*
		credString, ok := req.Parameters["credentials"]
		c.rwMutex.Lock()
		defer c.rwMutex.Unlock()
		if ok {
			jsonCred, err := json.Marshal(credString)
			if err != nil {
				glog.Errorf("Failed to marshal credentials: %v", err)
				return nil, err
			}
			var cred brokerapi.Credential
			err = json.Unmarshal(jsonCred, &cred)

			c.instanceMap[id] = &jenkinsServiceInstance{
				Name:       id,
				Credential: &cred,
			}
		} else {
			c.instanceMap[id] = &jenkinsServiceInstance{
				Name: id,
				Credential: &brokerapi.Credential{
					"special-key-1": "special-value-1",
					"special-key-2": "special-value-2",
				},
			}
		}*/
	if err := client.Install(releaseName(id), id); err != nil {
		return nil, err
	}
	glog.Infof("Created Jenkins Service Instance:\n\n")
	return &brokerapi.CreateServiceInstanceResponse{}, nil
}

func (c *jenkinsController) GetServiceInstance(id string) (string, error) {
	return "", errors.New("Unimplemented")
}

func (c *jenkinsController) RemoveServiceInstance(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	/*c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	_, ok := c.instanceMap[id]
	if ok {
		delete(c.instanceMap, id)
		return &brokerapi.DeleteServiceInstanceResponse{}, nil
	}
	*/
	if err := client.Delete(releaseName(id)); err != nil {
		return nil, err
	}

	return &brokerapi.DeleteServiceInstanceResponse{}, nil
}

//This might not work for now.
func (c *jenkinsController) Bind(
	instanceID,
	bindingID string,
	req *brokerapi.BindingRequest,
) (*brokerapi.CreateServiceBindingResponse, error) {
	port := "8080"
	username := "admin"
	password, err := client.GetPassword(releaseName(instanceID), instanceID)
	if err != nil {
		return nil, err
	}
	return &brokerapi.CreateServiceBindingResponse{
		Credentials: brokerapi.Credential{
			"username": username,
			"password": password,
			"port":     port,
		},
	}, nil
}

func (c *jenkinsController) UnBind(instanceID string, bindingID string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}

func releaseName(id string) string {
	return "i-" + id
}
