package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func ParaView() []any {
	return []any{
		storybook.PageHeader("Para", "Para.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			paragraphStory(),
		),
	}
}

func LeadView() []any {
	return []any{
		storybook.PageHeader("Lead", "Lead.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			leadStory(),
		),
	}
}

func CompactView() []any {
	return []any{
		storybook.PageHeader("Compact", "Compact.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			compactStory(),
		),
	}
}

func BlockquoteView() []any {
	return []any{
		storybook.PageHeader("Blockquote", "Blockquote.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			blockquoteStory(),
		),
	}
}

func LinkView() []any {
	return []any{
		storybook.PageHeader("Link", "Link.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			inlineLinkStory(),
			linkWithIconStory(),
			iconOnlyLinkStory(),
		),
	}
}

func DeletedView() []any {
	return []any{
		storybook.PageHeader("Deleted", "Deleted.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			deletedStory(),
		),
	}
}

func HighlightedView() []any {
	return []any{
		storybook.PageHeader("Highlighted", "Highlighted.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			highlightedStory(),
		),
	}
}

func StrongView() []any {
	return []any{
		storybook.PageHeader("Strong", "Strong.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			strongStory(),
		),
	}
}

func SmallerView() []any {
	return []any{
		storybook.PageHeader("Smaller", "Smaller.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			smallerStory(),
		),
	}
}

func EmView() []any {
	return []any{
		storybook.PageHeader("Em", "Em.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			emStory(),
		),
	}
}

func paragraphStory() dom.Element {
	const copy = "Carbon body text is used for supporting copy, descriptions, and longer-form content throughout the interface. It should stay readable across themes and layouts, giving product teams a consistent baseline for explanatory content, inline guidance, and general-purpose narrative text that supports the primary task without competing with headings or controls."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Para(copy),
	)

	return storybook.Story(
		"Paragraph",
		"Paragraphs use the Carbon body text token for readable supporting content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func leadStory() dom.Element {
	const copy = "Lead text introduces a page, section, or feature with more visual presence than standard paragraph copy. It works well for summaries, opening statements, and the first block of explanatory content when you want to establish context quickly before the reader moves into denser details or supporting information."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Lead(copy),
	)

	return storybook.Story(
		"Lead",
		"Lead text uses a larger body style for introductory or summary content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func compactStory() dom.Element {
	const copy = "Compact text is useful when space is constrained and the content plays a more supporting role, such as metadata, short descriptions, dense panels, or interface regions where the vertical rhythm needs to stay tight without sacrificing readability."
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Compact(copy),
	)

	return storybook.Story(
		"Compact",
		"Compact text uses a tighter body style for dense supporting content.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func blockquoteStory() dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Blockquote(
			"Good content systems need a clear typographic hierarchy so that long-form material remains readable without losing structure or emphasis.",
		).SetLabel("Carbon text storybook"),
	)

	return storybook.Story(
		"Blockquote",
		"Blockquotes highlight quoted or supporting content with an optional citation label.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func deletedStory() dom.Element {
	return inlineTextStory(
		"Deleted",
		"Deleted text is useful for showing removed or superseded inline content.",
		carbon.Para("Previous wording: ", carbon.Deleted("deprecated")),
	)
}

func highlightedStory() dom.Element {
	return inlineTextStory(
		"Highlighted",
		"Highlighted text draws attention to inline content without changing document flow.",
		carbon.Para("Status: ", carbon.Highlighted("needs review")),
	)
}

func strongStory() dom.Element {
	return inlineTextStory(
		"Strong",
		"Strong emphasis increases visual weight for important inline text.",
		carbon.Para(carbon.Strong("Important:"), " Confirm the deployment window before release."),
	)
}

func smallerStory() dom.Element {
	return inlineTextStory(
		"Smaller",
		"Smaller text is suited to compact supporting details and secondary metadata.",
		carbon.Para("Build 218 ", carbon.Smaller("experimental")),
	)
}

func emStory() dom.Element {
	return inlineTextStory(
		"Em",
		"Emphasis adds inline stress without switching to strong weight.",
		carbon.Para("This setting is ", carbon.Em("recommended"), " for most deployments."),
	)
}

func inlineTextStory(title, description string, sample mvc.View) dom.Element {
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:0.75rem;width:100%;padding:1rem 1rem 0"),
		sample,
	)

	return storybook.Story(
		title,
		description,
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
	)
}

func inlineLinkStory() dom.Element {
	const href = "#link"
	inlineLink := carbon.Link(
		href,
		carbon.With(carbon.LinkInline, carbon.SizeMedium),
		"Read the Carbon content guidelines",
	)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Para(
			"Links support inline navigation and can include a trailing icon for actions such as opening documentation or moving to related content. ",
			inlineLink,
			" within body copy to keep reading flow intact.",
		),
		carbon.Compact("Inline links keep navigation within running body copy without breaking reading flow."),
	)

	return storybook.Story(
		"Inline Link",
		"Inline links stay embedded in paragraph copy and support the same size and enabled state controls as other Carbon links.",
		canvas,
		inlineLink,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", carbon.SizeMedium, []carbon.Attr{carbon.SizeSmall, carbon.SizeMedium, carbon.SizeLarge}, func(size carbon.Attr) {
			inlineLink.Apply(carbon.With(size)...)
		}),
		storybook.CheckboxGroup("State", "Enabled", true, func(enabled bool) {
			inlineLink.SetEnabled(enabled)
		}),
	)
}

func linkWithIconStory() dom.Element {
	const href = "#link"
	icon := carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize20))
	link := carbon.Link(
		href,
		carbon.With(carbon.SizeMedium),
		"Read the Carbon content guidelines",
		icon,
	)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Para(link),
		carbon.Compact("A trailing icon helps communicate navigation intent while preserving the link label as visible text."),
	)

	return storybook.Story(
		"Link With Icon",
		"Standalone links can include a slotted Carbon icon that inherits the link color.",
		canvas,
		link,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", carbon.SizeMedium, []carbon.Attr{carbon.SizeSmall, carbon.SizeMedium, carbon.SizeLarge}, func(size carbon.Attr) {
			link.Apply(carbon.With(size)...)
			iconSize := carbon.IconSize20
			if size == carbon.SizeSmall {
				iconSize = carbon.IconSize16
			}
			icon.Apply(carbon.With(iconSize)...)
		}),
		storybook.CheckboxGroup("State", "Enabled", true, func(enabled bool) {
			link.SetEnabled(enabled)
		}),
	)
}

func iconOnlyLinkStory() dom.Element {
	const href = "#link"
	icon := carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize20))
	link := carbon.Link(
		href,
		carbon.With(carbon.SizeMedium),
		icon,
	)
	link.SetLabel("Open content guidelines")

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem 1rem 0"),
		carbon.Para(link),
		carbon.Compact("Icon-only links require an accessible label because the icon itself is decorative by default."),
	)

	return storybook.Story(
		"Icon-Only Link",
		"Icon-only links use the same component but rely on an accessible label instead of visible text.",
		canvas,
		link,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", carbon.SizeMedium, []carbon.Attr{carbon.SizeSmall, carbon.SizeMedium, carbon.SizeLarge}, func(size carbon.Attr) {
			link.Apply(carbon.With(size)...)
			iconSize := carbon.IconSize20
			if size == carbon.SizeSmall {
				iconSize = carbon.IconSize16
			}
			icon.Apply(carbon.With(iconSize)...)
		}),
		storybook.CheckboxGroup("State", "Enabled", true, func(enabled bool) {
			link.SetEnabled(enabled)
		}),
	)
}
