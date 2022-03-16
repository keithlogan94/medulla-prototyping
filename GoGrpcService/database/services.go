package database

import (
	context2 "context"
	"database_service/kubernetes"
)

type Server struct {
}

func (s *Server) CreateDatabase(ctx context2.Context, request *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {

	database := kubernetes.Database{}
	err := kubernetes.CreateDatabase(&database)
	if err != nil {
		panic(err)
	}
	var res = CreateDatabaseResponse{
		Database: &Database{
			Name:      "test",
			Role:      "test role",
			Collation: "test",
			Dialect:   "mysql",
		},
	}
	return &res, nil
}

func (s *Server) GetDatabases(ctx context2.Context, request *GetDatabasesRequest) (*GetDatabasesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateDatabases(ctx context2.Context, request *UpdateDatabasesRequest) (*UpdateDatabasesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteDatabases(ctx context2.Context, request *DeleteDatabasesRequest) (*DeleteDatabasesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreateModel(ctx context2.Context, request *CreateModelRequest) (*CreateModelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetModels(ctx context2.Context, request *GetModelsRequest) (*GetModelsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateModels(ctx context2.Context, request *UpdateModelsRequest) (*UpdateModelsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteModels(ctx context2.Context, request *DeleteModelsRequest) (*DeleteModelsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) QueryModel(ctx context2.Context, request *QueryModelRequest) (*QueryModelResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedDatabaseSvcServer() {
	//TODO implement me
	panic("implement me")
}
