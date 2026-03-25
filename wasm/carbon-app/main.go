package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	button "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/button"
	content "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/content"
	data "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/data"
	form "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/form"
	navigation "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/navigation"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

type docsPanelController struct {
	panel   mvc.View
	visible bool
}

func (c *docsPanelController) Visible() bool {
	return c.visible
}

func (c *docsPanelController) SetVisible(visible bool) mvc.View {
	c.visible = visible
	style := "position:fixed;top:3rem;right:0;height:calc(100vh - 3rem);overflow:auto;z-index:9000"
	if visible {
		style += ";inline-size:36rem;background:var(--cds-layer,#fff);border-inline-start:1px solid var(--cds-border-subtle,#e0e0e0)"
	}
	c.panel.Root().SetAttribute("style", style)
	if panel, ok := c.panel.(mvc.VisibleState); ok {
		panel.SetVisible(visible)
	}
	return c.panel
}

func main() {
	// Side bar navigation
	SideNav := carbon.SideNav(
		carbon.SideNavGroup("Content",
			carbon.SideNavGroupItem("#heading", "Head"),
			carbon.SideNavGroupItem("#para", "Para"),
			carbon.SideNavGroupItem("#lead", "Lead"),
			carbon.SideNavGroupItem("#compact", "Compact"),
			carbon.SideNavGroupItem("#blockquote", "Blockquote"),
			carbon.SideNavGroupItem("#link", "Link"),
			carbon.SideNavGroupItem("#deleted", "Deleted"),
			carbon.SideNavGroupItem("#highlighted", "Highlighted"),
			carbon.SideNavGroupItem("#strong", "Strong"),
			carbon.SideNavGroupItem("#smaller", "Smaller"),
			carbon.SideNavGroupItem("#em", "Em"),
			carbon.SideNavGroupItem("#markdown", "Markdown"),
			carbon.SideNavGroupItem("#code", "Code"),
			carbon.SideNavGroupItem("#codeblock", "CodeBlock"),
			carbon.SideNavGroupItem("#icon", "Icon"),
			carbon.SideNavGroupItem("#grid", "Grid"),
			carbon.SideNavGroupItem("#section", "Section"),
			carbon.SideNavGroupItem("#page", "Page"),
			carbon.SideNavGroupItem("#list", "List"),
			carbon.SideNavGroupItem("#structuredlist", "StructuredList"),
		),
		carbon.SideNavGroup("Components",
			carbon.SideNavGroupItem("#button", "Button"),
			carbon.SideNavGroupItem("#closebutton", "CloseButton"),
			carbon.SideNavGroupItem("#buttongroup", "ButtonGroup"),
			carbon.SideNavGroupItem("#tile", "Tiles"),
			carbon.SideNavGroupItem("#tag", "Tags"),
		),
		carbon.SideNavGroup("Forms",
			carbon.SideNavGroupItem("#form-introduction", "Introduction"),
			carbon.SideNavGroupItem("#input", "Input"),
			carbon.SideNavGroupItem("#checkbox", "Checkboxes"),
			carbon.SideNavGroupItem("#dropdown", "Dropdowns"),
		),
		carbon.SideNavGroup("Data",
			carbon.SideNavGroupItem("#table", "Table"),
			carbon.SideNavGroupItem("#pagination", "Pagination"),
		),
		carbon.SideNavGroup("Navigation",
			carbon.SideNavGroupItem("#header", "Header"),
			carbon.SideNavGroupItem("#sidenav", "SideNav"),
			carbon.SideNavGroupItem("#headerpanel", "HeaderPanel"),
		),
	)

	docsBody := carbon.Section(
		mvc.WithStyle("display:grid;gap:1rem"),
		carbon.Markdown("Select a story's top-right link to inspect its component documentation."),
	)
	closeBtn := carbon.CloseButton(mvc.WithStyle("position:absolute;top:0;right:0"))
	docsPanel := carbon.HeaderPanel(
		carbon.With(carbon.ThemeG10),
		mvc.HTML("DIV", mvc.WithStyle("position:relative;padding:1rem 1.5rem;display:grid;gap:1rem"),
			closeBtn,
			docsBody,
		),
	)
	docsPanelState := &docsPanelController{panel: docsPanel}
	closeBtn.AddEventListener(carbon.EventClick, func(dom.Event) {
		docsPanelState.SetVisible(false)
	})
	docsPanelState.SetVisible(false)
	storybook.SetDocsPanel(docsPanelState, docsBody)

	// Main content area — min-height fills the viewport below the fixed header.
	Content := carbon.Section(
		mvc.WithStyle("min-height:100vh"),
		router(SideNav),
	)

	mvc.New(carbon.Section(
		carbon.Header(
			carbon.HeaderNavGlobal(
				carbon.Button(
					carbon.Icon(carbon.IconLogoGithub, carbon.With(carbon.IconSize24)),
					mvc.WithAttr("href", "https://github.com/djthorpe/go-wasmbuild"),
					mvc.WithAttr("target", "_blank"),
					mvc.WithAriaLabel("GitHub repository"),
					mvc.WithStyle("color:white;padding:0 0.75rem;display:flex;align-items:center"),
				),
			),
		).SetLabel("#", "Carbon Design System", "Storybook"),
		SideNav,
		Content,
		docsPanel,
	), carbon.With(carbon.ThemeG90)).Run()
}

