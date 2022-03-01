
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;
using k8s;
using k8s.Models;

namespace DatabaseControllerKubeOps.Controller.Controllers;


internal class OnChange
{
    public static async void UpdateDatabase()
    {
        var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();
        var client = new Kubernetes(config);

        var pod = new V1Pod
        {
            Metadata = new V1ObjectMeta
            {
                Name = "database-sync",
                NamespaceProperty = "default",
                Labels = new Dictionary<string, string>
                {
                    { "app", "database-sync" }
                }
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


        V1Service service = new V1Service()
        {
            ApiVersion = $"{V1Service.KubeGroup}/{V1Service.KubeApiVersion}",
            Kind = V1Service.KubeKind,
            Metadata = new V1ObjectMeta()
            {
                Name = "database-sync-service",
            },
            Spec = new V1ServiceSpec
            {
                Type = "LoadBalancer",
                Selector = new Dictionary<string, string>
                {
                    ["app"] = "database-sync",
                },
                Ports = new List<V1ServicePort> {
                    new V1ServicePort {
                        Protocol = "TCP",
                        Port = 3000,
                        TargetPort = 3000,
                    },
                }
            }
        };

        client.CreateNamespacedService(service, "default");

        Console.WriteLine(result);
    }
}


[EntityRbac(typeof(V1Alpha1DataEntity), Verbs = RbacVerb.All)]
public class DataCtrl : IResourceController<V1Alpha1DataEntity>
{
    
    public Task<ResourceControllerResult> CreatedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("Created");
        OnChange.UpdateDatabase();
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> ReconcileAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("ReconcileAsync");
        OnChange.UpdateDatabase();
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> StatusModifiedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("StatusModifiedAsync");
        OnChange.UpdateDatabase();
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> DeletedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("DeletedAsync");
        return Task.FromResult<ResourceControllerResult>(null);
    }

}