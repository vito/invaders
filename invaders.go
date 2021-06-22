package invaders

import (
	"bytes"
	"fmt"
	"math/rand"
)

const Height = 11
const Width = 11
const Middle = Width / 2

type Shade int

const (
	Background Shade = iota
	Shade1
	Shade2
	Shade3
	Shade4
	Shade5
	Shade6
	Shade7
)

func (shade Shade) String() string {
	switch shade {
	case Background:
		return "  "
	default:
		return fmt.Sprintf("\x1b[3%dm██\x1b[0m", shade)
	}
}

type Invader [Height]Row

type Row [Width]Shade

func (row Row) IsEmpty() bool {
	for _, s := range row {
		if s != Background {
			return false
		}
	}

	return true
}

func (invader Invader) String() string {
	buf := new(bytes.Buffer)

	for _, row := range invader {
		for _, shade := range row {
			fmt.Fprint(buf, shade)
		}

		fmt.Fprintln(buf)
	}

	return buf.String()
}

func (invader *Invader) Set(r *rand.Rand) {
	invader.build(r, Antennas, Bodies, Arms, Legs, Eyes)
	invader.center()
}

func (invader *Invader) build(r *rand.Rand, parts ...[]Mask) {
	mask := Mask{}

	for row := range invader {
		invader[row] = Row{}
	}

	for _, masks := range parts {
		mask.Overlay(masks[r.Intn(len(masks))])
	}

	for row := range mask {
		for col := range mask[row] {
			distanceH := Middle - col
			if distanceH < 0 {
				distanceH = -distanceH
			}

			if col > Middle {
				invader[row][col] = invader[row][Middle-distanceH]
				continue
			}

			if mask[row][col] == 1 {
				invader[row][col] = Shade(r.Intn(7) + 1)
			}
		}
	}
}

func (invader *Invader) center() {
	empties := 0
	rows := []Row{}
	for _, row := range invader {
		if row.IsEmpty() {
			empties++
		} else {
			rows = append(rows, row)
		}
	}

	if empties%2 == 1 {
		empties++
	}

	spacing := empties / 2

	for i := 0; i < spacing; i++ {
		invader[i] = Row{}
	}

	copy(invader[spacing:], rows)

	for i := spacing + len(rows); i < Height; i++ {
		invader[i] = Row{}
	}
}
