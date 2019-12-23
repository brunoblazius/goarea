package goarea

import "math"

// Pi proporcao numerica
const Pi = 3.1416

// Circ calcula a circuferencia.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao e visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
