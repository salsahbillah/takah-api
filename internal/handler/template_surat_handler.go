package handler

import (
  "net/http"
  "strconv"

  "takah-api/internal/model"

  "github.com/gin-gonic/gin"
)

var templateSuratData = []model.TemplateSuratResponse{
  {
    ID:           1,
    TakahID:      1,
    TakahCode:    "SKET",
    TemplateName: "Template Surat Keterangan",
    Content:      "Yang bertanda tangan di bawah ini menerangkan bahwa...",
    CreatedAt:    "2026-06-18 10:00",
  },
}

func GetAllTemplateSurat(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "Data template surat berhasil diambil",
    "data":    templateSuratData,
  })
}

func GetTemplateSuratByID(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "ID template surat tidak valid"})
    return
  }

  for _, template := range templateSuratData {
    if template.ID == id {
      c.JSON(http.StatusOK, gin.H{
        "message": "Data template surat berhasil diambil",
        "data":    template,
      })
      return
    }
  }

  c.JSON(http.StatusNotFound, gin.H{"message": "Data template surat tidak ditemukan"})
}

func CreateTemplateSurat(c *gin.Context) {
  var request model.TemplateSuratRequest

  if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Data template surat wajib diisi dengan benar"})
    return
  }

  takah, found := findTakahByID(request.TakahID)
  if !found {
    c.JSON(http.StatusNotFound, gin.H{"message": "Master takah tidak ditemukan"})
    return
  }

  response := model.TemplateSuratResponse{
    ID:           len(templateSuratData) + 1,
    TakahID:      request.TakahID,
    TakahCode:    takah.Code,
    TemplateName: request.TemplateName,
    Content:      request.Content,
    CreatedAt:    "2026-06-18 10:00",
  }

  templateSuratData = append(templateSuratData, response)

  c.JSON(http.StatusCreated, gin.H{
    "message": "Data template surat berhasil dibuat",
    "data":    response,
  })
}

func UpdateTemplateSurat(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "ID template surat tidak valid"})
    return
  }

  var request model.TemplateSuratRequest

  if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Data template surat wajib diisi dengan benar"})
    return
  }

  takah, found := findTakahByID(request.TakahID)
  if !found {
    c.JSON(http.StatusNotFound, gin.H{"message": "Master takah tidak ditemukan"})
    return
  }

  response := model.TemplateSuratResponse{
    ID:           id,
    TakahID:      request.TakahID,
    TakahCode:    takah.Code,
    TemplateName: request.TemplateName,
    Content:      request.Content,
    CreatedAt:    "2026-06-18 10:00",
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Data template surat berhasil diupdate",
    "data":    response,
  })
}

func DeleteTemplateSurat(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "ID template surat tidak valid"})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Data template surat berhasil dihapus",
    "data": gin.H{
      "id": id,
    },
  })
}