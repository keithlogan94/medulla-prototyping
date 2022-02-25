using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using k8s;
using k8s.Models;
using Microsoft.Rest;
using KubernetesUtils;

namespace watch
{
    internal class Program
    {
        private async static Task Main(string[] args)
        {
            var dbs = await ParseCustomResource.GetResources<DatabaseSpec>("databases");
            Console.WriteLine(dbs);
           
        }
    }
}