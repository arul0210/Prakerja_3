package controllers

import (
	"net/http"

	"prakerja3/configs"
	"prakerja3/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetKaryawanController(c echo.Context) error {
	var karyawans []models.Karyawan

	result := configs.DB.Find(&karyawans)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal mengambil data karyawan", Data: nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status: true, Message: "Berhasil", Data: karyawans,
	})
}

func InsertKaryawanController(c echo.Context) error {
	var insertKaryawan models.Karyawan
	c.Bind(&insertKaryawan)

	//logic bisnis
	//cek database sudah data sudah ada
	//409
	result := configs.DB.First(&models.Karyawan{}, "email = ", insertKaryawan.Email)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "Email telah digunakan !!", Data: nil,
		})
	}

	//masuk ke database
	result = configs.DB.Create(&insertKaryawan)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal insert ke database !!", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "berhasil insert database karyawan", Data: insertKaryawan,
	})
}
