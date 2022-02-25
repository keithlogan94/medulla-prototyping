using k8s;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace KubernetesUtils
{
    internal class CustomKubernetesObjectFinder
    {

        public async Task<string> FindObjects(string plural)
        {
            Console.WriteLine("Running Controller");
            IKubernetes client = new Kubernetes(KubernetesClientConfiguration.BuildConfigFromConfigFile());
            return (await client
                .ListClusterCustomObjectAsync(
                "medulla.recro.com", 
                "v1beta1", 
                plural, 
                true)
                )
                .ToString();
        }


    }
}
