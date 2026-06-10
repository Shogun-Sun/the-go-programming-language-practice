// Пакет lengthconv выполняет преобразование длин.

package lengthconv

import "fmt"

type Meter float64
type Foot float64

const (
	MetersPerFoot = 0.3048
)

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (ft Foot) String() string {
	return fmt.Sprintf("%g ft", ft)
}
