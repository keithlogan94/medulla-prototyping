using k8s;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace KubernetesUtils
{
    internal class CustomKubernetesObjectFinder
    {

        public async Task<int> FindObjects(string plural)
        {
            Console.WriteLine("Running Controller");
            var config = KubernetesClientConfiguration.BuildConfigFromConfigFile();
            IKubernetes client = new Kubernetes(config);
            var databaseListResp = await client.ListClusterCustomObjectAsync("medulla.recro.com", "v1beta1", plural, true);
            Console.WriteLine(databaseListResp);
            return 1;
        }


    }
}
