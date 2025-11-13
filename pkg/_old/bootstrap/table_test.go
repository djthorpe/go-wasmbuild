package bootstrap

import (
	"strings"
	"testing"

	. "github.com/djthorpe/go-wasmbuild"
)

func TestTable_Basic(t *testing.T) {
	table := Table()

	if table == nil {
		t.Fatal("Table() returned nil")
	}

	// Check element type
	if tag := table.Element().TagName(); tag != "TABLE" {
		t.Errorf("Expected tag name 'TABLE', got '%s'", tag)
	}

	// Check table class
	if class := table.Element().GetAttribute("class"); !strings.Contains(class, "table") {
		t.Errorf("Expected class to contain 'table', got '%s'", class)
	}

	// Check that only TBODY is created by default (THEAD is optional)
	children := table.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (TBODY), got %d", len(children))
	}

	if len(children) >= 1 {
		if children[0].(Element).TagName() != "TBODY" {
			t.Errorf("Expected first child to be TBODY, got %s", children[0].(Element).TagName())
		}
	}
}

func TestTable_Header(t *testing.T) {
	table := Table().Header("Name", "Age", "City")

	// Check that THEAD was created
	children := table.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (THEAD, TBODY), got %d", len(children))
	}

	// Get THEAD element (should be first child)
	thead := children[0].(Element)
	if thead.TagName() != "THEAD" {
		t.Errorf("Expected first child to be THEAD, got %s", thead.TagName())
	}

	// Check that there's a TR in THEAD
	theadChildren := thead.ChildNodes()
	if len(theadChildren) != 1 {
		t.Errorf("Expected 1 TR in THEAD, got %d", len(theadChildren))
	}

	// Check TH cells
	tr := theadChildren[0].(Element)
	thCells := tr.ChildNodes()
	if len(thCells) != 3 {
		t.Errorf("Expected 3 TH cells, got %d", len(thCells))
	}

	expectedHeaders := []string{"Name", "Age", "City"}
	for i, cell := range thCells {
		if cell.(Element).TagName() != "TH" {
			t.Errorf("Expected TH element, got %s", cell.(Element).TagName())
		}
		if cell.TextContent() != expectedHeaders[i] {
			t.Errorf("Expected header '%s', got '%s'", expectedHeaders[i], cell.TextContent())
		}
	}
}

func TestTable_Footer(t *testing.T) {
	table := Table().Footer("Total", "100", "")

	// Check that TFOOT was created (without THEAD, should have TBODY and TFOOT)
	children := table.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (TBODY, TFOOT), got %d", len(children))
	}

	// Get TFOOT element (should be last child)
	tfoot := children[len(children)-1].(Element)
	if tfoot.TagName() != "TFOOT" {
		t.Errorf("Expected TFOOT element, got %s", tfoot.TagName())
	}

	// Check TR in TFOOT
	tfootChildren := tfoot.ChildNodes()
	if len(tfootChildren) != 1 {
		t.Errorf("Expected 1 TR in TFOOT, got %d", len(tfootChildren))
	}

	// Check TD cells
	tr := tfootChildren[0].(Element)
	tdCells := tr.ChildNodes()
	if len(tdCells) != 3 {
		t.Errorf("Expected 3 TD cells, got %d", len(tdCells))
	}

	expectedFooter := []string{"Total", "100", ""}
	for i, cell := range tdCells {
		if cell.(Element).TagName() != "TD" {
			t.Errorf("Expected TD element, got %s", cell.(Element).TagName())
		}
		if cell.TextContent() != expectedFooter[i] {
			t.Errorf("Expected footer '%s', got '%s'", expectedFooter[i], cell.TextContent())
		}
	}
}

func TestTable_WithOptions(t *testing.T) {
	table := Table(WithClass("table-striped"), WithClass("table-hover"))

	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table") {
		t.Errorf("Expected 'table' in class, got '%s'", class)
	}
	if !strings.Contains(class, "table-striped") {
		t.Errorf("Expected 'table-striped' in class, got '%s'", class)
	}
	if !strings.Contains(class, "table-hover") {
		t.Errorf("Expected 'table-hover' in class, got '%s'", class)
	}
}

func TestTable_Caption(t *testing.T) {
	table := Table().Caption("Employee List")

	// Check that CAPTION was created as first child
	children := table.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (CAPTION, TBODY), got %d", len(children))
	}

	// Get CAPTION element (should be first child)
	caption := children[0].(Element)
	if caption.TagName() != "CAPTION" {
		t.Errorf("Expected first child to be CAPTION, got %s", caption.TagName())
	}

	// Check caption text
	if caption.TextContent() != "Employee List" {
		t.Errorf("Expected caption 'Employee List', got '%s'", caption.TextContent())
	}
}

