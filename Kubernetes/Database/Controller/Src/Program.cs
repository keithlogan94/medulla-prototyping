using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
using k8s;
using k8s.Models;
using Microsoft.Rest;
using KubernetesUtils;
using Newtonsoft.Json;

namespace watch
{
    internal class Program
    {
        private async static Task Main(string[] args)
        {
            var finder = new CustomKubernetesObjectFinder();
            QueryCustomResource<DatabaseSpec> db = await ParseCustomResource.GetResources<DatabaseSpec>("databases");
            Console.WriteLine(db.apiVersion);
            //var str = await finder.FindObjects("databases");

            //Console.WriteLine(str["apiVersion"]);
            //var values = JsonConvert.DeserializeObject<Dictionary<string, string>>(str);
            //Console.WriteLine(values);
           
        }
    }
}