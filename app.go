package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SelectFile opens a file selection dialog and returns the selected file path
func (a *App) SelectFile() (string, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "ファイルを選択してください",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "PDFファイル (*.pdf)",
				Pattern:     "*.pdf",
			},
			{
				DisplayName: "画像ファイル (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return file, nil
}

// SelectDirectory opens a directory selection dialog and returns the selected directory path
func (a *App) SelectDirectory() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "ディレクトリを選択してください",
	})
	if err != nil {
		return "", fmt.Errorf("ディレクトリの選択中にエラーが発生しました: %w", err)
	}
	if dir == "" {
		return "", fmt.Errorf("ディレクトリが選択されませんでした")
	}
	return dir, nil
}

// AddStampToPDF adds a stamp image to a PDF file
func (a *App) AddStampToPDF(stampImagePath, inputPDFPath, outputName string) error {
	position := "tr"
	x := -52
	y := -37
	scale := "0.06"
	opacity := "0.5"
	rotation := "0"

	wm, err := api.ImageWatermark(
		stampImagePath,
		fmt.Sprintf("scale:%s, position:%s, offset: %d %d, opacity:%s, rotation: %s", scale, position, x, y, opacity, rotation),
		true,
		true,
		types.POINTS)
	if err != nil {
		return err
	}

	// 全角スペースを半角スペース削除
	re := regexp.MustCompile(`[\s\p{Zs}]+`)
	trimmedOutputName := re.ReplaceAllString(outputName, "")
	outputPDFPath := trimmedOutputName + ".pdf"

	err = api.AddWatermarksFile(inputPDFPath, outputPDFPath, nil, wm, nil)
	if err != nil {
		return err
	}
	return nil

}
