package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type notification struct {
	mvc.View
}

var _ mvc.View = (*notification)(nil)

// NotificationKind controls the semantic colour and icon of a notification.
type NotificationKind string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	NotificationInfo    NotificationKind = "info"
	NotificationSuccess NotificationKind = "success"
	NotificationWarning NotificationKind = "warning"
	NotificationError   NotificationKind = "error"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewNotification, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(notification), element, func(self, child mvc.View) {
			self.(*notification).View = child
		})
	})
}

// InlineNotification renders a <cds-inline-notification> embedded in the page
// flow. Pass WithNotificationKind, WithNotificationTitle, and body text as args.
//
//	cds.InlineNotification("Changes saved.", cds.WithNotificationKind(cds.NotificationSuccess), cds.WithNotificationTitle("Success"))
func InlineNotification(args ...any) *notification {
	return mvc.NewView(new(notification), ViewNotification, "cds-inline-notification", func(self, child mvc.View) {
		self.(*notification).View = child
	}, args).(*notification)
}

// ToastNotification renders a <cds-toast-notification> overlay. Position it
// via a fixed-position container in the markup. Pass WithNotificationKind,
// WithNotificationTitle, and WithNotificationSubtitle as args.
//
//	cds.ToastNotification(cds.WithNotificationKind(cds.NotificationSuccess), cds.WithNotificationTitle("Saved"), cds.WithNotificationSubtitle("Your changes have been saved."))
func ToastNotification(args ...any) *notification {
	return mvc.NewView(new(notification), ViewNotification, "cds-toast-notification", func(self, child mvc.View) {
		self.(*notification).View = child
	}, args).(*notification)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithNotificationKind sets the semantic style (info/success/warning/error).
func WithNotificationKind(k NotificationKind) mvc.Opt {
	return mvc.WithAttr("kind", string(k))
}

// WithNotificationTitle sets the bold title shown in the notification.
func WithNotificationTitle(title string) mvc.Opt {
	return mvc.WithAttr("title", title)
}

// WithNotificationSubtitle sets the secondary text beneath the title.
// Used primarily by ToastNotification.
func WithNotificationSubtitle(subtitle string) mvc.Opt {
	return mvc.WithAttr("subtitle", subtitle)
}

// WithNotificationHideClose hides the × dismiss button.
func WithNotificationHideClose() mvc.Opt {
	return mvc.WithAttr("hide-close-button", "")
}

// WithNotificationLowContrast renders the low-contrast variant of the notification.
func WithNotificationLowContrast() mvc.Opt {
	return mvc.WithAttr("low-contrast", "")
}

// WithNotificationTimeout sets an auto-dismiss timeout in milliseconds.
// Only applicable to ToastNotification.
func WithNotificationTimeout(ms int) mvc.Opt {
	return mvc.WithAttr("timeout", fmt.Sprintf("%d", ms))
}
