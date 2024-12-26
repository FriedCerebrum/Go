package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

// Функция для преобразования строки в оттенки серого
func filterParallel(img draw.RGBA64Image, wg *sync.WaitGroup, rowIndex int) {
	defer wg.Done() // Уменьшаем счётчик после завершения горутины

	bounds := img.Bounds()
	width := bounds.Max.X

	for x := 0; x < width; x++ {

		rgba := img.RGBA64At(x, rowIndex)

		gray := uint16((uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3)

		img.SetRGBA64(x, rowIndex, color.RGBA64{R: gray, G: gray, B: gray, A: rgba.A})
	}
}

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}
	fmt.Println("Текущая рабочая директория:", dir)

	file, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	//draw.RGBA64Image
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Ошибка: изображение не может быть преобразовано в draw.RGBA64Image")
		return
	}

	startTime := time.Now()

	bounds := drawImg.Bounds()
	height := bounds.Max.Y

	var wg sync.WaitGroup

	for y := 0; y < height; y++ {
		wg.Add(1)
		go filterParallel(drawImg, &wg, y)
	}

	wg.Wait()

	fmt.Println("Время обработки с параллельными горутинами:", time.Since(startTime))

	outputFile, err := os.Create("output_parallel.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла для сохранения:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Изображение успешно обработано и сохранено в output_parallel.png")
}
