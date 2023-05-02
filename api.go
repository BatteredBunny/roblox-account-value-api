package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var ErrorNoUserID = errors.New("no userid provided")

func (app *Application) handleCollectiblesAccountValue(c *gin.Context) {
	type handleCollectiblesAccountValueResponseCollectible struct {
		Name         string  `json:"name"`
		Price        uint64  `json:"price"`
		ID           uint64  `json:"id"`
		SerialNumber *uint64 `json:"serialnumber,omitempty"`
	}
	type handleCollectiblesAccountValueResponse struct {
		TotalRobux   uint64                                              `json:"total_robux"`
		InEuro       uint64                                              `json:"in_euro"`
		Collectibles []handleCollectiblesAccountValueResponseCollectible `json:"collectibles"`
	}

	useridRaw, exists := c.GetQuery("userid")
	if !exists {
		c.String(http.StatusBadRequest, ErrorNoUserID.Error())
		c.Abort()
		return
	}

	userid, err := strconv.ParseUint(useridRaw, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid userid")
		c.Abort()
		return
	}

	rawCollectibles, err := getAllCollectibles(userid)
	if err != nil {
		app.logWarning.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var collectibles []handleCollectiblesAccountValueResponseCollectible
	for _, collectible := range rawCollectibles {
		collectibles = append(collectibles, handleCollectiblesAccountValueResponseCollectible{
			Name:         collectible.Name,
			Price:        collectible.RecentAveragePrice,
			ID:           collectible.AssetId,
			SerialNumber: collectible.SerialNumber,
		})
	}

	var totalRobux uint64
	for _, collectible := range rawCollectibles {
		totalRobux += collectible.RecentAveragePrice
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, handleCollectiblesAccountValueResponse{
		TotalRobux:   totalRobux,
		InEuro:       totalRobux / app.config.RobuxPerEuro,
		Collectibles: collectibles,
	})
}

func (app *Application) handleCanViewInventory(c *gin.Context) {
	useridRaw, exists := c.GetQuery("userid")
	if !exists {
		c.String(http.StatusBadRequest, ErrorNoUserID.Error())
		c.Abort()
		return
	}

	userid, err := strconv.ParseUint(useridRaw, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid userid")
		c.Abort()
		return
	}

	canView, err := canViewInventoryAPI(userid)
	if err != nil {
		app.logWarning.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, strconv.FormatBool(*canView))
}

func (app *Application) handleProfileInfo(c *gin.Context) {
	type handleProfileInfoResponse struct {
		Username    string `json:"username"`
		DisplayName string `json:"displayname"`
		Avatar      string `json:"avatar"`
	}

	useridRaw, exists := c.GetQuery("userid")
	if !exists {
		c.String(http.StatusBadRequest, ErrorNoUserID.Error())
		c.Abort()
		return
	}

	userid, err := strconv.ParseUint(useridRaw, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid userid")
		c.Abort()
		return
	}

	info, err := profileInfoAPI(userid)
	if err != nil {
		app.logWarning.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	avatarUrl, err := profileAvatarAPI(userid)
	if err != nil {
		app.logWarning.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, handleProfileInfoResponse{
		Username:    info.Name,
		DisplayName: info.DisplayName,
		Avatar:      *avatarUrl,
	})
}

func (app *Application) handleExchangeRate(c *gin.Context) {
	type handleExchangeRateResponse struct {
		RobuxPerEuro uint64 `json:"robux_per_euro"`
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, handleExchangeRateResponse{RobuxPerEuro: app.config.RobuxPerEuro})
}
