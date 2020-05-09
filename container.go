package senses

import "github.com/hajimehoshi/ebiten"

type Container struct {
	*Visual
	Theme       *Theme
	Widgets     []Widget
	Orientation Orientation
	GrowAdder   float64
}

func NewContainer(orientation Orientation, theme *Theme) *Container {
	container := new(Container)
	container.Visual = NewVisual(theme)
	container.Widgets = make([]Widget, 0)
	container.Orientation = orientation
	return container
}

func (container *Container) Draw(screen *ebiten.Image) {
	container.Visual.Draw(screen)
	for _, widget := range container.Widgets {
		widget.Draw(screen)
	}
}

func (container *Container) Arrange() {
	container.Visual.Arrange()
	println(container.size.Width, container.size.Height)
	sizes := make([]*Size, 0)
	width := 0
	height := 0
	adder := 0
	for _, widget := range container.Widgets {
		if container.Orientation == Vertical {
			height = int(float64(container.size.Height) * widget.GrowRatio() / container.GrowAdder)
			width = container.size.Width
			sizes = append(sizes, NewSize(width, height))
			adder += height
		} else {
			width = int(float64(container.size.Width) * widget.GrowRatio() / container.GrowAdder)
			height = container.size.Height
			sizes = append(sizes, NewSize(width, height))
			adder += width
		}
	}
	var compensator int
	if container.Orientation == Vertical {
		compensator = container.size.Height - adder
	} else {
		compensator = container.size.Width - adder
	}
	for index := 0; index < compensator; index++ {
		if container.Orientation == Vertical {
			sizes[index].Height++
		} else {
			sizes[index].Width++
		}
	}
	adder = 0
	for index, widget := range container.Widgets {
		if container.Orientation == Vertical {
			widget.Position().Update(container.position.X, container.position.Y+adder)
			widget.Size().UpdateFrom(sizes[index])
			adder += sizes[index].Height
			widget.Arrange()
		} else {
			widget.Position().Update(container.position.X+adder, container.position.Y)
			widget.Size().UpdateFrom(sizes[index])
			adder += sizes[index].Width
			widget.Arrange()
		}
	}

}

func (container *Container) Add(widget Widget) {
	container.GrowAdder += widget.GrowRatio()
	container.Widgets = append(container.Widgets, widget)
	container.Arrange()
}
