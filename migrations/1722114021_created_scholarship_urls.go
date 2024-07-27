package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "caqe3al0ctn1ova",
			"created": "2024-07-27 21:00:21.692Z",
			"updated": "2024-07-27 21:00:21.692Z",
			"name": "scholarship_urls",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "wnkbvilj",
					"name": "url",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_WfNrA2X` + "`" + ` ON ` + "`" + `scholarship_urls` + "`" + ` (` + "`" + `url` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("caqe3al0ctn1ova")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
