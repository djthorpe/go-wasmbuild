package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tabs struct {
	mvc.View
}

type tab struct {
	mvc.View
}

// tabPane is an ephemeral data-holder consumed by TabSet; it is NOT a
// registered view and never appears in the DOM directly.
type tabPane struct {
	label    string
	content  []any
	selected bool
	disabled bool
}

var _ mvc.View = (*tabs)(nil)
var _ mvc.View = (*tab)(nil)

// TabsType controls the visual style of the tab bar.
type TabsType string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	TabsRegular   TabsType = ""          // default line tabs
	TabsContained TabsType = "contained" // box / pill tabs
)

// tabPaneSeq generates unique panel element IDs across the page lifetime.
var tabPaneSeq int

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTabs, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tabs), element, func(self, child mvc.View) {
			self.(*tabs).View = child
		})
	})
	mvc.RegisterView(ViewTab, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tab), element, func(self, child mvc.View) {
			self.(*tab).View = child
		})
	})
}

// Tabs returns a bare <cds-tabs> element for manual composition.
// For most use-cases prefer TabSet, which wires tabs and panels automatically.
func Tabs(args ...any) *tabs {
	return mvc.NewView(new(tabs), ViewTabs, "cds-tabs", func(self, child mvc.View) {
		self.(*tabs).View = child
	}, args).(*tabs)
}

// Tab returns a <cds-tab> element for use inside a bare Tabs container.
func Tab(label string, args ...any) *tab {
	args = append([]any{label, mvc.WithAttr("value", label)}, args...)
	return mvc.NewView(new(tab), ViewTab, "cds-tab", func(self, child mvc.View) {
		self.(*tab).View = child
	}, args).(*tab)
}

// TabPane holds the label and content for one tab+panel pair.
// Pass to TabSet together with any WithTabsType option.
//
//	cds.TabSet(
//	    cds.TabPane("Overview", cds.Para("…")),
//	    cds.TabPane("Details",  cds.Para("…")).Selected(),
//	    cds.TabPane("History",  cds.Para("…")).Disabled(),
//	    cds.WithTabsType(cds.TabsContained),
//	)
func TabPane(label string, content ...any) *tabPane {
	return &tabPane{label: label, content: content}
}

// Selected marks this pane's tab as initially expanded.
func (p *tabPane) Selected() *tabPane { p.selected = true; return p }

// Disabled marks this pane's tab as non-interactive.
func (p *tabPane) Disabled() *tabPane { p.disabled = true; return p }

// TabSet builds a complete, wired tab group from TabPane items and optional
// WithTabsType options. It generates unique panel IDs automatically, pre-hides
// all panels except the selected one, and lets the Carbon WC manage subsequent
// show/hide on every selection change.
//
//	cds.TabSet(
//	    cds.TabPane("First",  cds.Para(lorem)),
//	    cds.TabPane("Second", cds.Para(lorem)).Selected(),
//	    cds.TabPane("Third",  cds.Para(lorem)).Disabled(),
//	)
func TabSet(args ...any) mvc.View {
	var panes []*tabPane
	var tabsOpts []any
	type tabPanel struct {
		label string
		elem  dom.Element
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case *tabPane:
			panes = append(panes, v)
		case mvc.Opt:
			tabsOpts = append(tabsOpts, v)
		}
	}

	// Auto-select the first pane when none is explicitly marked selected.
	selIdx := -1
	for i, p := range panes {
		if p.selected {
			selIdx = i
			break
		}
	}
	if selIdx < 0 && len(panes) > 0 {
		selIdx = 0
	}

	// Build <cds-tab> children and panel <div>s in a single pass.
	// Set value="selectedLabel" on <cds-tabs> so the WC drives initial selection
	// and fires shouldUpdate on each tab, which toggles hidden on target panels.
	if selIdx >= 0 {
		tabsOpts = append(tabsOpts, mvc.WithAttr("value", panes[selIdx].label))
	}
	tabArgs := append([]any{}, tabsOpts...)
	var panels []any
	var panelRefs []tabPanel

	for i, p := range panes {
		tabPaneSeq++
		panelID := fmt.Sprintf("cds-panel-%d", tabPaneSeq)

		// <cds-tab value="label" target="panelID" [disabled]>label</cds-tab>
		// Do NOT set selected="" here — <cds-tabs value="…"> manages selection.
		ta := []any{
			p.label,
			mvc.WithAttr("value", p.label),
			mvc.WithAttr("target", panelID),
		}
		if p.disabled {
			ta = append(ta, mvc.WithAttr("disabled", ""))
		}
		t := mvc.NewView(new(tab), ViewTab, "cds-tab", func(self, child mvc.View) {
			self.(*tab).View = child
		}, ta...)
		tabArgs = append(tabArgs, t)

		// <div id="panelID" [hidden]>…panel content…</div>
		// Only the initially selected panel is visible; the WC toggles hidden on
		// subsequent selections via CDSContentSwitcherItem.shouldUpdate.
		pArgs := []any{
			mvc.WithAttr("id", panelID),
			mvc.WithAttr("style", "padding-top:var(--cds-spacing-05,1rem);"),
		}
		if i != selIdx {
			pArgs = append(pArgs, mvc.WithAttr("hidden", ""))
		}
		pArgs = append(pArgs, p.content...)
		panel := mvc.HTML("div", pArgs...)
		panels = append(panels, panel)
		panelRefs = append(panelRefs, tabPanel{label: p.label, elem: panel})
	}

	tabsEl := mvc.NewView(new(tabs), ViewTabs, "cds-tabs", func(self, child mvc.View) {
		self.(*tabs).View = child
	}, tabArgs...)
	tabsEl.AddEventListener("cds-tabs-selected", func(e dom.Event) {
		el, ok := e.Target().(dom.Element)
		if !ok {
			return
		}
		selected := el.Value()
		for _, panel := range panelRefs {
			if panel.label == selected {
				panel.elem.RemoveAttribute("hidden")
			} else {
				panel.elem.SetAttribute("hidden", "")
			}
		}
	})

	return Section(tabsEl, panels)
}

