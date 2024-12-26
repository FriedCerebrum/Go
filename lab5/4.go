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

var gaussianKernel = [][]float64{
	{1, 0, -1},
	{1, 0, -1},
	{1, 0, -1},
}

const kernelSize = 3

func applyConvolution(srcImg *image.RGBA64, dstImg *image.RGBA64, wg *sync.WaitGroup, y int, width, height int) {
	defer wg.Done()

	for x := 0; x < width; x++ {
		var rSum, gSum, bSum, aSum float64
		var weightSum float64

		for ky := 0; ky < kernelSize; ky++ {
			for kx := 0; kx < kernelSize; kx++ {

				px := x + kx - 1
				py := y + ky - 1

				if px >= 0 && px < width && py >= 0 && py < height {

					r, g, b, a := srcImg.RGBA64At(px, py).RGBA()
					weight := gaussianKernel[ky][kx]

					rSum += float64(r) * weight
					gSum += float64(g) * weight
					bSum += float64(b) * weight
					aSum += float64(a) * weight
					weightSum += weight
				}
			}
		}

		r := uint16(rSum)
		g := uint16(gSum)
		b := uint16(bSum)
		a := uint16(aSum)

		// Устанавливаем новый цвет пикселя
		dstImg.SetRGBA64(x, y, color.RGBA64{R: r, G: g, B: b, A: a})
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

	srcRGBA, ok := img.(*image.RGBA64)
	if !ok {

		bounds := img.Bounds()
		srcRGBA = image.NewRGBA64(bounds)
		draw.Draw(srcRGBA, bounds, img, bounds.Min, draw.Src)
	}

	dstRGBA := image.NewRGBA64(srcRGBA.Bounds())

	startTime := time.Now()

	bounds := srcRGBA.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var wg sync.WaitGroup

	numGoroutines := 16
	rowsPerGoroutine := height / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		startY := i * rowsPerGoroutine
		endY := startY + rowsPerGoroutine
		if i == numGoroutines-1 {
			endY = height
		}

		for y := startY; y < endY; y++ {
			wg.Add(1)
			go applyConvolution(srcRGBA, dstRGBA, &wg, y, width, height)
		}
	}

	wg.Wait()

	fmt.Println("Время обработки:", time.Since(startTime))

	// Сохраняем обработанное изображение
	outputFile, err := os.Create("output_convolution.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла для сохранения:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, dstRGBA)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Изображение успешно обработано и сохранено в output_convolution.png")
}
