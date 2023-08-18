package checkpoint

import (
	"os"
	"strings"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func egress2array(annotation string) []string {

	res := strings.Split(annotation, ",")
	return res
}

func inPod() bool {

	if k := os.Getenv("KUBERNETES_PORT"); k == "" {
		return false
	}

	return true
}

func getRestConfig() *rest.Config {

	var err error

	if inPod() {
		restconfig, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	} else {
		kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		restconfig, err = kubeconfig.ClientConfig()
		if err != nil {
			panic(err)
		}
	}

	return restconfig
}
