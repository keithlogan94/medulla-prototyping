using k8s.Models;
using KubeOps.Operator.Entities;
using KubeOps.Operator.Entities.Annotations;

namespace DatabaseControllerKubeOps.Controller.Entities;


[Description("A CustomResourceDefinition which allows building a database in Medulla.")]
[KubernetesEntity(
    ApiVersion = "v1alpha1",
    Kind = "Database",
    Group = "medulla.recro.com",
    PluralName = "databases")]
public class V1Alpha1DatabaseEntity : CustomKubernetesEntity
{
    [Required]
    public string? Name { get; set; }
}
