package kubernetes

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Database struct {
	Name string
	Uuid string
}

type Column struct {
	Role         string
	Name         string
	Comment      string
	AllowNull    string
	Type         string
	DefaultValue string
	FieldName    string
	PrimaryKey   bool
	Unique       bool
}

type Model struct {
	Name    string
	Role    string
	Columns []Column
}

func int32Ptr(i int32) *int32 { return &i }

func CreateDatabase(database *Database) *error {
	deploymentClient, deployment := getDeploymentClient(), &appsv1.Deployment{
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
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						apiv1.Container{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								apiv1.ContainerPort{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	return nil
}

func GetDatabases() ([]Database, *error) {
	return nil, nil
}

func UpdateDatabases(Database) *error {
	return nil
}

func DeleteDatabases(Database) *error {
	return nil
}

func CreateModel() *error {
	return nil
}

func GetModels() *error {
	return nil
}

func UpdateModels() *error {
	return nil
}

func DeleteModels() *error {
	return nil
}

func QueryDatabase() *error {
	return nil
}
