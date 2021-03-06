package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suusan2go/familog-api/domain/model"
	"github.com/suusan2go/familog-api/handler"
	"github.com/suusan2go/familog-api/registry"
)

// Map Generic Map
type Map map[string]interface{}

// JSONErrorHandler Handling Errors as JSON
func JSONErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	e := c.Echo()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else if e.Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = Map{"message": msg}
	}

	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			if err := c.NoContent(code); err != nil {
				goto ERROR
			}
		} else {
			if err := c.JSON(code, msg); err != nil {
				goto ERROR
			}
		}
	}
ERROR:
	e.Logger.Error(err)
}

func main() {
	db := model.InitDB()
	db.Migration()
	e := echo.New()
	e.Debug = true
	e.HTTPErrorHandler = JSONErrorHandler
	// middleware setting
	e.Use(middleware.Logger())
	reg := registry.Registry{DB: db}
	h := &handler.Handler{DB: db, Registry: &reg}

	// routing
	e.GET("/", h.GetAppInfo) // routing for healthcheck
	e.GET("/user", h.GetUser, h.Authenticate)
	e.PATCH("/user", h.PatchUser, h.Authenticate)
	e.POST("/device", h.PostDevice)
	e.POST("/push_notification_tokens", h.PostPushNotificationToken)
	e.POST("/session", h.PostSession)
	e.POST("/diaries", h.PostDiary, h.Authenticate)
	e.GET("/diaries", h.DiaryIndex, h.Authenticate)
	e.GET("/diaries/:id/invitation", h.GetDiaryInvitation, h.Authenticate)
	e.POST("/diaries/:id/invitation", h.PostDiaryInvitation, h.Authenticate)
	e.POST("/diary_invitation_verifications", h.PostDiaryInvitationVerification, h.Authenticate)
	e.POST("/diaries/:id/diary_entries", h.PostDiaryEntry, h.Authenticate)
	e.GET("/diaries/:id/diary_entries", h.GetDiaryEntries, h.Authenticate)
	e.GET("/diary_entries/:id", h.GetDiaryEntry, h.Authenticate)
	e.PATCH("/diary_entries/:id", h.PatchDiaryEntry, h.Authenticate)
	e.PATCH("/diary_entries/:id", h.PatchDiaryEntry, h.Authenticate)
	e.PATCH("/diary_entries/:diary_entry_id/images/:id", h.PatchDiaryEntryImage, h.Authenticate)
	e.DELETE("/diary_entries/:diary_entry_id/images/:id", h.DeleteDiaryEntryImage, h.Authenticate)

	e.Logger.Fatal(e.Start(":8080"))
}
