package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	configFlags := genericclioptions.NewConfigFlags(true)
	var outputFormat string

	rootCmd := &cobra.Command{
		Use:   "pod-fetcher",
		Short: "Fetch pods from a Kubernetes cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getPods(configFlags, outputFormat)
		},
	}

	configFlags.AddFlags(rootCmd.PersistentFlags())
	rootCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format (table, json, yaml, name)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error executing command: %v\n", err)
		os.Exit(1)
	}
}

func getPods(configFlags *genericclioptions.ConfigFlags, outputFormat string) error {
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

	return printPods(pods, outputFormat)
}

func printPods(pods *corev1.PodList, outputFormat string) error {
	var printer printers.ResourcePrinter

	switch outputFormat {
	case "json":
		printer = printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.JSONPrinter{})
	case "yaml":
		printer = printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
	case "table":
		printOpts := printers.PrintOptions{}
		printer = printers.NewTypeSetter(scheme.Scheme).ToPrinter(printers.NewTablePrinter(printOpts))
	case "wide":
		printOpts := printers.PrintOptions{
			WithNamespace: true,
			Wide:          true,
			ShowLabels:    true,
		}
		printer = printers.NewTypeSetter(scheme.Scheme).ToPrinter(printers.NewTablePrinter(printOpts))
	default:
		return fmt.Errorf("unknown output format: %s", outputFormat)
	}

	// Print the pod list using the chosen printer
	if err := printer.PrintObj(pods, os.Stdout); err != nil {
		return fmt.Errorf("failed to print pods: %w", err)
	}

	return nil
}
