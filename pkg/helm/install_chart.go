package helm

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func InstallChart(ns, name, url string, conf *action.Configuration) (interface{}, error) {
	cmd := action.NewInstall(conf)

	name, chart, err := cmd.NameAndChart([]string{name, url})
	if err != nil {
		return nil, err
	}
	cmd.ReleaseName = name

	cp, err := cmd.ChartPathOptions.LocateChart(chart, settings)
	if err != nil {
		return nil, err
	}

	ch, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}

	cmd.Namespace = ns
	release, err := cmd.Run(ch, nil)
	if err != nil {
		return nil, err
	}
	return release, nil
}
