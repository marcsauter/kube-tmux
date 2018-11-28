package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("[%s/%s]", config.CurrentContext, config.Contexts[config.CurrentContext].Namespace)
}
