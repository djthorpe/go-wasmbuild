package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	// notifPreviewStyle wraps notification example sections.
	notifPreviewStyle = "display:flex;flex-direction:column;gap:var(--cds-spacing-04,0.75rem);padding:var(--cds-spacing-04,0.75rem);border:1px solid var(--cds-border-subtle-01,#e0e0e0);"
)

func NotificationExamples() mvc.View {
	return cds.Section(
		cds.LeadPara(
			`Notifications use the `, cds.Code("cds-inline-notification"), ` and `,
			cds.Code("cds-toast-notification"), ` web components. `,
			`Set the semantic style with `, cds.Code("cds.WithNotificationKind()"), `. `,
			`Suppress the close button with `, cds.Code("cds.WithNotificationHideClose()"), `, `,
			`or auto-dismiss toasts with `, cds.Code("cds.WithNotificationTimeout(ms)"), `.`,
		),
		ExampleRow("Inline — Kinds", Example_Notification_001, "All four semantic styles: info, success, warning, and error."),
		ExampleRow("Inline — Low Contrast", Example_Notification_002, "Low-contrast variant for less-prominent messaging."),
		ExampleRow("Inline — Interactive", Example_Notification_003, "Trigger each kind with a button; dismiss via the × button or wait for auto-close."),
		ExampleRow("Toast — Interactive", Example_Notification_004, "Click a button to fire a real toast into a fixed screen overlay. Each toast auto-dismisses after 5 seconds, or click × to dismiss immediately."),
	)
}

func Example_Notification_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", notifPreviewStyle),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationInfo),
			cds.WithNotificationTitle("Info:"),
			"This is an informational message.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationSuccess),
			cds.WithNotificationTitle("Success:"),
			"Your changes have been saved.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationWarning),
			cds.WithNotificationTitle("Warning:"),
			"Some items could not be updated.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationError),
			cds.WithNotificationTitle("Error:"),
			"The operation failed. Please try again.",
		),
	), sourcecode()
}

func Example_Notification_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", notifPreviewStyle),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationInfo),
			cds.WithNotificationTitle("Info:"),
			cds.WithNotificationLowContrast(),
			"A low-contrast informational notice.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationSuccess),
			cds.WithNotificationTitle("Success:"),
			cds.WithNotificationLowContrast(),
			"A low-contrast success message.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationWarning),
			cds.WithNotificationTitle("Warning:"),
			cds.WithNotificationLowContrast(),
			"A low-contrast warning message.",
		),
		cds.InlineNotification(
			cds.WithNotificationKind(cds.NotificationError),
			cds.WithNotificationTitle("Error:"),
			cds.WithNotificationLowContrast(),
			"A low-contrast error message.",
		),
	), sourcecode()
}

func Example_Notification_003() (mvc.View, string) {
	bin := cds.Section(mvc.WithAttr("style", notifPreviewStyle))
	trigger := cds.Section(
		mvc.WithAttr("style", btnPreviewStyle),
		cds.Button("Info", mvc.WithAttr("id", "notif-info"), cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button("Success", mvc.WithAttr("id", "notif-success"), cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button("Warning", mvc.WithAttr("id", "notif-warning"), cds.WithButtonKind(cds.ButtonSecondary)),
		cds.Button("Error", mvc.WithAttr("id", "notif-error"), cds.WithButtonKind(cds.ButtonSecondary)),
	)
	trigger.AddEventListener("click", func(e dom.Event) {
		v := mvc.ViewFromEvent(e)
		if v == nil || v.Name() != cds.ViewButton {
			return
		}
		var kind cds.NotificationKind
		var title string
		switch v.ID() {
		case "notif-info":
			kind, title = cds.NotificationInfo, "Info:"
		case "notif-success":
			kind, title = cds.NotificationSuccess, "Success:"
		case "notif-warning":
			kind, title = cds.NotificationWarning, "Warning:"
		case "notif-error":
			kind, title = cds.NotificationError, "Error:"
		default:
			return
		}
		n := cds.InlineNotification(
			cds.WithNotificationKind(kind),
			cds.WithNotificationTitle(title),
			"Click × to dismiss this notification.",
		)
		bin.Root().AppendChild(n.Root())
	})
	bin.AddEventListener("cds-notification-closed", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			el.Remove()
		}
	})
	return cds.Section(trigger, bin), sourcecode()
}

// toastStack is the fixed-position overlay into which toast notifications are appended.
// Created lazily the first time Example_Notification_004 fires a toast; reused across
// route navigations for the lifetime of the WASM process.
var toastStack mvc.View

func getToastStack() mvc.View {
	if toastStack == nil {
		toastStack = mvc.New(
			mvc.WithAttr("id", "toast-stack"),
			mvc.WithAttr("style", "position:fixed;top:4.5rem;right:1rem;z-index:9999;"+
				"display:flex;flex-direction:column;gap:var(--cds-spacing-03,0.5rem);"),
		)
		// Remove each toast from the DOM when it closes (button or timeout).
		toastStack.AddEventListener("cds-notification-closed", func(e dom.Event) {
			if el, ok := e.Target().(dom.Element); ok {
				el.Remove()
			}
		})
	}
	return toastStack
}

func Example_Notification_004() (mvc.View, string) {
	type toastSpec struct {
		id    string
		kind  cds.NotificationKind
		title string
		sub   string
	}
	specs := []toastSpec{
		{"toast-info", cds.NotificationInfo, "Informational", "This is an informational toast message."},
		{"toast-success", cds.NotificationSuccess, "Changes saved", "Your changes have been saved successfully."},
		{"toast-warning", cds.NotificationWarning, "Approaching limit", "You have used 80\u0025 of your quota."},
		{"toast-error", cds.NotificationError, "Something went wrong", "Unable to connect. Please try again."},
	}

	btns := make([]any, 0, len(specs)+1)
	btns = append(btns, mvc.WithAttr("style", btnPreviewStyle))
	for _, s := range specs {
		btns = append(btns, cds.Button(s.title,
			mvc.WithAttr("id", s.id),
			cds.WithButtonKind(cds.ButtonSecondary),
		))
	}
	trigger := cds.Section(btns...)
	trigger.AddEventListener("click", func(e dom.Event) {
		v := mvc.ViewFromEvent(e)
		if v == nil || v.Name() != cds.ViewButton {
			return
		}
		for _, s := range specs {
			if v.ID() == s.id {
				getToastStack().Root().AppendChild(
					cds.ToastNotification(
						cds.WithNotificationKind(s.kind),
						cds.WithNotificationTitle(s.title),
						cds.WithNotificationSubtitle(s.sub),
						cds.WithNotificationTimeout(5000),
					).Root(),
				)
				break
			}
		}
	})
	return trigger, sourcecode()
}
