package bootstrap

///////////////////////////////////////////////////////////////////////////////
// TABLE-SPECIFIC OPTIONS

// WithStripedRows adds zebra-striping to table rows within the tbody
func WithStripedRows() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-striped-columns")
		o.classList.Add("table-striped")
		return nil
	}
}

// WithoutStriped removes zebra-striping from table rows
func WithoutStriped() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-striped", "table-striped-columns")
		return nil
	}
}

// WithStripedColumns adds zebra-striping to table columns
func WithStripedColumns() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-striped")
		o.classList.Add("table-striped-columns")
		return nil
	}
}

// WithHover enables a hover state on table rows within tbody
func WithHover() Opt {
	return WithClass("table-hover")
}

// WithoutHover removes the hover state from table rows
func WithoutHover() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-hover")
		return nil
	}
}

// WithActive highlights a table row or cell
func WithActive() Opt {
	return WithClass("table-active")
}

// WithBordered adds borders on all sides of the table and cells
func WithBordered() Opt {
	return WithClass("table-bordered")
}

// WithBorderless removes all borders from the table
func WithBorderless() Opt {
	return WithClass("table-borderless")
}

// WithoutBorder removes both bordered and borderless classes, returning to default border style
func WithoutBorder() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-bordered", "table-borderless")
		return nil
	}
}

// WithGroupDivider adds a thicker border between table groups (thead, tbody, tfoot)
// Adds class: table-group-divider
func WithGroupDivider() Opt {
	return WithClass("table-group-divider")
}

// WithoutGroupDivider removes the thicker border between table groups
func WithoutGroupDivider() Opt {
	return func(o *opts) error {
		o.classList.Remove("table-group-divider")
		return nil
	}
}

// WithAnimation enables animations for row operations (InsertBefore, Delete, Replace)
// Uses scaleY transforms for smooth height transitions
// This adds a special marker class that the Table constructor will detect
func WithAnimation() Opt {
	return WithClass("table-animated")
}
