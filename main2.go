package main

import (
	"fmt"
	"reflect"
	_ "time"
)
import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									Protocol:      v1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	Analize(deployment, 0)
	//iterateFields(deployment, 0)

}

func Analize(x interface{}, nivel int) {
	rType := reflect.TypeOf(x)
	fKind := rType.Kind()
	s := reflect.Indirect(reflect.ValueOf(x))

	if fKind == reflect.Struct {
		fmt.Printf("Struct: %s\n", rType.Name())
		for i := 0; i < s.NumField(); i++ {
			spacios(nivel + 4)
			fmt.Printf(" %s %s ", rType.Field(i).Name, s.Field(i).String())
			if rType.Field(i).Name == "Time" {
				return
			}
			Analize(s.Field(i).Interface(), nivel+8)
		}
	} else {
		fmt.Printf("%s\n", fKind)
	}
}

func int32Ptr(i int32) *int32 { return &i }

func spacios(i int) {
	for a := 1; a <= i; a++ {
		fmt.Printf(" ")
	}

}

func iterateFields(v interface{}, nivel int) {
	valueOf := reflect.ValueOf(v)
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		if field.Kind() == reflect.Struct {
			iterateFields(field.Interface(), nivel+4)

		}
		for a := 1; a < nivel; a++ {
			fmt.Printf(" ")
		}
		fmt.Println(field.Type().Name(), field.String())
	}
}
