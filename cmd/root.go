package cmd

import (
	"fmt"
	"github.com/jasonqiann/kube-plugin/pkg"
	"os"
)

func Execute() {

	config, err := pkg.ParseFlag()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err to parse flag, %v", err)
		os.Exit(1)
	}

	podDetailCommand := pkg.NewPodDetail(config)
	if err := podDetailCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