// TabSetParts is like TabSet but returns the tab-bar view and the panels view
// separately, so they can be placed in different DOM containers — for example
// a dark-themed header (tab bar) and a light body (panels).
func TabSetParts(args ...any) (tabBar, panels mvc.View) {
	var panes []*tabPane
	var tabsOpts []any
	type tabPanel struct {
		label string
		elem  dom.Element
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case *tabPane:
			panes = append(panes, v)
		case mvc.Opt:
			tabsOpts = append(tabsOpts, v)
		}
	}

	selIdx := -1
	for i, p := range panes {
		if p.selected {
			selIdx = i
			break
		}
	}
	if selIdx < 0 && len(panes) > 0 {
		selIdx = 0
	}

	if selIdx >= 0 {
		tabsOpts = append(tabsOpts, mvc.WithAttr("value", panes[selIdx].label))
	}
	tabArgs := append([]any{}, tabsOpts...)
	var panelList []any
	var panelRefs []tabPanel

	for i, p := range panes {
		tabPaneSeq++
		panelID := fmt.Sprintf("cds-panel-%d", tabPaneSeq)

		ta := []any{
			p.label,
			mvc.WithAttr("value", p.label),
			mvc.WithAttr("target", panelID),
		}
		if p.disabled {
			ta = append(ta, mvc.WithAttr("disabled", ""))
		}
		t := mvc.NewView(new(tab), ViewTab, "cds-tab", func(self, child mvc.View) {
			self.(*tab).View = child
		}, ta...)
		tabArgs = append(tabArgs, t)

		pArgs := []any{mvc.WithAttr("id", panelID)}
		if i != selIdx {
			pArgs = append(pArgs, mvc.WithAttr("hidden", ""))
		}
		pArgs = append(pArgs, p.content...)
		panel := mvc.HTML("div", pArgs...)
		panelList = append(panelList, panel)
		panelRefs = append(panelRefs, tabPanel{label: p.label, elem: panel})
	}

	tabsEl := mvc.NewView(new(tabs), ViewTabs, "cds-tabs", func(self, child mvc.View) {
		self.(*tabs).View = child
	}, tabArgs...)
	tabsEl.AddEventListener("cds-tabs-selected", func(e dom.Event) {
		el, ok := e.Target().(dom.Element)
		if !ok {
			return
		}
		selected := el.Value()
		for _, panel := range panelRefs {
			if panel.label == selected {
				panel.elem.RemoveAttribute("hidden")
			} else {
				panel.elem.SetAttribute("hidden", "")
			}
		}
	})

	return tabsEl, Section(panelList)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithTabsType sets the tab bar visual style.
func WithTabsType(t TabsType) mvc.Opt {
	return mvc.WithAttr("type", string(t))
}

// WithTabsSize sets the height of the tab bar. Valid values are "sm", "md", "lg".
func WithTabsSize(size string) mvc.Opt {
	return mvc.WithAttr("size", size)
}
