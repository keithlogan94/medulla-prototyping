
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;
using k8s;
using k8s.Models;
using System.Text.Json;
using System.Threading;

namespace DatabaseControllerKubeOps.Controller.Controllers;


internal class OnChange
{

    private static Random random = new Random();

    public static string RandomString(int length = 5)
    {
        string chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789".ToLower();
        return new string(Enumerable.Repeat(chars, length)
            .Select(s => s[random.Next(s.Length)]).ToArray());
    }

    public static async void UpdateDatabase(V1Alpha1DataEntity entity)
    {
        var unique = OnChange.RandomString();
        var config = KubernetesClientConfiguration.InClusterConfig();
        var client = new Kubernetes(config);

        var pod = new V1Pod
        {
            Metadata = new V1ObjectMeta
            {
                Name = "database-sync-" + unique,
                NamespaceProperty = "default",
                Labels = new Dictionary<string, string>
                {
                    { "app", "database-sync-" + unique }
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
                Name = "database-sync-service-" + unique,
            },
            Spec = new V1ServiceSpec
            {
                Type = "LoadBalancer",
                Selector = new Dictionary<string, string>
                {
                    ["app"] = "database-sync-" + unique,
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

        var values = new Dictionary<string, string>
        {
            { "data", JsonSerializer.Serialize(entity) },
        };
        var content = new FormUrlEncodedContent(values);

        HttpClient httpClient = new HttpClient();

        while (true)
        {
            try
            {
                var response = await httpClient.PostAsync("http://database-sync-service-" + unique + "/listen-for-database-schema", content);
                var responseString = await response.Content.ReadAsStringAsync();
                Console.WriteLine(responseString);
                break;
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.ToString());
                Console.WriteLine("Sleeping for 5 seconds until trying again");
                Thread.Sleep(5000);
            }
        }
        
    }
}


[EntityRbac(typeof(V1Alpha1DataEntity), Verbs = RbacVerb.All)]
public class DataCtrl : IResourceController<V1Alpha1DataEntity>
{
    
    public Task<ResourceControllerResult> CreatedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("Created");
        OnChange.UpdateDatabase(resource);
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> ReconcileAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("ReconcileAsync");
        OnChange.UpdateDatabase(resource);
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> StatusModifiedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("StatusModifiedAsync");
        OnChange.UpdateDatabase(resource);
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> DeletedAsync(V1Alpha1DataEntity resource)
    {
        Console.WriteLine("DeletedAsync");
        return Task.FromResult<ResourceControllerResult>(null);
    }

}