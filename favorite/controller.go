package favorite

import (
	"encoding/json"
	"net/http"

	"github.com/arikama/koran-backend/constants"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetFavCtrl(favManager FavManager, userManager managers.UserManager) gin.HandlerFunc {
	type Item struct {
		Id    int `json:"id"`
		Surah int `json:"surah"`
		Verse int `json:"verse"`
	}

	type Response struct {
		Items []*Item `json:"favorites"`
	}

	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())

		user, err := userManager.GetUser(accessToken)

		if err != nil {
			utils.JsonError(c, http.StatusNotFound, err)
			return
		}

		favs, err := favManager.GetFavs(user.Email)

		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		items := []*Item{}

		for _, fav := range favs {
			item := Item{
				Id:    fav.Id,
				Surah: fav.Surah,
				Verse: fav.Verse,
			}
			items = append(items, &item)
		}

		utils.JsonData(c, http.StatusOK, Response{
			Items: items,
		})
	}
}

func PostFavCtrl(favManager FavManager, userManager managers.UserManager) func(*gin.Context) {
	type Request struct {
		Surah int `json:"surah"`
		Verse int `json:"verse"`
	}

	type Item struct {
		Id    int `json:"id"`
		Surah int `json:"surah"`
		Verse int `json:"verse"`
	}

	type Response struct {
		Items []*Item `json:"favorites"`
	}

	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())

		user, err := userManager.GetUser(accessToken)

		if err != nil {
			utils.JsonError(c, http.StatusNotFound, err)
			return
		}

		request := Request{}

		err = json.NewDecoder(c.Request.Body).Decode(&request)

		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}

		err = favManager.CreateFav(user.Email, request.Surah, request.Verse)

		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		favs, err := favManager.GetFavs(user.Email)

		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		items := []*Item{}

		for _, fav := range favs {
			items = append(items, &Item{
				Id:    fav.Id,
				Surah: fav.Surah,
				Verse: fav.Verse,
			})
		}

		utils.JsonData(c, http.StatusOK, Response{
			Items: items,
		})
	}
}

func PostFavRemoveCtrl(favManager FavManager, userManager managers.UserManager) func(*gin.Context) {
	type Request struct {
		Id int `json:"id"`
	}

	type Item struct {
		Id    int `json:"id"`
		Surah int `json:"surah"`
		Verse int `json:"verse"`
	}

	type Response struct {
		Items []*Item `json:"favorites"`
	}

	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())

		user, err := userManager.GetUser(accessToken)

		if err != nil {
			utils.JsonError(c, http.StatusNotFound, err)
			return
		}

		request := Request{}

		err = json.NewDecoder(c.Request.Body).Decode(&request)

		if err != nil {
			utils.JsonError(c, http.StatusBadRequest, err)
			return
		}

		err = favManager.DeleteFav(request.Id)

		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		favs, err := favManager.GetFavs(user.Email)

		if err != nil {
			utils.JsonError(c, http.StatusInternalServerError, err)
			return
		}

		items := []*Item{}

		for _, fav := range favs {
			items = append(items, &Item{
				Id:    fav.Id,
				Surah: fav.Surah,
				Verse: fav.Verse,
			})
		}

		utils.JsonData(c, http.StatusOK, Response{
			Items: items,
		})
	}
}
