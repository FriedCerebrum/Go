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

// Функция для преобразования изображения в оттенки серого
func filter(img draw.RGBA64Image, wg *sync.WaitGroup, startY, endY, width int) {
	defer wg.Done()

	for y := startY; y < endY; y++ {
		for x := 0; x < width; x++ {
			// Получаем цвет пикселя
			rgba := img.RGBA64At(x, y)

			// Преобразуем в оттенки серого
			gray := uint16((uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3)

			// Применяем новый цвет
			img.SetRGBA64(x, y, color.RGBA64{R: gray, G: gray, B: gray, A: rgba.A})
		}
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
	width, height := bounds.Max.X, bounds.Max.Y

	var wg sync.WaitGroup

	numGoroutines := 16
	rowsPerGoroutine := height / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		startY := i * rowsPerGoroutine
		endY := startY + rowsPerGoroutine
		if i == numGoroutines-1 {
			endY = height // Обработка последней части
		}

		wg.Add(1)
		go filter(drawImg, &wg, startY, endY, width)
	}

	wg.Wait()

	fmt.Println("Время обработки:", time.Since(startTime))

	outputFile, err := os.Create("output.png")
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

	fmt.Println("Изображение успешно обработано и сохранено в output.png")
}
