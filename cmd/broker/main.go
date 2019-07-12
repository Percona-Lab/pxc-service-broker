package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"

	"github.com/Percona-Lab/pxc-service-broker/broker"
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
	brokerLogger.Fatal("http-listen", http.ListenAndServe(host+":"+port, nil))
}
