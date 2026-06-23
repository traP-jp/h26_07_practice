package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	traq "github.com/traPtitech/go-traq"
)

const traqAPIBaseURL = "https://q.trap.jp/api/v3"

func GetUnreadManagement(c *echo.Context) error {
	token := os.Getenv("TOKEN")
	config := traq.NewConfiguration()
	config.AddDefaultHeader("Cookie", "r_session="+token)
	apiClient := traq.NewAPIClient(config)

	user, _, err := apiClient.MeAPI.GetMe(c.Request().Context()).Execute()
	unreadMassages, _, err := apiClient.NotificationAPI.GetMyUnreadChannels(c.Request().Context()).Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"body": unreadMassages, "userID": user.Name})
}
