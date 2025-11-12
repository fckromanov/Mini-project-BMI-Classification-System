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
	switch {
	case imt < 16:
		category = "Выраженный дефицит массы"
		recommendation = "Больше кушайте"
	case imt >= 16 && imt < 18.5:
		category = "Недостаточная масса"
		recommendation = "Больше кушайте"
	case imt >= 18.5 && imt < 25:
		category = "Норма"
		recommendation = "Продолжайте поддерживать здоровый образ жизни!"
	case imt >= 25 && imt < 30:
		category = "Избыточная масса"
		recommendation = "Ходите в зал, следите за питанием"
	case imt >= 30 && imt < 35:
		category = "Ожирение 1 степени"
		recommendation = "Запишитесь к врачу"
	case imt >= 35 && imt < 40:
		category = "Ожирение 2 степени"
		recommendation = "Запишитесь к врачу"
	default:
		category = "Ожирение 3 степени"
		recommendation = "Запишитесь к врачу"
	}
	fmt.Printf("Категория: %s\n", category)
	fmt.Println(recommendation)
	fmt.Println("")
	fmt.Println("Спасибо, что пользуетесь нашим сервисом")
}
