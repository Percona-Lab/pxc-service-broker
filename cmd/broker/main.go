package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"

	"github.com/Percona-Lab/pxc-service-broker/broker"
	pxc "github.com/Percona-Lab/pxc-service-broker/pxccontroller"
)

func main() {
	brokerLogger := lager.NewLogger("pxc-broker")
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	brokerLogger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	brokerLogger.Info("Starting  broker")

	serviceBroker := &broker.PXCServiceBroker{
		/*InstanceCreators: map[string]broker.InstanceCreator{
			"shared":    localCreator,
			"dedicated": remoteRepo,
		},
		InstanceBinders: map[string]broker.InstanceBinder{
			"shared":    localRepo,
			"dedicated": remoteRepo,
		},
		Config: config,*/
	}

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: "",
		Password: "",
	}

	brokerAPI := brokerapi.New(serviceBroker, brokerLogger, brokerCredentials)

	http.Handle("/", brokerAPI)
	host := "localhost"
	port := "8081"

	log.Println("Create PXC")
	p, err := pxc.New()
	if err != nil {
		log.Println("Create PXC", err)
	}
	log.Println("Deploy Operator")
	err = p.DeployPXCOperator()
	if err != nil {
		log.Println("Deploy operator", err)
	}
	time.Sleep(80 * time.Second)
	err = p.DeployPXCCluster()
	if err != nil {
		log.Println("Deploy cluster", err)
	}
	brokerLogger.Fatal("http-listen", http.ListenAndServe(host+":"+port, nil))
}
