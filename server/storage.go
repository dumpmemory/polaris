package server

import (
	"fmt"
	"polaris/db"
	"polaris/log"
	"polaris/pkg/storage"
	"polaris/pkg/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Server) GetAllStorage(c *gin.Context) (interface{}, error) {
	data := s.db.GetAllStorage()
	return data, nil
}

func (s *Server) AddStorage(c *gin.Context) (interface{}, error) {
	var in db.StorageInfo
	if err := c.ShouldBindJSON(&in); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}

	if in.Implementation == "webdav" {
		//test webdav
		wd := in.ToWebDavSetting()
		st, err := storage.NewWebdavStorage(wd.URL, wd.User, wd.Password, wd.TvPath, false)
		if err != nil {
			return nil, errors.Wrap(err, "new webdav")
		}
		fs, err := st.ReadDir(".")
		if err != nil {
			return nil, errors.Wrap(err, "test read")
		}
		for _, f := range fs {
			log.Infof("file name: %v", f.Name())
		}
	}
	log.Infof("received add storage input: %v", in)
	err := s.db.AddStorage(&in)
	return nil, err
}

func (s *Server) DeleteStorage(c *gin.Context) (interface{}, error) {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, fmt.Errorf("id is not int: %v", ids)
	}
	err = s.db.DeleteStorage(id)
	return nil, err
}

func (s *Server) SuggestedSeriesFolderName(c *gin.Context) (interface{}, error) {
	ids := c.Param("tmdb_id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, fmt.Errorf("id is not int: %v", ids)
	}
	d, err := s.MustTMDB().GetTvDetails(id, s.language)
	if err != nil {
		return nil, errors.Wrap(err, "get tv details")
	} 
	name := d.Name
	originalName := d.OriginalName
	year := strings.Split(d.FirstAirDate, "-")[0]

	name = fmt.Sprintf("%s %s", name, originalName)

	if !utils.ContainsChineseChar(name) {
		name = originalName
	}
	if year != "" {
		name = fmt.Sprintf("%s (%s)", name, year)
	}
	log.Infof("tv series of tmdb id %v suggestting name is %v", id, name)
	return gin.H{"name": name}, nil
}

func (s *Server) SuggestedMovieFolderName(c *gin.Context) (interface{}, error) {
	ids := c.Param("tmdb_id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, fmt.Errorf("id is not int: %v", ids)
	}
	d1, err := s.MustTMDB().GetMovieDetails(id, s.language)
	if err != nil {
		return nil, errors.Wrap(err, "get movie details")
	}
	name := d1.Title
	originalName := d1.OriginalTitle
	year := strings.Split(d1.ReleaseDate, "-")[0]

	name = fmt.Sprintf("%s %s", name, originalName)

	if !utils.ContainsChineseChar(name) {
		name = originalName
	}
	if year != "" {
		name = fmt.Sprintf("%s (%s)", name, year)
	}
	log.Infof("tv series of tmdb id %v suggestting name is %v", id, name)
	return gin.H{"name": name}, nil
}
