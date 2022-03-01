
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;

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