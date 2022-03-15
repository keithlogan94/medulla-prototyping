namespace DatabaseGrpcService
{
    public class GreeterService : GrpcDatabase.DatabaseSvc.DatabaseSvcBase
    {
        private readonly ILogger<GrpcDatabase.DatabaseSvc.DatabaseSvcBase> _logger;
        public GreeterService(ILogger<GrpcDatabase.DatabaseSvc.DatabaseSvcBase> logger)
        {
            _logger = logger;
        }

        public override async Task<GrpcDatabase.CreateDatabaseResponse> CreateDatabase(GrpcDatabase.CreateDatabaseRequest request, Grpc.Core.ServerCallContext context)
        {
            System.Console.WriteLine("Grpc");
            return new GrpcDatabase.CreateDatabaseResponse() {
                
            };
        }
    }
}
