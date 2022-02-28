
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;

namespace DatabaseControllerKubeOps.Controller.Controllers;

[EntityRbac(typeof(V1Alpha1DatabaseEntity), Verbs = RbacVerb.All)]
public class DatabaseTableCtrl : IResourceController<V1Alpha1DatabaseTableEntity>
{

    public Task<ResourceControllerResult> CreatedAsync(V1Alpha1DatabaseEntity resource)
    {
        //return Task.FromResult<ResourceControllerResult>(null); // This wont trigger a requeue.
        return Task.FromResult(ResourceControllerResult.RequeueEvent(TimeSpan.FromSeconds(15))); // This will requeue the event in 15 seconds.
    }

    public Task<ResourceControllerResult> ReconcileAsync(V1Alpha1DatabaseEntity resource)
    {
        //return Task.FromResult<ResourceControllerResult>(null); // This wont trigger a requeue.
        return Task.FromResult(ResourceControllerResult.RequeueEvent(TimeSpan.FromSeconds(15))); // This will requeue the event in 15 seconds.
    }

    public Task<ResourceControllerResult> StatusModifiedAsync(V1Alpha1DatabaseEntity resource)
    {
        //return Task.FromResult<ResourceControllerResult>(null); // This wont trigger a requeue.
        return Task.FromResult(ResourceControllerResult.RequeueEvent(TimeSpan.FromSeconds(15))); // This will requeue the event in 15 seconds.
    }

    public Task<ResourceControllerResult> DeletedAsync(V1Alpha1DatabaseEntity resource)
    {
        //return Task.FromResult<ResourceControllerResult>(null); // This wont trigger a requeue.
        return Task.FromResult(ResourceControllerResult.RequeueEvent(TimeSpan.FromSeconds(15))); // This will requeue the event in 15 seconds.
    }


}
