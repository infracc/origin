package v1

import (
	v1 "github.com/openshift/origin/pkg/build/apis/build/v1"
	"github.com/openshift/origin/pkg/build/generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type BuildV1Interface interface {
	RESTClient() rest.Interface
	BuildsGetter
	BuildConfigsGetter
}

// BuildV1Client is used to interact with features provided by the build.openshift.io group.
type BuildV1Client struct {
	restClient rest.Interface
}

func (c *BuildV1Client) Builds(namespace string) BuildResourceInterface {
	return newBuilds(c, namespace)
}

func (c *BuildV1Client) BuildConfigs(namespace string) BuildConfigInterface {
	return newBuildConfigs(c, namespace)
}

// NewForConfig creates a new BuildV1Client for the given config.
func NewForConfig(c *rest.Config) (*BuildV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BuildV1Client{client}, nil
}

// NewForConfigOrDie creates a new BuildV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BuildV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BuildV1Client for the given RESTClient.
func New(c rest.Interface) *BuildV1Client {
	return &BuildV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BuildV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
