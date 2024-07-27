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
			"id": "mx1q2zgzlvu65f1",
			"created": "2024-07-27 21:02:03.163Z",
			"updated": "2024-07-27 21:02:03.163Z",
			"name": "scholarship_bodies",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "uj2bucbj",
					"name": "url",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "caqe3al0ctn1ova",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "sk3nwkgr",
					"name": "text",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_Il1NcVo` + "`" + ` ON ` + "`" + `scholarship_bodies` + "`" + ` (` + "`" + `url` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("mx1q2zgzlvu65f1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