func TestTable_CaptionWithHeader(t *testing.T) {
	table := Table().Caption("Employee List").Header("Name", "Age")

	// Check that both CAPTION and THEAD are present in correct order
	children := table.Element().ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children (CAPTION, THEAD, TBODY), got %d", len(children))
	}

	// Verify order: CAPTION, THEAD, TBODY
	if children[0].(Element).TagName() != "CAPTION" {
		t.Errorf("Expected first child to be CAPTION, got %s", children[0].(Element).TagName())
	}
	if children[1].(Element).TagName() != "THEAD" {
		t.Errorf("Expected second child to be THEAD, got %s", children[1].(Element).TagName())
	}
	if children[2].(Element).TagName() != "TBODY" {
		t.Errorf("Expected third child to be TBODY, got %s", children[2].(Element).TagName())
	}
}

func TestTable_Row(t *testing.T) {
	row := Row("John", "30", "NYC")

	// Check element type
	if tag := row.Element().TagName(); tag != "TR" {
		t.Errorf("Expected tag name 'TR', got '%s'", tag)
	}

	// Check TD cells
	cells := row.Element().ChildNodes()
	if len(cells) != 3 {
		t.Errorf("Expected 3 TD cells, got %d", len(cells))
	}

	expectedContent := []string{"John", "30", "NYC"}
	for i, cell := range cells {
		if cell.(Element).TagName() != "TD" {
			t.Errorf("Expected TD element, got %s", cell.(Element).TagName())
		}
		if cell.TextContent() != expectedContent[i] {
			t.Errorf("Expected content '%s', got '%s'", expectedContent[i], cell.TextContent())
		}
	}
}

func TestTable_WithRows(t *testing.T) {
	table := Table().
		Header("Name", "Age", "City").
		Append(
			Row("John", "30", "NYC"),
			Row("Jane", "25", "LA"),
		)

	// Get TBODY
	children := table.Element().ChildNodes()
	tbody := children[len(children)-1].(Element)

	// Check that TBODY has 2 rows
	rows := tbody.ChildNodes()
	if len(rows) != 2 {
		t.Errorf("Expected 2 rows in TBODY, got %d", len(rows))
	}

	// Check first row
	firstRow := rows[0].(Element)
	if firstRow.TagName() != "TR" {
		t.Errorf("Expected TR element, got %s", firstRow.TagName())
	}

	firstCells := firstRow.ChildNodes()
	if len(firstCells) != 3 {
		t.Errorf("Expected 3 cells in first row, got %d", len(firstCells))
	}

	if firstCells[0].TextContent() != "John" {
		t.Errorf("Expected 'John', got '%s'", firstCells[0].TextContent())
	}

	// Check second row
	secondRow := rows[1].(Element)
	secondCells := secondRow.ChildNodes()
	if secondCells[0].TextContent() != "Jane" {
		t.Errorf("Expected 'Jane', got '%s'", secondCells[0].TextContent())
	}
}

func TestTable_AppendInvalidComponent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when appending non-Row component")
		}
	}()

	// Try to append a Badge (which is not a Row/TR element)
	badge := Badge()
	badge.Append("test")
	Table().Append(badge)
}

func TestTable_InsertInvalidComponent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when inserting non-Row component")
		}
	}()

	// Try to insert a Button (which is not a Row/TR element)
	button := Button(PRIMARY)
	button.Append("test")
	Table().Insert(button)
}

func TestTable_AppendNonComponent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when appending non-Component")
		}
	}()

	// Try to append a string (not a Component)
	Table().Append("test")
}

func TestTable_Count(t *testing.T) {
	table := Table()

	// Empty table should have count 0
	if count := table.Count(); count != 0 {
		t.Errorf("Expected count 0 for empty table, got %d", count)
	}

	// Add some rows
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
	)

	// Should have count 3
	if count := table.Count(); count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}
}

func TestTable_InsertBefore(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
	)

	// Insert before index 1
	table.InsertBefore(1, Row("Alice", "28", "Seattle"))

	// Should now have 4 rows
	if count := table.Count(); count != 4 {
		t.Errorf("Expected count 4, got %d", count)
	}

	// Check that Alice is at index 1
	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	row1 := rows[1].(Element)
	cells := row1.ChildNodes()
	if cells[0].TextContent() != "Alice" {
		t.Errorf("Expected 'Alice' at index 1, got '%s'", cells[0].TextContent())
	}

	// Check that Jane moved to index 2
	row2 := rows[2].(Element)
	cells2 := row2.ChildNodes()
	if cells2[0].TextContent() != "Jane" {
		t.Errorf("Expected 'Jane' at index 2, got '%s'", cells2[0].TextContent())
	}
}

func TestTable_InsertBefore_AtEnd(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// Insert with index -1 should append to end
	table.InsertBefore(-1, Row("Bob", "35", "Chicago"))

	// Should have 3 rows
	if count := table.Count(); count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}

	// Check that Bob is last
	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	lastRow := rows[2].(Element)
	cells := lastRow.ChildNodes()
	if cells[0].TextContent() != "Bob" {
		t.Errorf("Expected 'Bob' at last position, got '%s'", cells[0].TextContent())
	}
}