func router(nav mvc.View) mvc.View {
	type ItemSelector interface {
		Item(string) mvc.View
	}
	return mvc.Router().
		Active(nav.(mvc.ActiveGroup)).
		Page("#heading", carbon.Page(content.HeadingView()...), nav.(ItemSelector).Item("#heading")).
		Page("#para", carbon.Page(content.ParaView()...), nav.(ItemSelector).Item("#para")).
		Page("#lead", carbon.Page(content.LeadView()...), nav.(ItemSelector).Item("#lead")).
		Page("#compact", carbon.Page(content.CompactView()...), nav.(ItemSelector).Item("#compact")).
		Page("#blockquote", carbon.Page(content.BlockquoteView()...), nav.(ItemSelector).Item("#blockquote")).
		Page("#link", carbon.Page(content.LinkView()...), nav.(ItemSelector).Item("#link")).
		Page("#deleted", carbon.Page(content.DeletedView()...), nav.(ItemSelector).Item("#deleted")).
		Page("#highlighted", carbon.Page(content.HighlightedView()...), nav.(ItemSelector).Item("#highlighted")).
		Page("#strong", carbon.Page(content.StrongView()...), nav.(ItemSelector).Item("#strong")).
		Page("#smaller", carbon.Page(content.SmallerView()...), nav.(ItemSelector).Item("#smaller")).
		Page("#em", carbon.Page(content.EmView()...), nav.(ItemSelector).Item("#em")).
		Page("#markdown", carbon.Page(content.MarkdownView()...), nav.(ItemSelector).Item("#markdown")).
		Page("#code", carbon.Page(content.CodeView()...), nav.(ItemSelector).Item("#code")).
		Page("#codeblock", carbon.Page(content.CodeBlockView()...), nav.(ItemSelector).Item("#codeblock")).
		Page("#icon", carbon.Page(content.IconView()...), nav.(ItemSelector).Item("#icon")).
		Page("#grid", carbon.Page(content.GridView()...), nav.(ItemSelector).Item("#grid")).
		Page("#section", carbon.Page(content.SectionView()...), nav.(ItemSelector).Item("#section")).
		Page("#page", carbon.Page(content.PageView()...), nav.(ItemSelector).Item("#page")).
		Page("#list", carbon.Page(content.ListView()...), nav.(ItemSelector).Item("#list")).
		Page("#button", carbon.Page(button.View()...), nav.(ItemSelector).Item("#button")).
		Page("#closebutton", carbon.Page(button.CloseButtonView()...), nav.(ItemSelector).Item("#closebutton")).
		Page("#buttongroup", carbon.Page(button.GroupView()...), nav.(ItemSelector).Item("#buttongroup")).
		Page("#tile", carbon.Page(content.TileView()...), nav.(ItemSelector).Item("#tile")).
		Page("#tag", carbon.Page(content.TagView()...), nav.(ItemSelector).Item("#tag")).
		Page("#form-introduction", carbon.Page(form.IntroductionView()...), nav.(ItemSelector).Item("#form-introduction")).
		Page("#input", carbon.Page(form.InputView()...), nav.(ItemSelector).Item("#input")).
		Page("#checkbox", carbon.Page(form.CheckboxView()...), nav.(ItemSelector).Item("#checkbox")).
		Page("#dropdown", carbon.Page(form.DropdownView()...), nav.(ItemSelector).Item("#dropdown")).
		Page("#structuredlist", carbon.Page(content.StructuredListView()...), nav.(ItemSelector).Item("#structuredlist")).
		Page("#table", carbon.Page(data.TableView()...), nav.(ItemSelector).Item("#table")).
		Page("#pagination", carbon.Page(data.PaginationView()...), nav.(ItemSelector).Item("#pagination")).
		Page("#header", carbon.Page(navigation.HeaderView()...), nav.(ItemSelector).Item("#header")).
		Page("#sidenav", carbon.Page(navigation.SideNavView()...), nav.(ItemSelector).Item("#sidenav")).
		Page("#headerpanel", carbon.Page(navigation.PanelView()...), nav.(ItemSelector).Item("#headerpanel"))
}
