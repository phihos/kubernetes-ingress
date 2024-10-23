package k8s

import (
	"time"

	"github.com/jrhouston/k8slock"
	k8sclientset "k8s.io/client-go/kubernetes"
)

type Lease struct {
	name      string
	namespace string
	clientID  string
	ttl       time.Duration
	clientset *k8sclientset.Clientset
}

func NewLease(name string, namespace string, ttl time.Duration, clientID string, clientset *k8sclientset.Clientset) *Lease {
	return &Lease{name, namespace, clientID, ttl, clientset}
}

func (l Lease) Lock() error {
	locker, err := k8slock.NewLocker(
		l.name,
		k8slock.Namespace(l.namespace),
		k8slock.Clientset(l.clientset),
		k8slock.TTL(l.ttl),
		k8slock.ClientID(l.clientID),
	)
	if err != nil {
		return err
	}
	locker.Lock()
	return nil
}

func (l Lease) Unlock() error {
	locker, err := k8slock.NewLocker(
		l.name,
		k8slock.Namespace(l.namespace),
		k8slock.Clientset(l.clientset),
		k8slock.TTL(l.ttl),
		k8slock.ClientID(l.clientID),
	)
	if err != nil {
		return err
	}
	locker.Unlock()
	return nil
}
