package actions

import (
	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chartutil"
	kubefake "helm.sh/helm/v3/pkg/kube/fake"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"helm.sh/helm/v3/pkg/time"
	"io/ioutil"
	"testing"
)

func TestListReleases(t *testing.T) {
	store := storage.Init(driver.NewMemory())
	err := store.Create(&release.Release{
		Name:      "test",
		Namespace: "test-namespace",
		Info: &release.Info{
			FirstDeployed: time.Time{},
			Status:        "deployed",
		},
	})
	actionConfig := &action.Configuration{
		Releases:     store,
		KubeClient:   &kubefake.PrintingKubeClient{Out: ioutil.Discard},
		Capabilities: chartutil.DefaultCapabilities,
		Log:          func(format string, v ...interface{}) {},
	}
	action := NewHelmActions(actionConfig)
	rels, err := action.ListReleases()
	if err != nil {
		t.Error(err.Error())
	}
	assert.Len(t, rels, 1, "Release list should return 1 release")
	assert.Equal(t, rels[0].Name, "test", "Release name isn't matching")
	assert.Equal(t, rels[0].Namespace, "test-namespace", "Namespace name isn't matching")
	assert.Equal(t, rels[0].Info.Status, release.StatusDeployed, "Chart status is not deployed")
}
