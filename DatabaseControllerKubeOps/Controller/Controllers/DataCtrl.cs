
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;
using k8s;
using k8s.Models;

namespace DatabaseControllerKubeOps.Controller.Controllers;


[EntityRbac(typeof(V1Alpha1DataEntity), Verbs = RbacVerb.All)]
public class DataCtrl : IResourceController<V1Alpha1DataEntity>
{
    
    public Task<ResourceControllerResult> CreatedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("Created");
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> ReconcileAsync(V1Alpha1DataEntity resource)
    {
        var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();
        var client = new Kubernetes(config);

        var pod = new V1Pod
        {
            Metadata = new V1ObjectMeta
            {
                Name = "database-sync",
                NamespaceProperty = "default"
            },
            Spec = new V1PodSpec
            {
                Containers = new List<V1Container>()
                {
                    new V1Container()
                    {
                        Name = "database-sync",
                        Image = "keithlogan94/database-sync:latest"
                    }
                }
            }
        };

        var result = client.CreateNamespacedPod(pod, "default");
        Console.WriteLine(result);

        Console.WriteLine("ReconcileAsync");
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> StatusModifiedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("StatusModifiedAsync");
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> DeletedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("DeletedAsync");
        return Task.FromResult<ResourceControllerResult>(null);
    }

}