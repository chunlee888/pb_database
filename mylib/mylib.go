package mylib

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)


func AddRouteHello(e *core.ServeEvent)  {
	e.Router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/hello",
		Handler: func(c echo.Context) error {
			return c.String(200, "Hello world!")
		},
		Middlewares: []echo.MiddlewareFunc{
			apis.ActivityLogger(e.App),
		},
	})
}

func CreateCollection(e *core.ServeEvent, name string) error {
	_ , err := e.App.Dao().FindCollectionByNameOrId(name)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Collection '" + name + "' is already exist!")
		return nil
	}

	collection := &models.Collection{
		Name:       name,
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
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
		),
	}

	if err := e.App.Dao().SaveCollection(collection); err != nil {
		log.Println("Not able to save collection '" + name + "'")
		log.Fatal(err)
	}

	return nil
}