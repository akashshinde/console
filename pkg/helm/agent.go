package helm

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"net/http"
)

var settings = cli.New()

func GetActionConfigurations(host, ns, token string, transport *http.RoundTripper) *action.Configuration {

	conf := &rest.Config{
		Host:        host,
		BearerToken: token,
		Transport:   *transport,
	}
	clientset, _ := kubernetes.NewForConfig(conf)
	tr := true
	kubeConf := &genericclioptions.ConfigFlags{
		APIServer:   &host,
		Insecure:    &tr,
		BearerToken: &token,
		Namespace: &ns,
	}
	store := createStorage(ns, clientset)
	config := &action.Configuration{
		RESTClientGetter: kubeConf,
		Releases:         store,
		KubeClient:       kube.New(kubeConf),
		RegistryClient:   nil,
		Capabilities:     nil,
		Log:              klog.Infof,
	}
	return config
}

func createStorage(namespace string, clientset *kubernetes.Clientset) *storage.Storage {
	var store *storage.Storage
	d := driver.NewSecrets(clientset.CoreV1().Secrets(namespace))
	d.Log = klog.Infof
	store = storage.Init(d)
	return store
}
