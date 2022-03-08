package uo

import (
	"github.com/codablock/kluctl/pkg/types/k8s"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (uo *UnstructuredObject) GetK8sGVK() schema.GroupVersionKind {
	kind, _, err := uo.GetNestedString("kind")
	if err != nil {
		log.Fatal(err)
	}
	apiVersion, _, err := uo.GetNestedString("apiVersion")
	if err != nil {
		log.Fatal(err)
	}
	gv, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		log.Fatal(err)
	}
	return schema.GroupVersionKind{
		Group:   gv.Group,
		Version: gv.Version,
		Kind:    kind,
	}
}

func (uo *UnstructuredObject) SetK8sGVK(gvk schema.GroupVersionKind) {
	err := uo.SetNestedField(gvk.GroupVersion().String(), "apiVersion")
	if err != nil {
		log.Fatal(err)
	}
	err = uo.SetNestedField(gvk.Kind, "kind")
	if err != nil {
		log.Fatal(err)
	}
}

func (uo *UnstructuredObject) SetK8sGVKs(g string, v string, k string) {
	uo.SetK8sGVK(schema.GroupVersionKind{Group: g, Version: v, Kind: k})
}

func (uo *UnstructuredObject) GetK8sName() string {
	s, _, err := uo.GetNestedString("metadata", "name")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (uo *UnstructuredObject) SetK8sName(name string) {
	err := uo.SetNestedField(name, "metadata", "name")
	if err != nil {
		log.Fatal(err)
	}
}

func (uo *UnstructuredObject) GetK8sNamespace() string {
	s, _, err := uo.GetNestedString("metadata", "namespace")
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (uo *UnstructuredObject) SetK8sNamespace(namespace string) {
	if namespace != "" {
		err := uo.SetNestedField(namespace, "metadata", "namespace")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := uo.RemoveNestedField("metadata", "namespace")
		if err != nil {
			log.Fatal(err)
		}
	}

}

func (uo *UnstructuredObject) GetK8sRef() k8s.ObjectRef {
	return k8s.ObjectRef{
		GVK:       uo.GetK8sGVK(),
		Name:      uo.GetK8sName(),
		Namespace: uo.GetK8sNamespace(),
	}
}

func (uo *UnstructuredObject) GetK8sLabels() map[string]string {
	ret, ok, err := uo.GetNestedStringMapCopy("metadata", "labels")
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		return map[string]string{}
	}
	return ret
}

func (uo *UnstructuredObject) SetK8sLabels(labels map[string]string) {
	_ = uo.RemoveNestedField("metadata", "labels")
	for k, v := range labels {
		uo.SetK8sLabel(k, v)
	}
}

func (uo *UnstructuredObject) GetK8sLabel(name string) *string {
	ret, ok, err := uo.GetNestedString("metadata", "labels", name)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		return nil
	}
	return &ret
}

func (uo *UnstructuredObject) SetK8sLabel(name string, value string) {
	err := uo.SetNestedField(value, "metadata", "labels", name)
	if err != nil {
		log.Fatal(err)
	}
}

func (uo *UnstructuredObject) GetK8sAnnotations() map[string]string {
	ret, ok, err := uo.GetNestedStringMapCopy("metadata", "annotations")
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		return map[string]string{}
	}
	return ret
}

func (uo *UnstructuredObject) GetK8sAnnotation(name string) *string {
	ret, ok, err := uo.GetNestedString("metadata", "annotations", name)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		return nil
	}
	return &ret
}

func (uo *UnstructuredObject) SetK8sAnnotations(annotations map[string]string) {
	_ = uo.RemoveNestedField("metadata", "annotations")
	for k, v := range annotations {
		uo.SetK8sAnnotation(k, v)
	}
}

func (uo *UnstructuredObject) SetK8sAnnotation(name string, value string) {
	err := uo.SetNestedField(value, "metadata", "annotations", name)
	if err != nil {
		log.Fatal(err)
	}
}

func (uo *UnstructuredObject) GetK8sResourceVersion() string {
	ret, _, _ := uo.GetNestedString("metadata", "resourceVersion")
	return ret
}

func (uo *UnstructuredObject) SetK8sResourceVersion(rv string) {
	if rv == "" {
		_ = uo.RemoveNestedField("metadata", "resourceVersion")
	} else {
		err := uo.SetNestedField(rv, "metadata", "resourceVersion")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (uo *UnstructuredObject) GetK8sManagedFields() []metav1.ManagedFieldsEntry {
	return uo.ToUnstructured().GetManagedFields()
}
