package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func main() {
	configFlags := genericclioptions.NewConfigFlags(true)
	rootCmd := &cobra.Command{
		Use:   "pod-fetcher",
		Short: "Fetch pods from a Kubernetes cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getPods(configFlags)
		},
	}

	configFlags.AddFlags(rootCmd.PersistentFlags())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error executing command: %v\n", err)
		os.Exit(1)
	}
}

func getPods(configFlags *genericclioptions.ConfigFlags) error {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to build config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create clientset: %w", err)
	}

	namespace := *configFlags.Namespace
	if namespace == "" {
		namespace = "default"
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list pods: %w", err)
	}

	fmt.Printf("Pods in namespace '%s':\n", namespace)
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}

	return nil
}
