package broker

import (
	"context"
	"errors"

	"github.com/pivotal-cf/brokerapi"
)

const (
	PlanNameShared    = "shared-vm"
	PlanNameDedicated = "dedicated-vm"
)

type InstanceCredentials struct {
	Host     string
	Port     int
	Password string
}

type InstanceCreator interface {
	Create(instanceID string) error
	Destroy(instanceID string) error
	InstanceExists(instanceID string) (bool, error)
}

type InstanceBinder interface {
	Bind(instanceID string, bindingID string) (InstanceCredentials, error)
	Unbind(instanceID string, bindingID string) error
	InstanceExists(instanceID string) (bool, error)
}

type PXCServiceBroker struct {
	InstanceCreators map[string]InstanceCreator
	InstanceBinders  map[string]InstanceBinder
	//Config           brokerconfig.Config
}

func (psb *PXCServiceBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	planList := []brokerapi.ServicePlan{}
	for _, plan := range psb.plans() {
		planList = append(planList, *plan)
	}

	return []brokerapi.Service{
		brokerapi.Service{
			ID:          "pxc-service-broker-id",
			Name:        "percona-xtradb-cluster",
			Description: "database",
			Bindable:    true,
			Plans:       planList,
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName:         "PXC",
				LongDescription:     "PerconaXtraBDCluster",
				DocumentationUrl:    "",
				SupportUrl:          "",
				ImageUrl:            "",
				ProviderDisplayName: "percona",
			},
			Tags: []string{
				"pxc",
			},
		},
		/*
			brokerapi.Service{
				ID:          psb.Config.RedisConfiguration.ServiceID,
				Name:        psb.Config.RedisConfiguration.ServiceName,
				Description: psb.Config.RedisConfiguration.Description,
				Bindable:    true,
				Plans:       planList,
				Metadata: &brokerapi.ServiceMetadata{
					DisplayName:         psb.Config.RedisConfiguration.DisplayName,
					LongDescription:     psb.Config.RedisConfiguration.LongDescription,
					DocumentationUrl:    psb.Config.RedisConfiguration.DocumentationURL,
					SupportUrl:          psb.Config.RedisConfiguration.SupportURL,
					ImageUrl:            fmt.Sprintf("data:image/png;base64,%s", psb.Config.RedisConfiguration.IconImage),
					ProviderDisplayName: psb.Config.RedisConfiguration.ProviderDisplayName,
				},
				Tags: []string{
					"pivotal",
					"redis",
				},
			},*/
	}, nil
}

//Provision ...
func (psb *PXCServiceBroker) Provision(ctx context.Context, instanceID string, serviceDetails brokerapi.ProvisionDetails, asyncAllowed bool) (spec brokerapi.ProvisionedServiceSpec, err error) {
	spec = brokerapi.ProvisionedServiceSpec{}

	/*if psb.instanceExists(instanceID) {
		return spec, brokerapi.ErrInstanceAlreadyExists
	}

	if serviceDetails.PlanID == "" {
		return spec, errors.New("plan_id required")
	}

	planIdentifier := ""
	for key, plan := range psb.plans() {
		if plan.ID == serviceDetails.PlanID {
			planIdentifier = key
			break
		}
	}

	if planIdentifier == "" {
		return spec, errors.New("plan_id not recognized")
	}

	instanceCreator, ok := psb.InstanceCreators[planIdentifier]
	if !ok {
		return spec, errors.New("instance creator not found for plan")
	}

	err = instanceCreator.Create(instanceID)
	if err != nil {
		return spec, err
	}*/

	return spec, nil
}

func (psb *PXCServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	spec := brokerapi.DeprovisionServiceSpec{}

	/*for _, instanceCreator := range psb.InstanceCreators {
		instanceExists, _ := instanceCreator.InstanceExists(instanceID)
		if instanceExists {
			return spec, instanceCreator.Destroy(instanceID)
		}
	}*/
	return spec, brokerapi.ErrInstanceDoesNotExist
}

func (psb *PXCServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	/*binding := brokerapi.Binding{}

	for _, repo := range psb.InstanceBinders {
		instanceExists, _ := repo.InstanceExists(instanceID)
		if instanceExists {
			instanceCredentials, err := repo.Bind(instanceID, bindingID)
			if err != nil {
				return binding, err
			}
			credentialsMap := map[string]interface{}{
				"host":     instanceCredentials.Host,
				"port":     instanceCredentials.Port,
				"password": instanceCredentials.Password,
			}

			binding.Credentials = credentialsMap
			return binding, nil
		}
	}*/
	return brokerapi.Binding{}, brokerapi.ErrInstanceDoesNotExist
}

func (psb *PXCServiceBroker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	/*for _, repo := range psb.InstanceBinders {
		instanceExists, _ := repo.InstanceExists(instanceID)
		if instanceExists {
			err := repo.Unbind(instanceID, bindingID)
			if err != nil {
				return brokerapi.UnbindSpec{}, brokerapi.ErrBindingDoesNotExist
			}
			return brokerapi.UnbindSpec{}, nil
		}
	}*/

	return brokerapi.UnbindSpec{}, brokerapi.ErrInstanceDoesNotExist
}

func (psb *PXCServiceBroker) plans() map[string]*brokerapi.ServicePlan {
	plans := map[string]*brokerapi.ServicePlan{}
	/*
		if psb.Config.SharedEnabled() {
			plans["shared"] = &brokerapi.ServicePlan{
				ID:          psb.Config.RedisConfiguration.SharedVMPlanID,
				Name:        PlanNameShared,
				Description: "This plan provides a Redis server on a shared VM configured for data persistence.",
				Metadata: &brokerapi.ServicePlanMetadata{
					Bullets: []string{
						"Each instance shares the same VM",
						"Single dedicated Redis process",
						"Suitable for development & testing workloads",
					},
					DisplayName: "Shared-VM",
				},
			}
		}

		if psb.Config.DedicatedEnabled() {
			plans["dedicated"] = &brokerapi.ServicePlan{
				ID:          psb.Config.RedisConfiguration.DedicatedVMPlanID,
				Name:        PlanNameDedicated,
				Description: "This plan provides a Redis server configured for data persistence. ",
				Metadata: &brokerapi.ServicePlanMetadata{
					Bullets: []string{
						"Dedicated VM per instance",
						"Single dedicated Redis process",
						"Suitable for production workloads",
					},
					DisplayName: "Dedicated-VM",
				},
			}
		}*/

	return plans
}

func (psb *PXCServiceBroker) instanceExists(instanceID string) bool {
	/*for _, instanceCreator := range psb.InstanceCreators {
		instanceExists, _ := instanceCreator.InstanceExists(instanceID)
		if instanceExists {
			return true
		}
	}*/
	return false
}

// LastOperation ...
// If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
func (psb *PXCServiceBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("not implemented")
}

func (psb *PXCServiceBroker) Update(cxt context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, errors.New("not implemented")
}

func (psb *PXCServiceBroker) GetBinding(ctx context.Context, instanceID, bindingID string) (brokerapi.GetBindingSpec, error) {
	return brokerapi.GetBindingSpec{}, errors.New("not implemented")
}

func (psb *PXCServiceBroker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	return brokerapi.GetInstanceDetailsSpec{}, errors.New("not implemented")
}

func (psb *PXCServiceBroker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, errors.New("not implemented")
}
