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
			"id": "08fzz8ptsd4uzy5",
			"created": "2024-07-28 00:22:51.305Z",
			"updated": "2024-07-28 00:22:51.305Z",
			"name": "scholarship_shortlist",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "4rncht2z",
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
					"id": "xti6gzol",
					"name": "status",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"not_started",
							"in_progress",
							"editing_reviewing",
							"submitted",
							"not_planned"
						]
					}
				},
				{
					"system": false,
					"id": "4ujtdbhy",
					"name": "notes",
					"type": "editor",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_9M7USFs` + "`" + ` ON ` + "`" + `scholarship_shortlist` + "`" + ` (` + "`" + `scholarship_id` + "`" + `)"
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

		collection, err := dao.FindCollectionByNameOrId("08fzz8ptsd4uzy5")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
