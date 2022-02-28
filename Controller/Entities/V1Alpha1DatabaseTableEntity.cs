using k8s.Models;
using KubeOps.Operator.Entities;
using System.ComponentModel;
using System.ComponentModel.DataAnnotations;

namespace DatabaseControllerKubeOps.Controller.Entities;

public class ValidateSpec
{
    public string? Is { get; set; }
    public string? Not { get; set; }
    public bool? IsEmail { get; set; }
    public bool? IsUrl { get; set; }
    public bool? IsIp { get; set; }
    public bool? IsIpV4 { get; set; }
    public bool? IsIpV6 { get; set; }
    public bool? IsAlpha { get; set; }
    public bool? IsAlphaNumeric { get; set; }
    public bool? IsNumeric { get; set; }
    public bool? IsInt { get; set; }
    public bool? IsFloat { get; set; }
    public bool? IsDecimal { get; set; }
    public bool? IsLowercase { get; set; }
    public bool? IsUppercase { get; set; }
    public bool? IsNull { get; set; }
    public bool? NotNull { get; set; }
    public bool? NotEmpty { get; set; }
}

public class ColumnSpec
{
    public string? ColumnName { get; set; }
    public string? Type { get; set; }
    public bool AllowNull { get; set; }
    public string? DefaultValue { get; set; }
    public bool AutoIncrement { get; set; }
    public bool PrimaryKey { get; set; }
    public string? Field { get; set; }
    public string? Unique { get; set; }
    public string? Comment { get; set; }
    public ValidateSpec? Validate { get; set; }

}


[Description("A CustomResourceDefinition which allows building a database table in Medulla.")]
[KubernetesEntity(
    ApiVersion = "v1alpha1",
    Kind = "DatabaseTable",
    Group = "medulla.recro.com",
    PluralName = "databasetables")]
public class V1Alpha1DatabaseTableEntity : CustomKubernetesEntity
{
    [Required]
    public string? Name { get; set; }
    [Required]
    public List<Dictionary<string, ColumnSpec>>? Columns { get; set; }
}
