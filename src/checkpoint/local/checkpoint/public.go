package checkpoint

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	v1 "github.com/openshift/api/network/v1"
	networkv1client "github.com/openshift/client-go/network/clientset/versioned/typed/network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	k8snetclient "k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
)

const egressAnnotation string = "checkpoint.com/egress-rules"
const ingressAnnotation string = "checkpoint.com/ingress-rules"

type Ruler interface {
	Rules(SessionT) RulesMapT
}

type Ingress struct{}
type Egress struct{}

type SessionT struct {
	coreclient    corev1client.CoreV1Client
	networkclient networkv1client.NetworkV1Client
	k8snetclient  k8snetclient.NetworkingV1Client
}

type LocksT struct {
	Egress  sync.RWMutex
	Ingress sync.RWMutex
}

type RulesMapT map[string][]string

type RulesT struct {
	Egress  RulesMapT
	Ingress RulesMapT
	Locks   LocksT
}

var (
	netnamespace v1.NetNamespace
	ingress      k8snetclient.IngressClassInterface
	restconfig   *rest.Config
)

func Session() SessionT {

	restconfig := getRestConfig()

	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	networkclient, err := networkv1client.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	k8snetworkclient, err := k8snetclient.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}

	res := SessionT{
		coreclient:    *coreclient,
		networkclient: *networkclient,
		k8snetclient:  *k8snetworkclient,
	}

	return res
}

func (egress Egress) Rules(session SessionT, ctx context.Context) (RulesMap RulesMapT, err error) {
	var rulesmap RulesMapT = make(RulesMapT)

	namespaces, err := session.coreclient.Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		error := errors.New("Cannot get list of namespaces, please login or check service account!")
		return rulesmap, error
	}
	netnamespaces, err := session.networkclient.NetNamespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		error := errors.New("Cannot get list of netnamespaces, are you using openshift ?")
		return rulesmap, error
	}

	for _, namespace := range namespaces.Items {
		for _, netv := range netnamespaces.Items {
			if netv.Name == namespace.Name {
				netnamespace = netv
				break
			}
		}
		for k, v := range namespace.Annotations {
			if k == egressAnnotation {
				t := strings.ReplaceAll(strings.ReplaceAll(strings.TrimRight(strings.TrimLeft(v, "["), "]"), "\"", ""), " ", "")
				rules := egress2array(t)
				for _, e := range netnamespace.EgressIPs {
					for _, ru := range rules {
						rulesmap[ru] = append(rulesmap[ru], string(e))
					}
				}
			}
		}
	}
	return rulesmap, nil
}

func (ingress Ingress) Rules(session SessionT, ctx context.Context) (RulesMap RulesMapT, err error) {

	var (
		rulesmap        RulesMapT = make(RulesMapT)
		ingressRuleList []map[string][]string
	)

	namespaces, err := session.coreclient.Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		error := errors.New("Cannot get list of namespaces, please login or check service account!")
		return rulesmap, error
	}

	for _, namespace := range namespaces.Items {
		ruleAnnotation := ""

		for k, v := range namespace.Annotations {
			if k == ingressAnnotation {
				ruleAnnotation = v
				break
			}
		}

		if len(ruleAnnotation) == 0 {
			continue
		}

		json.Unmarshal([]byte(ruleAnnotation), &ingressRuleList)
		if ingressRuleList == nil {
			continue
		}

		ingressObjectList, _ := session.k8snetclient.Ingresses(namespace.Name).List(ctx, metav1.ListOptions{})
		if err != nil {
			error := errors.New("Cannot get list of ingress objects, please login or check service account!")
			return rulesmap, error
		}

		for _, v := range ingressRuleList {
			for in, rs := range v {
				for _, j := range ingressObjectList.Items {
					if in == j.Name {
						ipAddress := j.Status.LoadBalancer.Ingress[0].IP
						for _, ru := range rs {
							rulesmap[ru] = append(rulesmap[ru], ipAddress)
						}
					} else {
						continue
					}
				}
			}
		}
		if err != nil {
			error := errors.New("Cannot get list of ingresses, please check verbs in role for sa")
			return rulesmap, error
		}
	}

	return rulesmap, nil
}
