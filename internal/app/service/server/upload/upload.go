package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// SaveUploadedFile salva o arquivo recebido e retorna o path no diretório /tmp
func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	uniqueName := fmt.Sprintf("%s_%s%s", file.Filename[:len(file.Filename)-len(ext)], uuid.NewString(), ext)

	// Cria o diretório /tmp se não existir
	tmpDir := filepath.Join("./internal/app", "tmp")
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("erro ao criar diretório tmp: %w", err)
	}

	// Gera o caminho completo para o novo arquivo
	destPath := filepath.Join(tmpDir, uniqueName)

	dst, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo no destino: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	return destPath, nil
}

// CopyLocalFileToTmp copia um arquivo de um path local e salva no diretório /tmp, retornando o novo caminho
func CopyLocalFileToTmp(originalPath string) (string, error) {
	// Abre o arquivo de origem
	src, err := os.Open(originalPath)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir arquivo original: %w", err)
	}
	defer src.Close()

	// Gera nome único baseado no original
	ext := filepath.Ext(originalPath)
	name := filepath.Base(originalPath[:len(originalPath)-len(ext)])
	uniqueName := fmt.Sprintf("%s_%s%s", name, uuid.NewString(), ext)

	// Garante que o diretório tmp existe
	tmpDir := filepath.Join("./internal/app", "tmp")
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("erro ao criar diretório tmp: %w", err)
	}

	// Define caminho final
	destPath := filepath.Join(tmpDir, uniqueName)

	// Cria o destino
	dst, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo no destino: %w", err)
	}
	defer dst.Close()

	// Copia conteúdo
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	return destPath, nil
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("erro ao remover arquivo: %w", err)
	}
	return nil
}
