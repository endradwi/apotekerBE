package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// All Movies godoc
// @Summary Movies
// @Description  Get All Movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param search query string false "Search Movie"
// @Param page query int false "Page Movie"
// @Param limit query int false "Limit Movie"
// @Param sort query string false "Sort Movie"
// @Success 200 {object} Response{results=models.ListMovie}
// @Router /movies [get]
func GetAllMovies(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	sortmovie := ctx.DefaultQuery("sort", "ASC")
	if sortmovie != "ASC" {
		sortmovie = "DESC"
	}

	var movies models.ListMovie
	var count int
	get := lib.Redis().Get(context.Background(), ctx.Request.RequestURI)
	getCount := lib.Redis().Get(context.Background(),
		fmt.Sprintf("count+%s", ctx.Request.RequestURI))
	if get.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &movies)
	} else {
		movies = models.FindAllMovie(page, limit, search, sortmovie)
		change, _ := json.Marshal(movies)
		lib.Redis().Set(
			context.Background(),
			ctx.Request.RequestURI,
			string(change),
			0,
		)
	}
	if getCount.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &count)
	} else {
		count = models.CountData(search)
		change, _ := json.Marshal(count)
		lib.Redis().Set(context.Background(),
			fmt.Sprintf("count+%s", ctx.Request.RequestURI),
			string(change),
			0,
		)
	}
	// count = models.CountData(search)
	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	log.Println("errorrrr", totalPage)
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	prevPage := page - 1
	if prevPage < 2 {
		prevPage = 0
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: "See All Users",
		PageInfo: PageInfo{
			CurentPage: page,
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPage:  totalPage,
			TotalData:  count,
		},
		Results: movies,
	})
}

// Update Movie godoc
// @Schemes
// @Description  Updated Movies
// @Tags Movies
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Edit_Movie formData models.Movie_body true "Edit Movie"
// @Success 200 {object} Response{results=models.Movie_Data}
// @Router /movies/:id [get]
func EditMovie(ctx *gin.Context) {
	paramId, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.FindOneMovie(paramId)
	// if paramId != profile.Id {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// 	})
	// 	return
	// }
	// handling body form without file

	var updatedMovie models.Movie_body
	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	file, _ := ctx.FormFile("image")

	if file.Filename != "" {
		filename := uuid.New().String()

		// handling extentioin .jpg dll
		splitedFilename := strings.Split(file.Filename, ".")
		ext := splitedFilename[len(splitedFilename)-1]
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Must Fill .jpg, .jpeg, .png",
			})
			return
		}

		// handling name file
		ctx.SaveUploadedFile(file, fmt.Sprintf("updload/profile/%s", storedFile))
		movie.Image = storedFile
	}

	// Validation Size File
	maxfile := 1 * 1024 * 1024
	if file.Size > int64(maxfile) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "File to Large",
		})
		return
	}

	updated := models.UpdateMovie(updatedMovie)

	fmt.Println("data upload =", updated)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update Movie Success",
		Results: updated,
	})
}

func GetDetailMoviesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	find := models.FindDetailMovie(id)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "You Get Movie By ID",
		Results: find,
	})

}
func GetMoviesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	find := models.FindOneMovie(id)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "You Get Movie By ID",
		Results: find,
	})

}

// Add Movie godoc
// @Schemes
// @Description Add Movies
// @Tags Movies
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Edit_Movie formData models.Movie_body true "Add New Movie"
// @Success 200 {object} Response{results=models.Movie_Data}
// @Router /movies/:id [post]
func SaveMovies(ctx *gin.Context) {
	var formData models.Movie_body
	text := ctx.ShouldBind(&formData)
	log.Println("data=", text)
	// file, _ := ctx.FormFile("image")
	// filename := uuid.New().String()
	// splitfile := strings.Split(file.Filename, ".")[1]
	// ext := splitfile[len(splitfile)-1]
	// storedFile := fmt.Sprintf("%s.%s", filename, ext)
	// if file.Size > 2<<8 {
	// ctx.JSON(400, Response{
	// Success: false,
	// Message: "Image to large",
	// })
	// return
	// }
	// ctx.SaveUploadedFile(file, fmt.Sprintf("uploads/movies/%s", storedFile))
	// if file.Filename != "" {
	// formData.Image = storedFile
	// }

	temp, _ := models.InsertMovie(formData)

	// data := lib.Redis().Scan(context.Background(), 0, "", 0).Iterator()
	// for data.Next(context.Background()) {

	// }
	// var row = Movie{
	// 	Tittle:      formData.Tittle,
	// 	Image:       formData.Image,
	// 	Description: formData.Description,
	// }
	//
	// row.Id = len(ListMovie) + 1
	// ListMovie = append(ListMovie, row)
	ctx.JSON(200, Response{
		Success: true,
		Message: "Your Movie Saved",
		Results: temp,
	})
}

// Add Movie godoc
// @Schemes
// @Description  Delete One Movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param id query string true "Delete Movie"
// @Success 200 {object} Response{results=models.ListMovie}
// @Router /movies/:id [delete]
func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	deleted := models.DeleteMovie(id)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Deleted Success",
		Results: deleted,
	})

}
