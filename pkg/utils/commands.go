package utils

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/crds"
	log "github.com/sirupsen/logrus"
	corev1api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/kubectl/pkg/scheme"
)

const defaultTimeout = 30 * time.Second

// PodCommandExecutor ...
type PodCommandExecutor interface {
	ExecutePodCommand(string, []string) error
	ApplyPreCommands(crds.Rule) error
	ApplyPostCommands(crds.Rule) error
}

// PodCommands ...
type PodCommands struct {
	namespace             string
	name                  string
	container             string
	client                *kubernetes.Clientset
	restConfig            *rest.Config
	streamExecutorFactory streamExecutorFactory
}

// NewPodCommands ...
func NewPodCommands(podNamespace, podName, podContainer string, clientset *kubernetes.Clientset, config *rest.Config) PodCommandExecutor {
	return &PodCommands{
		namespace:             podNamespace,
		name:                  podName,
		container:             podContainer,
		client:                clientset,
		restConfig:            config,
		streamExecutorFactory: &defaultStreamExecutorFactory{},
	}
}

func (c *PodCommands) validate() error {
	if c.name == "" {
		return fmt.Errorf("validate:: exec command validate commands [%+v] name is nil", c)
	}
	if c.namespace == "" {
		return fmt.Errorf("validate:: exec command validate commands [%+v] namespace is nil", c)
	}
	return nil
}

// ApplyPreCommands ...
func (c *PodCommands) ApplyPreCommands(rule crds.Rule) error {
	log.Infof("ApplyPreCommands:: start to exec pre commands")
	c.container = rule.Spec.ContainerName
	commands := []string{}
	for _, rule := range rule.Spec.Rules {
		if rule.Type == "command" {
			commands = append(commands, rule.Value)
		} else {
			err := fmt.Errorf("ApplyPreCommands:: rule type [%s] not supportted", rule.Type)
			log.Error(err)
			return err
		}
	}
	for _, command := range commands {
		err := c.ExecutePodCommand(rule.Name, []string{"sh", "-c", command})
		if err != nil {
			log.Error(err)
			return err
		}
	}
	// err := c.ExecutePodCommand("FlushCache", []string{"sh", "-c", "fsync"})
	// if err != nil {
	// 	return err
	// }
	return nil
}

// ApplyPostCommands ...
func (c *PodCommands) ApplyPostCommands(rule crds.Rule) error {
	log.Infof("ApplyPreCommands:: start to exec post commands")
	c.container = rule.Spec.ContainerName
	commands := []string{}
	for _, rule := range rule.Spec.Rules {
		if rule.Type == "command" {
			commands = append(commands, rule.Value)
		} else {
			return fmt.Errorf("ApplyPostCommands:: rule type [%s] not supportted", rule.Type)
		}
	}
	for _, command := range commands {
		err := c.ExecutePodCommand(rule.Name, []string{"sh", "-c", command})
		if err != nil {
			return err
		}
	}
	return nil
}

// ExecutePodCommand ...
func (c *PodCommands) ExecutePodCommand(hookName string, command []string) error {
	if err := c.validate(); err != nil {
		return err
	}
	if len(command) == 0 {
		return fmt.Errorf("ExecutePodCommand:: command is required")
	}
	pod, err := c.getContainers()
	if err != nil {
		return err
	}
	if c.container == "" {
		c.setDefaultHookContainer(pod)
	}
	err = c.ensureContainerExists(pod)
	if err != nil {
		return err
	}

	hookLog := log.WithFields(
		log.Fields{
			"hookName":      hookName,
			"hookContainer": c.container,
			"hookCommand":   command,
			"hookOnError":   "Fail",
			"hookTimeout":   defaultTimeout,
		},
	)
	hookLog.Info("ExecutePodCommand:: running exec hook")

	req := c.client.CoreV1().RESTClient().Post().Resource("pods").Name(c.name).
		Namespace(c.namespace).SubResource("exec")

	req.VersionedParams(
		&corev1api.PodExecOptions{
			Command: command,
			Stdout:  true,
			Stderr:  true,
		}, scheme.ParameterCodec,
	)
	exec, err := c.streamExecutorFactory.NewSPDYExecutor(c.restConfig, "POST", req.URL())
	if err != nil {
		return err
	}
	var stdout, stderr bytes.Buffer
	errChan := make(chan error)
	go func() {
		err = exec.Stream(remotecommand.StreamOptions{
			Stdout: &stdout,
			Stderr: &stderr,
		})
		errChan <- err
	}()

	var timeoutCh <-chan time.Time
	timer := time.NewTimer(defaultTimeout)
	defer timer.Stop()
	timeoutCh = timer.C

	select {
	case err = <-errChan:
	case <-timeoutCh:
		return fmt.Errorf("ExecutePodCommand:: timeout after: %v", defaultTimeout)
	}
	log.Infof("ExecutePodCommand:: stdout: %v", stdout.String())
	log.Infof("ExecutePodCommand:: stderr: %v", stderr.String())

	if err != nil {
		return err
	}
	return nil
}

func (c *PodCommands) getContainers() (*corev1api.Pod, error) {
	pod, err := c.client.CoreV1().Pods(c.namespace).Get(context.Background(), c.name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("getContainers:: failed to get pods: %v", err)
	}
	return pod, nil
}

func (c *PodCommands) ensureContainerExists(pod *corev1api.Pod) error {
	if len(pod.Spec.Containers) > 0 {
		for _, container := range pod.Spec.Containers {
			if container.Name == c.container {
				return nil
			}
		}
	} else {
		return fmt.Errorf("ensureContainerExists:: pod containers count is zero: %v", pod.Spec.Containers)
	}
	return fmt.Errorf("ensureContainerExists:: container name [%s] not in [%s] pod", c.container, c.name)
}

func (c *PodCommands) setDefaultHookContainer(pod *corev1api.Pod) error {
	if len(pod.Spec.Containers) < 1 {
		return fmt.Errorf("need at least 1 container")
	}

	c.container = pod.Spec.Containers[0].Name

	return nil
}

type streamExecutorFactory interface {
	NewSPDYExecutor(config *rest.Config, method string, url *url.URL) (remotecommand.Executor, error)
}

type defaultStreamExecutorFactory struct{}

// NewSPDYExecutor ...
func (f *defaultStreamExecutorFactory) NewSPDYExecutor(config *rest.Config, method string, url *url.URL) (remotecommand.Executor, error) {
	return remotecommand.NewSPDYExecutor(config, method, url)
}
