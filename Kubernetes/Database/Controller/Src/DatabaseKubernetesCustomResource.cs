using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;


namespace KubernetesUtils
{

    public class MetaData
    {
        public string resourceVersion { get; set; }
        public int generation { get; set; }
        public DateTime creationTimestamp { get; set; }

        public string name { get; set; }

        public string uid { get; set; }

    }


    public class DatabaseSpec
    {
        public string databaseName { get; set; }

    }


    //[{"spec":{ "databaseName":"another_db_table"}},{ "apiVersion":"medulla.recro.com/v1beta1","kind":"Database","metadata":{ "creationTimestamp":"2022-02-23T15:33:05Z","generation":1,"managedFields":[{ "apiVersion":"medulla.recro.com/v1beta1","fieldsType":"FieldsV1","fieldsV1":{ "f:spec":{ ".":{ },"f:databaseName":{ } } },"manager":"kubectl-create","operation":"Update","time":"2022-02-23T15:33:05Z"}],"resourceVersion":"135275","uid":"869e5248-91a0-48a5-82da-fb820da0a353"},"spec":{ "databaseName":"my_db_table_name"} }]
    public class CustomResourceItem<Spec>
    {
        public string apiVersion { get; set; }
        public string kind { get; set; }
        public MetaData metadata { get; set; }
        public Spec spec { get; set; }

    }


    public class QueryCustomResource<Spec>
    {
        /*
         * {"apiVersion":"medulla.recro.com/v1beta1","kind":"DatabaseList","metadata":{ "continue":"","resourceVersion":"204447"}}
         */

        public string apiVersion { get; set; }
        public List<CustomResourceItem<Spec>> items { get; set; }

        public MetaData metadata { get; set; }

    }


    public class ParseCustomResource
    {

        public static async Task<QueryCustomResource<Spec>> GetResources<Spec>(string plural)
        {
            var finder = new CustomKubernetesObjectFinder();
            var str = await finder.FindObjects(plural);
            QueryCustomResource<Spec>? obj =
                JsonSerializer.Deserialize<QueryCustomResource<Spec>>(str);
            return obj;
        }

    }


}
