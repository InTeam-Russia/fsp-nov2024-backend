package api

import (
	"net/http"

	"github.com/InTeam-Russia/go-backend-template/internal/apierr"
	"github.com/InTeam-Russia/go-backend-template/internal/filters/filter"
	"github.com/InTeam-Russia/go-backend-template/internal/helpers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRoutes(
	r *gin.Engine,
	logger *zap.Logger,
	filterRepo filter.FilterRepository,
) {
	f := r.Group("/events/filters")

	f.GET("/sports", func(c *gin.Context) {
		data, err := filterRepo.GetSports()
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/teams", func(c *gin.Context) {
		sport := helpers.StrPtr(c.Query("sport"))
		if *sport == "" {
			sport = nil
		}
		data, err := filterRepo.GetTeams(sport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/competitionTypes", func(c *gin.Context) {
		team := helpers.StrPtr(c.Query("team"))
		if *team == "" {
			team = nil
		}
		data, err := filterRepo.GetCompetitionTypes(team)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/genderAndAgeGroups", func(c *gin.Context) {
		team := helpers.StrPtr(c.Query("team"))
		if *team == "" {
			team = nil
		}
		data, err := filterRepo.GetGenderAndAgeGroups(team)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/programsDisciplines", func(c *gin.Context) {
		team := helpers.StrPtr(c.Query("team"))
		if *team == "" {
			team = nil
		}
		data, err := filterRepo.GetProgramsDisciplines(team)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/countries", func(c *gin.Context) {
		data, err := filterRepo.GetCountries()
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/regions", func(c *gin.Context) {
		country := helpers.StrPtr(c.Query("country"))
		if *country == "" {
			country = nil
		}
		data, err := filterRepo.GetRegions(country)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})

	f.GET("/cities", func(c *gin.Context) {
		country := helpers.StrPtr(c.Query("country"))
		region := helpers.StrPtr(c.Query("region"))
		if *country == "" {
			country = nil
		}
		if *region == "" {
			region = nil
		}
		data, err := filterRepo.GetCities(country, region)
		if err != nil {
			c.JSON(http.StatusInternalServerError, apierr.InternalServerError)
			logger.Error(err.Error())
			return
		}
		c.JSON(http.StatusOK, data)
	})
}
