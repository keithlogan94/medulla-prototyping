namespace DatabaseGrpcService
{
    public class GreeterService : GrpcDatabase.DatabaseSvc.DatabaseSvcBase
    {
        private readonly ILogger<GrpcDatabase.DatabaseSvc.DatabaseSvcBase> _logger;
        public GreeterService(ILogger<GrpcDatabase.DatabaseSvc.DatabaseSvcBase> logger)
        {
            _logger = logger;
        }

        public override Task<GrpcDatabase.CreateDatabaseResponse> CreateDatabase(GrpcDatabase.CreateDatabaseRequest request, Grpc.Core.ServerCallContext context)
        {
            return Task.FromResult(new GrpcDatabase.CreateDatabaseResponse
            {
                Uuid = "test",
                Database = new GrpcDatabase.Database(),
            });
        }
    }
}
