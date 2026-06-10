// Пакет weightconv выполняет преобразование весов
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	KgPerP = 0.45359237
)

func (kg Kilogram) String() string {
	return fmt.Sprintf("%g kg", kg)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g p", p)
}
