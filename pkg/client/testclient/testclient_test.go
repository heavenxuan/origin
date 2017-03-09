package testclient

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	kapi "k8s.io/kubernetes/pkg/api"

	deployapi "github.com/openshift/origin/pkg/deploy/api"
)

func TestNewClient(t *testing.T) {
	o, err := ReadObjectsFromPath("../../../test/integration/testdata/test-deployment-config.yaml", "test", kapi.Codecs.UniversalDecoder(), kapi.Scheme)
	if err != nil {
		t.Fatal(err)
	}
	oc, _ := NewFixtureClients(o...)
	list, err := oc.DeploymentConfigs("test").List(metainternal.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}

	// same result
	list, err = oc.DeploymentConfigs("test").List(metainternal.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected list %#v", list)
	}
	t.Logf("list: %#v", list)
}

func TestErrors(t *testing.T) {
	oc, _ := NewErrorClients(errors.NewNotFound(deployapi.Resource("DeploymentConfigList"), ""))
	_, err := oc.DeploymentConfigs("test").List(metainternal.ListOptions{})
	if !errors.IsNotFound(err) {
		t.Fatalf("unexpected error: %v", err)
	}

	oc, _ = NewErrorClients(errors.NewForbidden(deployapi.Resource("DeploymentConfigList"), "", nil))
	_, err = oc.DeploymentConfigs("test").List(metainternal.ListOptions{})
	if !errors.IsForbidden(err) {
		t.Fatalf("unexpected error: %v", err)
	}
}
