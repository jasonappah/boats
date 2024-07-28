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
			"id": "wdne3tgrhus90bk",
			"created": "2024-07-28 00:17:55.360Z",
			"updated": "2024-07-28 00:17:55.360Z",
			"name": "parsed_scholarships",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "0hhblbed",
					"name": "scholarship_id",
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
					"id": "qopv2dgp",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "ue8uy5jn",
					"name": "cleaned_text",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "cs46jjtd",
					"name": "summary",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "ok4vus2h",
					"name": "eligible",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "qlfv05nn",
					"name": "viewed",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "kmmcsx6x",
					"name": "due_date",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_QOI7oO6` + "`" + ` ON ` + "`" + `parsed_scholarships` + "`" + ` (` + "`" + `scholarship_id` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_XJwivBu` + "`" + ` ON ` + "`" + `parsed_scholarships` + "`" + ` (\n  ` + "`" + `eligible` + "`" + `,\n  ` + "`" + `viewed` + "`" + `,\n  ` + "`" + `due_date` + "`" + `\n)"
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

		collection, err := dao.FindCollectionByNameOrId("wdne3tgrhus90bk")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
