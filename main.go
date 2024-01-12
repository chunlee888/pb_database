package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

// func CreateNewCollection(pb  &pocketbase.PocketBase) {
//     collection := &models.Collection{
//     Name:       "example",
//     Type:       models.CollectionTypeBase,
//     ListRule:   nil,
//     ViewRule:   types.Pointer("@request.auth.id != ''"),
//     CreateRule: types.Pointer(""),
//     UpdateRule: types.Pointer("@request.auth.id != ''"),
//     DeleteRule: nil,
//     Schema:     schema.NewSchema(
//         &schema.SchemaField{
//             Name:     "title",
//             Type:     schema.FieldTypeText,
//             Required: true,
//             Options:  &schema.TextOptions{
//                 Max: types.Pointer(10),
//             },
//         },
//         &schema.SchemaField{
//             Name:     "user",
//             Type:     schema.FieldTypeRelation,
//             Required: true,
//             Options:  &schema.RelationOptions{
//                 MaxSelect:     types.Pointer(1),
//                 CollectionId:  "ae40239d2bc4477",
//                 CascadeDelete: true,
//             },
//         },
//     ),
//     Indexes: types.JsonArray[string]{
//         "CREATE UNIQUE INDEX idx_user ON example (user)",
//         },
//     }

//     if err := pb.Dao().SaveCollection(collection); err != nil {
//     	log.Fatal(err)
//     }
// }

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// add new "GET /hello" route to the app router (echo)
        e.Router.AddRoute(echo.Route{
            Method: http.MethodGet,
            Path:   "/hello",
            Handler: func(c echo.Context) error {
                return c.String(200, "Hello world!")
            },
            Middlewares: []echo.MiddlewareFunc{
                apis.ActivityLogger(app),
            },
        })

		_ , err := app.Dao().FindCollectionByNameOrId("example")

		if err != nil {
            log.Println("errrr 1111")
            log.Println(err)
		} else {
            log.Println("Collection example is already exist!")
            return nil
        }

		collection := &models.Collection{
			Name:       "example",
			Type:       models.CollectionTypeBase,
			ListRule:   nil,
			ViewRule:   types.Pointer("@request.auth.id != ''"),
			CreateRule: types.Pointer(""),
			UpdateRule: types.Pointer("@request.auth.id != ''"),
			DeleteRule: nil,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "title",
					Type:     schema.FieldTypeText,
					Required: true,
					Options: &schema.TextOptions{
						Max: types.Pointer(10),
					},
				},
				&schema.SchemaField{
					Name:     "user",
					Type:     schema.FieldTypeRelation,
					Required: true,
					Options: &schema.RelationOptions{
						MaxSelect:     types.Pointer(1),
						CollectionId:  "ae40239d2bc4477",
						CascadeDelete: true,
					},
				},
			),
			Indexes: types.JsonArray[string]{
				"CREATE UNIQUE INDEX idx_user ON example (user)",
			},
		}

		if err := e.App.Dao().SaveCollection(collection); err != nil {
            log.Println("xxxx")
			log.Fatal(err)
		}

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
