package api

import (
	"github.com/gin-gonic/gin"
	"golang-telegram-bot/models"
	"net/http"
)


// Albums slice to seed record album data.
var InitialAlbums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @Summary List albums
// @Description get all albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {object} models.AlbumList
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	albums := models.AlbumList{Albums: InitialAlbums}
	c.IndentedJSON(http.StatusOK, albums)
}


// @Summary create album
// @Description post albums
// @Tags albums
// @Accept json
// @Produce json
// @Param album body models.Album true "post albums"
// @Success 200 {object} models.Album
// @Router /albums [post]
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	InitialAlbums = append(InitialAlbums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}


// @Summary get albums by id
// @Description Get albums by Id
// @ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path integer  true "album ID"
// @Success 200 {object} models.Album
// @Router /albums/{id} [get]
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range InitialAlbums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
