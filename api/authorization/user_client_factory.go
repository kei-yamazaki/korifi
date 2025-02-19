package authorization

import (
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	k8sclient "k8s.io/client-go/kubernetes"

	apierrors "code.cloudfoundry.org/korifi/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type UserK8sClientFactory interface {
	BuildClient(Info) (client.WithWatch, error)
	BuildK8sClient(info Info) (k8sclient.Interface, error)
}

type UnprivilegedClientFactory struct {
	config  *rest.Config
	mapper  meta.RESTMapper
	backoff wait.Backoff
}

func NewUnprivilegedClientFactory(config *rest.Config, mapper meta.RESTMapper, backoff wait.Backoff) UnprivilegedClientFactory {
	return UnprivilegedClientFactory{
		config:  rest.AnonymousClientConfig(rest.CopyConfig(config)),
		mapper:  mapper,
		backoff: backoff,
	}
}

func (f UnprivilegedClientFactory) BuildClient(authInfo Info) (client.WithWatch, error) {
	config := rest.CopyConfig(f.config)

	switch strings.ToLower(authInfo.Scheme()) {
	case BearerScheme:
		config.BearerToken = authInfo.Token

	case CertScheme:
		certBlock, rst := pem.Decode(authInfo.CertData)
		if certBlock == nil {
			return nil, fmt.Errorf("failed to decode cert PEM")
		}

		keyBlock, _ := pem.Decode(rst)
		if keyBlock == nil {
			return nil, fmt.Errorf("failed to decode key PEM")
		}

		config.CertData = pem.EncodeToMemory(certBlock)
		config.KeyData = pem.EncodeToMemory(keyBlock)

	default:
		return nil, apierrors.NewNotAuthenticatedError(errors.New("unsupported Authorization header scheme"))
	}

	userClient, err := client.NewWithWatch(config, client.Options{
		Scheme: scheme.Scheme,
		Mapper: f.mapper,
	})
	if err != nil {
		return nil, apierrors.FromK8sError(err, "")
	}

	return NewAuthRetryingClient(userClient, f.backoff), nil
}

func (f UnprivilegedClientFactory) BuildK8sClient(authInfo Info) (k8sclient.Interface, error) {
	config := rest.CopyConfig(f.config)

	switch strings.ToLower(authInfo.Scheme()) {
	case BearerScheme:
		config.BearerToken = authInfo.Token

	case CertScheme:
		certBlock, rst := pem.Decode(authInfo.CertData)
		if certBlock == nil {
			return nil, fmt.Errorf("failed to decode cert PEM")
		}

		keyBlock, _ := pem.Decode(rst)
		if keyBlock == nil {
			return nil, fmt.Errorf("failed to decode key PEM")
		}

		config.CertData = pem.EncodeToMemory(certBlock)
		config.KeyData = pem.EncodeToMemory(keyBlock)

	default:
		return nil, apierrors.NewNotAuthenticatedError(errors.New("unsupported Authorization header scheme"))
	}

	userK8sClient, err := k8sclient.NewForConfig(config)
	if err != nil {
		return nil, apierrors.FromK8sError(err, "")
	}

	return userK8sClient, nil
}
