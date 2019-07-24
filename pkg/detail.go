package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"os"
)

type PodDetailOptions struct {
	client    *kubernetes.Clientset
	namespace string
}

func NewPodDetailOptions(c *Configuration) (*PodDetailOptions, error) {

	client, err := initKubeClient(c)
	if err != nil {
		return nil, err
	}
	p := &PodDetailOptions{
		client:    client,
		namespace: c.namespace,
	}
	return p, nil
}

func (p *PodDetailOptions) Run() {
	var ns string
	if p.namespace == "" {
		ns = v1.NamespaceAll
	} else {
		ns = p.namespace
	}

	podList, err := p.client.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		fmt.Errorf("err go get pod list, %v", err)
	}

	fmt.Printf("namespace:%s has %d pod\n", ns, len(podList.Items))

}

func initKubeClient(config *Configuration) (*kubernetes.Clientset, error) {
	var cfg *rest.Config
	var err error
	if config.configFile == "" {
		klog.Infof("no --kubeconfig, use in-cluster kubernetes config")
		cfg, err = rest.InClusterConfig()
		if err != nil {
			klog.Errorf("use in cluster config failed %v", err)
			return nil, err
		}
	} else {
		cfg, err = clientcmd.BuildConfigFromFlags("", config.configFile)
		if err != nil {
			klog.Errorf("use --kubeconfig %s failed %v", config.configFile, err)
			return nil, err
		}
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Errorf("init kubernetes client failed %v", err)
		return nil, err
	}

	return kubeClient, nil
}

func NewPodDetail(c *Configuration) *cobra.Command {

	podDetailOptions, err := NewPodDetailOptions(c)
	if err != nil {
		fmt.Errorf("err to new pod detail server, err", err)
		os.Exit(1)
	}

	podDetail := &cobra.Command{
		Use:   "pod detail",
		Short: "show pod detail in k8s cluster",
		Run: func(cmd *cobra.Command, args []string) {
			podDetailOptions.Run()
		},
	}

	return podDetail
}
