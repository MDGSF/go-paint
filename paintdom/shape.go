package paintdom

type coord = float64
type ShapeID = string

type Shape interface {
	GetID() ShapeID
}

type shapeBase struct {
	ID ShapeID `json:"id"`
}

func (p *shapeBase) GetID() ShapeID {
	return p.ID
}

// ---------------------------------------------------

type ShapeStyle struct {
	LineWidth coord  `json:"lineWidth"`
	LineColor string `json:"lineColor"`
	FillColor string `json:"fillColor"`
}

type Point struct {
	X coord `json:"x"`
	Y coord `json:"y"`
}

// ---------------------------------------------------

type Path struct {
}

// ---------------------------------------------------
