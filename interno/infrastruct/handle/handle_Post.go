package handle

import (
	"log/slog"
	"net/http"

	"github.com/andreylsant/etldecliente/interno/csv"
	"github.com/gin-gonic/gin"
)

type Handle struct {}

// Gostaria de Receber um arquivo
func (h *Handle) Handle_Post(c *gin.Context) {
	// Pega o arquivo enviado no campo "file"
	// Define o caminho onde o arquivo será salvo

	csvReader := csv.ReaderEtl{}

	// 1. Pega o cabeçalho do arquivo enviado no multipart/form-data
	fileHeader, err := c.FormFile("file")
	if err != nil {
		slog.Error("[Erro ao obter cabeçalho do arquivo]", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não enviado"})
		return
	}

	// 2. Abre o arquivo. O objeto 'file' retornado implementa io.Reader
	file, err := fileHeader.Open()
	if err != nil {
		slog.Error("[Erro ao abrir o arquivo]", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar o arquivo"})
		return
	}

	defer file.Close() // Importante fechar para liberar memória do servidor

	csvReader.Reader(file)

	c.JSON(http.StatusOK, gin.H{"message": "ETL executado com sucesso!"})
}
