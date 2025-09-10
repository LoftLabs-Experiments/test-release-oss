package framework

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/loft-sh/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	PollInterval            = 5 * time.Second
	PollTimeout             = time.Minute
	PollTimeoutLong         = 2 * time.Minute
	DefaultTestAppName      = "testapp"
	DefaultTestAppNamespace = "testapp"
	DefaultClientTimeout    = 100 * time.Second // the default in client-go is 32 (which is not enough, as we are occasionally experiencing client-side throttling in e2e tests)
	DefaultClientBurst      = 100               // the default in client-go is 10 (which is not enough, as we are occasionally experiencing client-side throttling in e2e tests)
	DefaultClientQPS        = 50                // the default in client-go is 5 (which is not enough, as we are occasionally experiencing client-side throttling in e2e tests)
)

var DefaultFramework = &Framework{}

type Framework struct {
	// The context to use for testing
	Context context.Context

	// TestAppName is the name of the test app instance which we are testing
	TestAppName string

	// TestAppNamespace is the namespace in cluster of the current
	// test app instance which we are testing
	TestAppNamespace string

	// The suffix to append to the synced resources in the namespace
	Suffix string

	// Config is the kubernetes rest config of the
	// kubernetes cluster were we are testing in
	Config *rest.Config

	// Client is the kubernetes client of the current
	// kubernetes cluster were we are testing in
	Client *kubernetes.Clientset

	// CRClient is the controller runtime client of the current
	// kubernetes cluster were we are testing in
	CRClient client.Client

	// KubeConfigFile is a file containing kube config
	// of the current kubernetes cluster which we are testing.
	// This file shall be deleted in the end of the test suite execution.
	KubeConfigFile *os.File

	// Log is the logger that should be used
	Log log.Logger

	// ClientTimeout value used in the clients
	ClientTimeout time.Duration

	// ClientBurst value used in the clients
	ClientBurst int

	// ClientQPS value used in the clients
	ClientQPS float32
}

func CreateFramework(ctx context.Context) error {
	// setup loggers
	l := log.GetInstance()

	name := os.Getenv("TESTAPP_NAME")
	if name == "" {
		name = DefaultTestAppName
	}
	ns := os.Getenv("TESTAPP_NAMESPACE")
	if ns == "" {
		ns = DefaultTestAppNamespace
	}
	timeoutEnvVar := os.Getenv("TESTAPP_CLIENT_TIMEOUT")
	var timeout time.Duration
	timeoutInt, err := strconv.Atoi(timeoutEnvVar)
	if err == nil {
		timeout = time.Duration(timeoutInt) * time.Second
	} else {
		timeout = DefaultClientTimeout
	}

	clientBurstEnvVar := os.Getenv("TESTAPP_CLIENT_BURST")
	var clientBurst int
	clientBurst, err = strconv.Atoi(clientBurstEnvVar)
	if err != nil {
		clientBurst = DefaultClientBurst
	}

	clientQPSEnvVar := os.Getenv("TESTAPP_CLIENT_QPS")
	var clientQPS int
	clientQPS, err = strconv.Atoi(clientQPSEnvVar)
	if err != nil {
		clientQPS = DefaultClientQPS
	}

	suffix := os.Getenv("TESTAPP_SUFFIX")
	if suffix == "" {
		suffix = "testapp"
	}

	l.Infof("Testing TestApp named: %s in namespace: %s", name, ns)
	config, err := ctrl.GetConfig()
	if err != nil {
		return err
	}
	config.Timeout = timeout
	config.Burst = clientBurst
	config.QPS = float32(clientQPS)

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	crClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	// create the framework
	DefaultFramework = &Framework{
		Context:          ctx,
		TestAppName:      name,
		TestAppNamespace: ns,
		Suffix:           suffix,
		Config:           config,
		Client:           kubeClient,
		CRClient:         crClient,
		Log:              l,
		ClientTimeout:    timeout,
		ClientBurst:      clientBurst,
		ClientQPS:        float32(clientQPS),
	}

	l.Done("Framework successfully initialized")
	return nil
}

func (f *Framework) RefreshClient() error {
	// Simplified version - just use current kubeconfig
	config, err := ctrl.GetConfig()
	if err != nil {
		return fmt.Errorf("could not get kubeconfig: %w", err)
	}
	config.Timeout = f.ClientTimeout
	config.Burst = f.ClientBurst
	config.QPS = f.ClientQPS

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	crClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return err
	}

	// Test connectivity
	_, err = kubeClient.CoreV1().ServiceAccounts("default").Get(f.Context, "default", metav1.GetOptions{})
	if err != nil {
		return err
	}

	f.Config = config
	f.Client = kubeClient
	f.CRClient = crClient
	return nil
}

func (f *Framework) Cleanup() error {
	if f.KubeConfigFile != nil {
		return os.Remove(f.KubeConfigFile.Name())
	}
	return nil
}
