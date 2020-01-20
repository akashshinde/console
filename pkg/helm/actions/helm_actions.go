package actions

import "helm.sh/helm/v3/pkg/action"

type HelmActions struct {
	conf *action.Configuration
}

func NewHelmActions(conf *action.Configuration) *HelmActions {
	return &HelmActions{conf: conf}
}
