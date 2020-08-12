package chartproxy

import (
	"strconv"
	"testing"

	"github.com/openshift/console/pkg/helm/actions/fake"
	"github.com/openshift/console/pkg/helm/testdata"
)

func TestHelmRepoGetter_List(t *testing.T) {
	tests := []struct {
		testName        string
		indexFile       []string
		expectedEntries int
		defaultRepoName string
	}{
		{
			testName:        "correct index.yaml should report correct no of chart entries - 2 index file",
			indexFile:       []string{testdata.AzureRepoYaml, sampleRepoYaml},
			expectedEntries: 2,
		},
		{
			testName:        "correct index.yaml should report correct no of chart entries - 1 index file",
			indexFile:       []string{testdata.AzureRepoYaml},
			expectedEntries: 1,
		},
		{
			testName:        "default repo index file should be served in case no chart repo configured ",
			indexFile:       []string{},
			expectedEntries: 1,
			defaultRepoName: "redhat-helm-charts",
		},
	}

	for _, tt := range tests {
		client := fake.SetupClusterWithHelmCrs(tt.indexFile)
		cfg := config{repoUrl: "https://default-url.com"}
		cfg.Configure()
		repoGetter := NewRepoGetter(client, nil)
		repos, err := repoGetter.List()
		if err != nil {
			t.Error(err)
		}

		if len(repos) != tt.expectedEntries {
			t.Errorf("index length mismatch expected is %d received %d", tt.expectedEntries, len(repos))
		}
		for i, repo := range repos {
			var repoName string
			if tt.defaultRepoName != "" {
				repoName = tt.defaultRepoName
			} else {
				repoName = "sample-repo-" + strconv.Itoa(i+1)
			}
			if repo.Name != repoName {
				t.Errorf("Repo name mismatch expected is %s received %s", repoName, repo.Name)
			}
		}
	}
}
