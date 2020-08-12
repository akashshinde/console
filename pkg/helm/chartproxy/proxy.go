package chartproxy

import (
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type proxy struct {
	config *rest.Config

	DynamicClient  dynamic.Interface
	coreV1Client   v1.CoreV1Interface
	helmRepoGetter HelmRepoGetter

	IndexFileGetter
}

type IndexFileGetter interface {
	IndexFile() (*repo.IndexFile, error)
}

type RestConfigProvider func() (*rest.Config, error)

type ProxyOption func(p *proxy) error

func dynamicKubeClientProvider(p *proxy) error {
	client, err := dynamic.NewForConfig(p.config)
	if err != nil {
		return err
	}
	p.DynamicClient = client
	return nil
}

func coreClientProvider(p *proxy) error {
	client, err := kubernetes.NewForConfig(p.config)
	if err != nil {
		return err
	}
	p.coreV1Client = client.CoreV1()
	return nil
}

var defaultOptions = []ProxyOption{dynamicKubeClientProvider, coreClientProvider}

func New(k8sConfig RestConfigProvider, opts ...ProxyOption) (IndexFileGetter, error) {
	config, err := k8sConfig()
	if err != nil {
		return nil, err
	}
	p := &proxy{
		config: config,
	}

	if len(opts) == 0 {
		opts = defaultOptions
	}

	for _, opt := range opts {
		opt(p)
	}
	p.helmRepoGetter = NewRepoGetter(p.DynamicClient, p.coreV1Client)
	return p, nil
}

func (p *proxy) IndexFile() (*repo.IndexFile, error) {
	helmRepos, err := p.helmRepoGetter.List()
	if err != nil {
		return nil, err
	}

	var indexFiles []*repo.IndexFile
	for _, helmRepo := range helmRepos {
		idxFile, err := helmRepo.IndexFile()
		if err != nil {
			return nil, err
		}
		indexFiles = append(indexFiles, idxFile)
	}
	return mergeIndexFiles(indexFiles...), nil
}
