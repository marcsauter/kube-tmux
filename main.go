package main

import (
	"os"
	"text/template"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	format := "[{{.Context}}/{{.Namespace}}]"
	if len(os.Args) > 1 {
		format = os.Args[1]
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		os.Exit(1)
	}
	template.Must(template.New("tmux").Parse(format)).Execute(os.Stdout, struct {
		Context, Namespace string
	}{
		config.CurrentContext,
		config.Contexts[config.CurrentContext].Namespace,
	})
}
