package main

import "fmt"

func main() {
	name := "Ivan"
	weight := 65
	height := 180

	imt := float64(weight) / (float64(height) / 100.0)

	var category string
	var recommendation string

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     АНАЛИЗ ИНДЕКСА МАССЫ ТЕЛА         ║")
	fmt.Println("╚════════════════════════════════════════╝")

	fmt.Printf("Пациент: %s\n", name)
	fmt.Printf("Рост: %d см\n", height)
	fmt.Printf("Вес: %d кг\n", weight)
	fmt.Println("")
	fmt.Printf("Индекс массы тела: %.2f\n", imt)
	fmt.Println("")
	fmt.Println("✅ Рекомендация:")
	if imt < 16 {
		category = "Выраженный дефицит массы"
	} else if imt >= 16 && imt < 18.5 {
		category = "Недостаточная масса"
	} else if imt >= 18.5 && imt < 25 {
		category = "Норма"
	} else if imt >= 25 && imt < 30 {
		category = "Избыточная масса"
	} else if imt >= 30 && imt < 35 {
		category = "Ожирение 1 степени"
	} else if imt >= 35 && imt < 40 {
		category = "Ожирение 2 степени"
	} else {
		category = "Ожирение 3 степени"
	}
	fmt.Printf("Категория: %s\n", category)
	if imt < 16 {
		recommendation = "Больше кушайте"
	} else if imt >= 16 && imt < 18.5 {
		recommendation = "Больше кушайте"
	} else if imt >= 18.5 && imt < 25 {
		recommendation = "Продолжайте поддерживать здоровый образ жизни!"
	} else if imt >= 25 && imt < 30 {
		recommendation = "Ходите в зал, следите за питанием"
	} else if imt >= 30 && imt < 35 {
		recommendation = "Запишитесь к врачу"
	} else if imt >= 35 && imt < 40 {
		recommendation = "Запишитесь к врачу"
	} else {
		recommendation = "Запишитесь к врачу"
	}
	fmt.Println(recommendation)
	fmt.Println("")
	fmt.Println("Спасибо, что пользуетесь нашим сервисом")
}
