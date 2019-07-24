package pkg

import "github.com/spf13/pflag"

type Configuration struct {
	configFile string
	namespace  string
}

func ParseFlag() (*Configuration, error) {

	var (
		argKubeConfig = pflag.String("--kubeconfig", "/root/.kube/config", "kube config"+
			"file path to contact with k8s cluster")
		argNamespace = pflag.String("--namespace", "", "namespace which to use")
	)
	pflag.Parse()

	return &Configuration{
		configFile: *argKubeConfig,
		namespace:  *argNamespace,
	}, nil
}
