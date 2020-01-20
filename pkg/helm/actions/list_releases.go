package actions

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func (a *HelmActions) ListReleases() ([]*release.Release, error) {
	cmd := action.NewList(a.conf)

	releases, err := cmd.Run()
	if err != nil {
		return nil, err
	}
	if releases == nil {
		var rs []*release.Release
		return rs, nil
	}
	return releases, nil
}
