package helm

type HelmRequest struct {
	Name      string
	Namespace string
	Url       string
	Values    map[string]interface{}
}
