package text

var space = []byte(" ")

type Segment struct {
	Start int

	Stop int

	Padding int

	ForceNewline bool
}

func NewSegment(start, stop int) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func NewSegmentPadding(start, stop, n int) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func (t *Segment) Value(buffer []byte) []byte { _ = "STUB: not implemented"; return nil }

func (t *Segment) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *Segment) Between(other Segment) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func (t *Segment) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (t *Segment) TrimRightSpace(buffer []byte) Segment {
	_ = "STUB: not implemented"
	return *new(Segment)
}

func (t *Segment) TrimLeftSpace(buffer []byte) Segment {
	_ = "STUB: not implemented"
	return *new(Segment)
}

func (t *Segment) TrimLeftSpaceWidth(width int, buffer []byte) Segment {
	_ = "STUB: not implemented"
	return *new(Segment)
}

func (t *Segment) WithStart(v int) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func (t *Segment) WithStop(v int) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func (t *Segment) ConcatPadding(v []byte) []byte { _ = "STUB: not implemented"; return nil }

type Segments struct {
	values []Segment
}

func NewSegments() *Segments { _ = "STUB: not implemented"; return nil }

func (s *Segments) Append(t Segment) { _ = "STUB: not implemented"; return }

func (s *Segments) AppendAll(t []Segment) { _ = "STUB: not implemented"; return }

func (s *Segments) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *Segments) At(i int) Segment { _ = "STUB: not implemented"; return *new(Segment) }

func (s *Segments) Set(i int, v Segment) { _ = "STUB: not implemented"; return }

func (s *Segments) SetSliced(lo, hi int) { _ = "STUB: not implemented"; return }

func (s *Segments) Sliced(lo, hi int) []Segment { _ = "STUB: not implemented"; return nil }

func (s *Segments) Clear() { _ = "STUB: not implemented"; return }

func (s *Segments) Unshift(v Segment) { _ = "STUB: not implemented"; return }

func (s *Segments) Value(buffer []byte) []byte { _ = "STUB: not implemented"; return nil }
