package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type collectiblesAPIResponse struct {
	PreviousPageCursor string                        `json:"previousPageCursor"`
	NextPageCursor     string                        `json:"nextPageCursor"`
	Data               []collectiblesAPIResponseData `json:"data"`
}

type collectiblesAPIResponseData struct {
	UserAssetId                uint64  `json:"userAssetId"`
	SerialNumber               *uint64 `json:"serialNumber"`
	AssetId                    uint64  `json:"assetId"`
	Name                       string  `json:"name"`
	RecentAveragePrice         uint64  `json:"recentAveragePrice"`
	OriginalPrice              *uint64 `json:"originalPrice"`
	AssetStock                 *uint64 `json:"assetStock"`
	BuildersClubMembershipType uint64  `json:"buildersClubMembershipType"`
}

// collectiblesAPI is a wrapper for https://inventory.roblox.com/v1/users/%d/assets/collectibles
func collectiblesAPI(userid uint64, cursor string) (body *collectiblesAPIResponse, err error) {
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

// profileAvatarAPI is a wrapper for https://thumbnails.roblox.com/v1/users/avatar
func profileAvatarAPI(userid uint64) (*string, error) {
	resp, err := http.Get(fmt.Sprintf("https://thumbnails.roblox.com/v1/users/avatar?userIds=%d&size=720x720&format=Png&isCircular=true", userid))
	if err != nil {
		return nil, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body struct {
		Data []struct {
			TargetId int    `json:"targetId"`
			State    string `json:"state"`
			ImageUrl string `json:"imageUrl"`
		} `json:"data"`
	}

	if err = json.Unmarshal(rawBody, &body); err != nil {
		return nil, err
	}

	return &body.Data[0].ImageUrl, nil
}

type profileInfoAPIResponse struct {
	Description            string      `json:"description"`
	Created                time.Time   `json:"created"`
	IsBanned               bool        `json:"isBanned"`
	ExternalAppDisplayName interface{} `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool        `json:"hasVerifiedBadge"`
	Id                     int         `json:"id"`
	Name                   string      `json:"name"`
	DisplayName            string      `json:"displayName"`
}

// profileInfoAPI is a wrapper for https://users.roblox.com/v1/users/%d
func profileInfoAPI(userid uint64) (body *profileInfoAPIResponse, err error) {
	resp, err := http.Get(fmt.Sprintf("https://users.roblox.com/v1/users/%d", userid))
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

// canViewInventoryAPI is a wrapper for https://inventory.roblox.com/v1/users/%d/can-view-inventory
func canViewInventoryAPI(userid uint64) (canView *bool, err error) {
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
