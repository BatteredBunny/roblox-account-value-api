package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ErrorNoUserID = errors.New("no userid provided")

type collectiblesApiResponseData struct {
	UserAssetId                uint64  `json:"userAssetId"`
	SerialNumber               *uint64 `json:"serialNumber"`
	AssetId                    uint64  `json:"assetId"`
	Name                       string  `json:"name"`
	RecentAveragePrice         uint64  `json:"recentAveragePrice"`
	OriginalPrice              *uint64 `json:"originalPrice"`
	AssetStock                 *uint64 `json:"assetStock"`
	BuildersClubMembershipType uint64  `json:"buildersClubMembershipType"`
}

type collectiblesApiResponse struct {
	PreviousPageCursor string                        `json:"previousPageCursor"`
	NextPageCursor     string                        `json:"nextPageCursor"`
	Data               []collectiblesApiResponseData `json:"data"`
}

func collectiblesAPI(userid uint64, cursor string) (body *collectiblesApiResponse, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("https://inventory.roblox.com/v1/users/%d/assets/collectibles", userid))
	if err != nil {
		return nil, err
	}

	q := requestURL.Query()
	q.Set("limit", "100")
	q.Set("sortOrder", "Asc")
	q.Set("cursor", cursor)

	resp, err := http.Get(requestURL.String() + "?" + q.Encode())
	if err != nil {
		return nil, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawBody, &body)

	return
}

type accountValueResponse struct {
	TotalRobux uint64 `json:"total_robux"`
	InEuro     uint64 `json:"in_euro"`
}

func (app *Application) limitedAccountValueAPI(c *gin.Context) {
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

	var items []collectiblesApiResponseData
	nextCursor := ""
	for {
		body, err := collectiblesAPI(userid, nextCursor)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		items = append(items, body.Data...)
		if body.NextPageCursor == "" {
			break
		}

		body.NextPageCursor = nextCursor
	}

	var totalRobux uint64
	for _, item := range items {
		totalRobux += item.RecentAveragePrice
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, accountValueResponse{
		TotalRobux: totalRobux,
		InEuro:     totalRobux / app.config.RobuxToEuroRate,
	})
}

func canViewInventory(userid uint64) (canView *bool, err error) {
	resp, err := http.Get(fmt.Sprintf("https://inventory.roblox.com/v1/users/%d/can-view-inventory", userid))
	if err != nil {
		return nil, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body struct {
		CanView bool `json:"canView"`
	}
	if err = json.Unmarshal(rawBody, &body); err != nil {
		return nil, err
	}

	return &body.CanView, nil
}

func (app *Application) canViewInventoryAPI(c *gin.Context) {
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

	canView, err := canViewInventory(userid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, strconv.FormatBool(*canView))
}
