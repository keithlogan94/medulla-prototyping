﻿
using KubeOps.Operator.Controller;
using KubeOps.Operator.Rbac;
using DatabaseControllerKubeOps.Controller.Entities;
using KubeOps.Operator.Controller.Results;

namespace DatabaseControllerKubeOps.Controller.Controllers;


[EntityRbac(typeof(V1Alpha1DatabaseEntity), Verbs = RbacVerb.All)]
public class DatabaseCtrl : IResourceController<V1Alpha1DatabaseEntity>
{

    public Task<ResourceControllerResult> CreatedAsync(V1Alpha1DatabaseEntity resource)
    {
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> ReconcileAsync(V1Alpha1DatabaseEntity resource)
    {
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> StatusModifiedAsync(V1Alpha1DatabaseEntity resource)
    {
        return Task.FromResult<ResourceControllerResult>(null);
    }

    public Task<ResourceControllerResult> DeletedAsync(V1Alpha1DatabaseEntity resource)
    {
        return Task.FromResult<ResourceControllerResult>(null);
    }

}