func TestTable_InsertBefore_OutOfBounds(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// Insert with index 10 (out of bounds) should do nothing
	table.InsertBefore(10, Row("Bob", "35", "Chicago"))

	// Should still have 2 rows
	if count := table.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestTable_Delete(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
	)

	// Delete index 1
	table.Delete(1)

	// Should now have 2 rows
	if count := table.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}

	// Check that Jane is gone
	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	row0 := rows[0].(Element)
	row1 := rows[1].(Element)

	cells0 := row0.ChildNodes()
	cells1 := row1.ChildNodes()

	if cells0[0].TextContent() != "John" {
		t.Errorf("Expected 'John' at index 0, got '%s'", cells0[0].TextContent())
	}
	if cells1[0].TextContent() != "Bob" {
		t.Errorf("Expected 'Bob' at index 1, got '%s'", cells1[0].TextContent())
	}
}

func TestTable_Delete_OutOfBounds(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// Delete index 10 (out of bounds) should do nothing
	table.Delete(10)

	// Should still have 2 rows
	if count := table.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestTable_Replace(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
	)

	// Replace index 1
	table.Replace(1, Row("Alice", "28", "Seattle"))

	// Should still have 3 rows
	if count := table.Count(); count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}

	// Check that Jane was replaced with Alice
	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	row1 := rows[1].(Element)
	cells := row1.ChildNodes()

	if cells[0].TextContent() != "Alice" {
		t.Errorf("Expected 'Alice' at index 1, got '%s'", cells[0].TextContent())
	}
	if cells[1].TextContent() != "28" {
		t.Errorf("Expected '28' at index 1, got '%s'", cells[1].TextContent())
	}
}

func TestTable_Replace_OutOfBounds(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// Replace index 10 (out of bounds) should do nothing
	table.Replace(10, Row("Bob", "35", "Chicago"))

	// Should still have 2 rows
	if count := table.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestTable_Active(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
		Row("Alice", "28", "Seattle"),
	)

	// Mark rows 1 and 3 as active
	result := table.Active(1, 3)

	// Check return value
	if len(result) != 2 || result[0] != 1 || result[1] != 3 {
		t.Errorf("Expected Active to return [1, 3], got %v", result)
	}

	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()

	// Check row 0 - should not be active
	row0 := rows[0].(Element)
	if row0.ClassList().Contains("table-active") {
		t.Error("Row 0 should not be active")
	}

	// Check row 1 - should be active
	row1 := rows[1].(Element)
	if !row1.ClassList().Contains("table-active") {
		t.Error("Row 1 should be active")
	}

	// Check row 2 - should not be active
	row2 := rows[2].(Element)
	if row2.ClassList().Contains("table-active") {
		t.Error("Row 2 should not be active")
	}

	// Check row 3 - should be active
	row3 := rows[3].(Element)
	if !row3.ClassList().Contains("table-active") {
		t.Error("Row 3 should be active")
	}

	// Test get mode - should return [1, 3]
	activeRows := table.Active()
	if len(activeRows) != 2 || activeRows[0] != 1 || activeRows[1] != 3 {
		t.Errorf("Expected Active() to return [1, 3], got %v", activeRows)
	}
}

func TestTable_Active_ClearsPrevious(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
		Row("Bob", "35", "Chicago"),
	)

	// Mark row 1 as active
	table.Active(1)

	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	row1 := rows[1].(Element)

	if !row1.ClassList().Contains("table-active") {
		t.Error("Row 1 should be active")
	}

	// Now mark row 2 as active (should clear row 1)
	table.Active(2)

	// Row 1 should no longer be active
	if row1.ClassList().Contains("table-active") {
		t.Error("Row 1 should no longer be active")
	}

	// Row 2 should be active
	row2 := rows[2].(Element)
	if !row2.ClassList().Contains("table-active") {
		t.Error("Row 2 should be active")
	}
}

func TestTable_Active_OutOfBounds(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// Mark row 10 (out of bounds) as active - should be ignored
	table.Active(10)

	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()

	// No rows should be active
	for i, row := range rows {
		elem := row.(Element)
		if elem.ClassList().Contains("table-active") {
			t.Errorf("Row %d should not be active", i)
		}
	}
}

func TestTable_Active_Empty(t *testing.T) {
	table := Table()
	table.Append(
		Row("John", "30", "NYC"),
		Row("Jane", "25", "LA"),
	)

	// First mark row 0 as active
	table.Active(0)

	tbody := table.Element().FirstChild().(Element)
	rows := tbody.ChildNodes()
	row0 := rows[0].(Element)

	if !row0.ClassList().Contains("table-active") {
		t.Error("Row 0 should be active")
	}

	// Call Active with no indices - should return current active rows [0]
	activeRows := table.Active()
	if len(activeRows) != 1 || activeRows[0] != 0 {
		t.Errorf("Expected Active() to return [0], got %v", activeRows)
	}

	// Row should still be active
	if !row0.ClassList().Contains("table-active") {
		t.Error("Row 0 should still be active")
	}
}
