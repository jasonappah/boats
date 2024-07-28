package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mx1q2zgzlvu65f1")
		if err != nil {
			return err
		}

		collection.Name = "scholarship_html"

		if err := json.Unmarshal([]byte(`[
			"CREATE UNIQUE INDEX ` + "`" + `idx_Il1NcVo` + "`" + ` ON ` + "`" + `scholarship_html` + "`" + ` (` + "`" + `scholarship_id` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// update
		edit_scholarship_id := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "uj2bucbj",
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
		}`), edit_scholarship_id); err != nil {
			return err
		}
		collection.Schema.AddField(edit_scholarship_id)

		// update
		edit_html := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "sk3nwkgr",
			"name": "html",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_html); err != nil {
			return err
		}
		collection.Schema.AddField(edit_html)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("mx1q2zgzlvu65f1")
		if err != nil {
			return err
		}

		collection.Name = "scholarship_bodies"

		if err := json.Unmarshal([]byte(`[
			"CREATE UNIQUE INDEX ` + "`" + `idx_Il1NcVo` + "`" + ` ON ` + "`" + `scholarship_bodies` + "`" + ` (` + "`" + `url` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		// update
		edit_scholarship_id := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_scholarship_id); err != nil {
			return err
		}
		collection.Schema.AddField(edit_scholarship_id)

		// update
		edit_html := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), edit_html); err != nil {
			return err
		}
		collection.Schema.AddField(edit_html)

		return dao.SaveCollection(collection)
	})
}
