package database

import "database_service/kubernetes"

func changeKubernetesDatabasesToResponseDatabases(kdbs []kubernetes.Database) []*Database {
	var dbs []*Database
	for i := 0; i < len(kdbs); i++ {
		dbs = append(dbs,
			&Database{
				Name:      kdbs[i].Name,
				Role:      "",
				Collation: "",
				Dialect:   "",
			})
	}
	return dbs
}
