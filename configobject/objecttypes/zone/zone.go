package zone

import (
	"git.icinga.com/icingadb/icingadb-main/configobject"
	"git.icinga.com/icingadb/icingadb-main/connection"
	"git.icinga.com/icingadb/icingadb-main/utils"
)

var (
	ObjectInformation configobject.ObjectInformation
	Fields         = []string{
		"id",
		"env_id",
		"name_checksum",
		"properties_checksum",
		"parents_checksum",
		"name",
		"name_ci",
		"is_global",
		"parent_id",
	}
)

type Zone struct {
	Id                  string  `json:"id"`
	EnvId               string  `json:"environment_id"`
	NameChecksum        string  `json:"name_checksum"`
	PropertiesChecksum  string  `json:"properties_checksum"`
	ParentsChecksum     string  `json:"parents_checksum"`
	Name                string  `json:"name"`
	NameCi              *string `json:"name_ci"`
	IsGlobal			bool	`json:"is_global"`
	ParentId            string  `json:"parent_id"`
}

func NewZone() connection.Row {
	z := Zone{}
	z.NameCi = &z.Name

	return &z
}

func (z *Zone) InsertValues() []interface{} {
	v := z.UpdateValues()

	return append([]interface{}{utils.Checksum(z.Id)}, v...)
}

func (z *Zone) UpdateValues() []interface{} {
	v := make([]interface{}, 0)

	v = append(
		v,
		utils.Checksum(z.EnvId),
		utils.Checksum(z.NameChecksum),
		utils.Checksum(z.PropertiesChecksum),
		utils.Checksum(z.ParentsChecksum),
		z.Name,
		z.NameCi,
		utils.Bool[z.IsGlobal],
		utils.Checksum(z.ParentId),
	)

	return v
}

func (z *Zone) GetId() string {
	return z.Id
}

func (z *Zone) SetId(id string) {
	z.Id = id
}

func init() {
	ObjectInformation = configobject.ObjectInformation{
		ObjectType: "zone",
		Factory: NewZone,
		BulkInsertStmt: connection.NewBulkInsertStmt("zone", Fields),
		BulkDeleteStmt: connection.NewBulkDeleteStmt("zone"),
		BulkUpdateStmt: connection.NewBulkUpdateStmt("zone", Fields),
	}
}