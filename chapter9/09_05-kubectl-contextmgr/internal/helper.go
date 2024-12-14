package helper

import (
	"fmt"
	"log"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Context struct {
	Status    bool
	Name      string
	Cluster   string
	IsCurrent bool
}

func LoadKubeConfig(kubeconfig string) (*api.Config, error) {
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		fmt.Printf("Error loading kubeconfig: %v\n", err)
		return nil, err
	}

	return config, nil
}

func GetContext(config *api.Config) []Context {
	var contexts []Context
	for name, context := range config.Contexts {
		contexts = append(contexts, Context{
			Name:      name,
			Cluster:   context.Cluster,
			Status:    isClusterReachable(name, config),
			IsCurrent: name == config.CurrentContext,
		})
	}
	return contexts
}

func SwitchContext(contextName string, config *api.Config) (bool, error) {
	config.CurrentContext = contextName
	err := clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), *config, false)
	if err != nil {
		fmt.Printf("Error switching context: %v\n", err)
		return false, err
	}
	return true, nil
}

func isClusterReachable(contextName string, config *api.Config) bool {
	clientConfig := clientcmd.NewNonInteractiveClientConfig(*config, contextName, &clientcmd.ConfigOverrides{}, nil)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		log.Printf("Error creating rest config: %v\n", err)
		return false
	}

	restConfig.Timeout = 5 * time.Second

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Printf("Error creating clientset: %v\n", err)
		return false
	}

	_, err = clientset.Discovery().ServerVersion()
	return err == nil
}
