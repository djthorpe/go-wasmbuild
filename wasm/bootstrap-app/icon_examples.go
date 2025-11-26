package main

import (
	// Packages
	"strings"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func IconExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("icon_examples.md"),
		bs.HRule(),
		AllIcons(),
	)
}

func AllIcons() mvc.View {
	icons := make([]any, 0, len(iconsNames))
	for _, name := range iconsNames {
		name = strings.TrimSuffix(name, ".svg")
		icons = append(icons, bs.Col(
			bs.Container(bs.WithFlex(bs.Middle), mvc.WithClass("m-3", "p-2", "text-center"), bs.WithBorder(),
				bs.Icon(name, mvc.WithClass("fs-1")),
				bs.Smaller(name, mvc.WithClass("text-nowrap", "text-center")),
			),
		))
	}
	return bs.Row(icons...)
}

var (
	iconsNames = []string{
		"0-circle-fill.svg",
		"0-circle.svg",
		"0-square-fill.svg",
		"0-square.svg",
		"1-circle-fill.svg",
		"1-circle.svg",

		"1-square-fill.svg",

		"1-square.svg",

		"123.svg",

		"2-circle-fill.svg",

		"2-circle.svg",

		"2-square-fill.svg",

		"2-square.svg",

		"3-circle-fill.svg",

		"3-circle.svg",

		"3-square-fill.svg",

		"3-square.svg",

		"4-circle-fill.svg",

		"4-circle.svg",

		"4-square-fill.svg",

		"4-square.svg",

		"5-circle-fill.svg",

		"5-circle.svg",

		"5-square-fill.svg",

		"5-square.svg",

		"6-circle-fill.svg",

		"6-circle.svg",

		"6-square-fill.svg",

		"6-square.svg",

		"7-circle-fill.svg",

		"7-circle.svg",

		"7-square-fill.svg",

		"7-square.svg",

		"8-circle-fill.svg",

		"8-circle.svg",

		"8-square-fill.svg",

		"8-square.svg",

		"9-circle-fill.svg",

		"9-circle.svg",

		"9-square-fill.svg",

		"9-square.svg",

		"activity.svg",

		"airplane-engines-fill.svg",

		"airplane-engines.svg",

		"airplane-fill.svg",

		"airplane.svg",

		"alarm-fill.svg",

		"alarm.svg",

		"alexa.svg",

		"align-bottom.svg",

		"align-center.svg",

		"align-end.svg",

		"align-middle.svg",

		"align-start.svg",

		"align-top.svg",

		"alipay.svg",

		"alphabet-uppercase.svg",

		"alphabet.svg",

		"alt.svg",

		"amazon.svg",

		"amd.svg",

		"android.svg",

		"android2.svg",

		"anthropic.svg",

		"app-indicator.svg",

		"app.svg",

		"apple-music.svg",

		"apple.svg",

		"archive-fill.svg",

		"archive.svg",

		"arrow-90deg-down.svg",

		"arrow-90deg-left.svg",

		"arrow-90deg-right.svg",

		"arrow-90deg-up.svg",

		"arrow-bar-down.svg",

		"arrow-bar-left.svg",

		"arrow-bar-right.svg",

		"arrow-bar-up.svg",

		"arrow-clockwise.svg",

		"arrow-counterclockwise.svg",

		"arrow-down-circle-fill.svg",

		"arrow-down-circle.svg",

		"arrow-down-left-circle-fill.svg",

		"arrow-down-left-circle.svg",

		"arrow-down-left-square-fill.svg",

		"arrow-down-left-square.svg",

		"arrow-down-left.svg",

		"arrow-down-right-circle-fill.svg",

		"arrow-down-right-circle.svg",

		"arrow-down-right-square-fill.svg",

		"arrow-down-right-square.svg",

		"arrow-down-right.svg",

		"arrow-down-short.svg",

		"arrow-down-square-fill.svg",

		"arrow-down-square.svg",

		"arrow-down-up.svg",

		"arrow-down.svg",

		"arrow-left-circle-fill.svg",

		"arrow-left-circle.svg",

		"arrow-left-right.svg",

		"arrow-left-short.svg",

		"arrow-left-square-fill.svg",

		"arrow-left-square.svg",

		"arrow-left.svg",

		"arrow-repeat.svg",

		"arrow-return-left.svg",

		"arrow-return-right.svg",

		"arrow-right-circle-fill.svg",

		"arrow-right-circle.svg",

		"arrow-right-short.svg",

		"arrow-right-square-fill.svg",

		"arrow-right-square.svg",

		"arrow-right.svg",

		"arrow-through-heart-fill.svg",

		"arrow-through-heart.svg",

		"arrow-up-circle-fill.svg",

		"arrow-up-circle.svg",

		"arrow-up-left-circle-fill.svg",

		"arrow-up-left-circle.svg",

		"arrow-up-left-square-fill.svg",

		"arrow-up-left-square.svg",

		"arrow-up-left.svg",

		"arrow-up-right-circle-fill.svg",

		"arrow-up-right-circle.svg",

		"arrow-up-right-square-fill.svg",

		"arrow-up-right-square.svg",

		"arrow-up-right.svg",

		"arrow-up-short.svg",

		"arrow-up-square-fill.svg",

		"arrow-up-square.svg",

		"arrow-up.svg",

		"arrows-angle-contract.svg",

		"arrows-angle-expand.svg",

		"arrows-collapse-vertical.svg",

		"arrows-collapse.svg",

		"arrows-expand-vertical.svg",

		"arrows-expand.svg",

		"arrows-fullscreen.svg",

		"arrows-move.svg",

		"arrows-vertical.svg",

		"arrows.svg",

		"aspect-ratio-fill.svg",

		"aspect-ratio.svg",

		"asterisk.svg",

		"at.svg",

		"award-fill.svg",

		"award.svg",

		"back.svg",

		"backpack-fill.svg",

		"backpack.svg",

		"backpack2-fill.svg",

		"backpack2.svg",

		"backpack3-fill.svg",

		"backpack3.svg",

		"backpack4-fill.svg",

		"backpack4.svg",

		"backspace-fill.svg",

		"backspace-reverse-fill.svg",

		"backspace-reverse.svg",

		"backspace.svg",

		"badge-3d-fill.svg",

		"badge-3d.svg",

		"badge-4k-fill.svg",

		"badge-4k.svg",

		"badge-8k-fill.svg",

		"badge-8k.svg",

		"badge-ad-fill.svg",

		"badge-ad.svg",

		"badge-ar-fill.svg",

		"badge-ar.svg",

		"badge-cc-fill.svg",

		"badge-cc.svg",

		"badge-hd-fill.svg",

		"badge-hd.svg",

		"badge-sd-fill.svg",

		"badge-sd.svg",

		"badge-tm-fill.svg",

		"badge-tm.svg",

		"badge-vo-fill.svg",

		"badge-vo.svg",

		"badge-vr-fill.svg",

		"badge-vr.svg",

		"badge-wc-fill.svg",

		"badge-wc.svg",

		"bag-check-fill.svg",

		"bag-check.svg",

		"bag-dash-fill.svg",

		"bag-dash.svg",

		"bag-fill.svg",

		"bag-heart-fill.svg",

		"bag-heart.svg",

		"bag-plus-fill.svg",

		"bag-plus.svg",

		"bag-x-fill.svg",

		"bag-x.svg",

		"bag.svg",

		"balloon-fill.svg",

		"balloon-heart-fill.svg",

		"balloon-heart.svg",

		"balloon.svg",

		"ban-fill.svg",

		"ban.svg",

		"bandaid-fill.svg",

		"bandaid.svg",

		"bank.svg",

		"bank2.svg",

		"bar-chart-fill.svg",

		"bar-chart-line-fill.svg",

		"bar-chart-line.svg",

		"bar-chart-steps.svg",

		"bar-chart.svg",

		"basket-fill.svg",

		"basket.svg",

		"basket2-fill.svg",

		"basket2.svg",

		"basket3-fill.svg",
		"basket3.svg",

		"battery-charging.svg",

		"battery-full.svg",

		"battery-half.svg",

		"battery-low.svg",

		"battery.svg",

		"beaker-fill.svg",

		"beaker.svg",

		"behance.svg",

		"bell-fill.svg",

		"bell-slash-fill.svg",

		"bell-slash.svg",

		"bell.svg",

		"bezier.svg",

		"bezier2.svg",

		"bicycle.svg",

		"bing.svg",

		"binoculars-fill.svg",

		"binoculars.svg",

		"blockquote-left.svg",

		"blockquote-right.svg",

		"bluesky.svg",
		"bluetooth.svg",

		"body-text.svg",

		"book-fill.svg",

		"book-half.svg",

		"book.svg",

		"bookmark-check-fill.svg",

		"bookmark-check.svg",

		"bookmark-dash-fill.svg",

		"bookmark-dash.svg",

		"bookmark-fill.svg",

		"bookmark-heart-fill.svg",

		"bookmark-heart.svg",

		"bookmark-plus-fill.svg",

		"bookmark-plus.svg",

		"bookmark-star-fill.svg",

		"bookmark-star.svg",

		"bookmark-x-fill.svg",

		"bookmark-x.svg",

		"bookmark.svg",

		"bookmarks-fill.svg",

		"bookmarks.svg",

		"bookshelf.svg",

		"boombox-fill.svg",

		"boombox.svg",

		"bootstrap-fill.svg",

		"bootstrap-reboot.svg",

		"bootstrap.svg",

		"border-all.svg",

		"border-bottom.svg",

		"border-center.svg",

		"border-inner.svg",

		"border-left.svg",

		"border-middle.svg",

		"border-outer.svg",

		"border-right.svg",

		"border-style.svg",

		"border-top.svg",

		"border-width.svg",

		"border.svg",

		"bounding-box-circles.svg",

		"bounding-box.svg",

		"box-arrow-down-left.svg",

		"box-arrow-down-right.svg",

		"box-arrow-down.svg",

		"box-arrow-in-down-left.svg",

		"box-arrow-in-down-right.svg",

		"box-arrow-in-down.svg",

		"box-arrow-in-left.svg",

		"box-arrow-in-right.svg",

		"box-arrow-in-up-left.svg",

		"box-arrow-in-up-right.svg",

		"box-arrow-in-up.svg",

		"box-arrow-left.svg",

		"box-arrow-right.svg",

		"box-arrow-up-left.svg",

		"box-arrow-up-right.svg",

		"box-arrow-up.svg",

		"box-fill.svg",

		"box-seam-fill.svg",

		"box-seam.svg",

		"box.svg",

		"box2-fill.svg",

		"box2-heart-fill.svg",

		"box2-heart.svg",

		"box2.svg",

		"boxes.svg",

		"braces-asterisk.svg",

		"braces.svg",

		"bricks.svg",

		"briefcase-fill.svg",

		"briefcase.svg",

		"brightness-alt-high-fill.svg",

		"brightness-alt-high.svg",

		"brightness-alt-low-fill.svg",

		"brightness-alt-low.svg",

		"brightness-high-fill.svg",

		"brightness-high.svg",

		"brightness-low-fill.svg",

		"brightness-low.svg",

		"brilliance.svg",

		"broadcast-pin.svg",

		"broadcast.svg",

		"browser-chrome.svg",

		"browser-edge.svg",

		"browser-firefox.svg",

		"browser-safari.svg",

		"brush-fill.svg",

		"brush.svg",

		"bucket-fill.svg",

		"bucket.svg",

		"bug-fill.svg",

		"bug.svg",

		"building-add.svg",

		"building-check.svg",

		"building-dash.svg",

		"building-down.svg",

		"building-exclamation.svg",

		"building-fill-add.svg",

		"building-fill-check.svg",

		"building-fill-dash.svg",

		"building-fill-down.svg",

		"building-fill-exclamation.svg",

		"building-fill-gear.svg",

		"building-fill-lock.svg",

		"building-fill-slash.svg",

		"building-fill-up.svg",

		"building-fill-x.svg",

		"building-fill.svg",

		"building-gear.svg",

		"building-lock.svg",

		"building-slash.svg",

		"building-up.svg",

		"building-x.svg",

		"building.svg",

		"buildings-fill.svg",

		"buildings.svg",

		"bullseye.svg",

		"bus-front-fill.svg",

		"bus-front.svg",

		"c-circle-fill.svg",

		"c-circle.svg",

		"c-square-fill.svg",

		"c-square.svg",

		"cake-fill.svg",

		"cake.svg",

		"cake2-fill.svg",

		"cake2.svg",

		"calculator-fill.svg",

		"calculator.svg",

		"calendar-check-fill.svg",

		"calendar-check.svg",

		"calendar-date-fill.svg",

		"calendar-date.svg",

		"calendar-day-fill.svg",

		"calendar-day.svg",

		"calendar-event-fill.svg",

		"calendar-event.svg",

		"calendar-fill.svg",

		"calendar-heart-fill.svg",

		"calendar-heart.svg",

		"calendar-minus-fill.svg",

		"calendar-minus.svg",

		"calendar-month-fill.svg",

		"calendar-month.svg",

		"calendar-plus-fill.svg",

		"calendar-plus.svg",

		"calendar-range-fill.svg",

		"calendar-range.svg",

		"calendar-week-fill.svg",

		"calendar-week.svg",

		"calendar-x-fill.svg",

		"calendar-x.svg",

		"calendar.svg",

		"calendar2-check-fill.svg",

		"calendar2-check.svg",

		"calendar2-date-fill.svg",

		"calendar2-date.svg",

		"calendar2-day-fill.svg",

		"calendar2-day.svg",

		"calendar2-event-fill.svg",

		"calendar2-event.svg",

		"calendar2-fill.svg",

		"calendar2-heart-fill.svg",

		"calendar2-heart.svg",

		"calendar2-minus-fill.svg",

		"calendar2-minus.svg",

		"calendar2-month-fill.svg",

		"calendar2-month.svg",

		"calendar2-plus-fill.svg",

		"calendar2-plus.svg",

		"calendar2-range-fill.svg",

		"calendar2-range.svg",

		"calendar2-week-fill.svg",

		"calendar2-week.svg",

		"calendar2-x-fill.svg",

		"calendar2-x.svg",

		"calendar2.svg",

		"calendar3-event-fill.svg",
		"calendar3-event.svg",

		"calendar3-fill.svg",

		"calendar3-range-fill.svg",
		"calendar3-range.svg",

		"calendar3-week-fill.svg",
		"calendar3-week.svg",

		"calendar3.svg",

		"calendar4-event.svg",

		"calendar4-range.svg",

		"calendar4-week.svg",

		"calendar4.svg",

		"camera-fill.svg",

		"camera-reels-fill.svg",

		"camera-reels.svg",

		"camera-video-fill.svg",

		"camera-video-off-fill.svg",

		"camera-video-off.svg",

		"camera-video.svg",

		"camera.svg",

		"camera2.svg",

		"capslock-fill.svg",

		"capslock.svg",

		"capsule-pill.svg",

		"capsule.svg",

		"car-front-fill.svg",

		"car-front.svg",

		"card-checklist.svg",

		"card-heading.svg",

		"card-image.svg",

		"card-list.svg",

		"card-text.svg",

		"caret-down-fill.svg",
		"caret-down-square-fill.svg",

		"caret-down-square.svg",

		"caret-down.svg",

		"caret-left-fill.svg",
		"caret-left-square-fill.svg",

		"caret-left-square.svg",

		"caret-left.svg",

		"caret-right-fill.svg",
		"caret-right-square-fill.svg",

		"caret-right-square.svg",

		"caret-right.svg",

		"caret-up-fill.svg",
		"caret-up-square-fill.svg",

		"caret-up-square.svg",

		"caret-up.svg",

		"cart-check-fill.svg",

		"cart-check.svg",

		"cart-dash-fill.svg",

		"cart-dash.svg",

		"cart-fill.svg",

		"cart-plus-fill.svg",

		"cart-plus.svg",

		"cart-x-fill.svg",

		"cart-x.svg",

		"cart.svg",

		"cart2.svg",

		"cart3.svg",

		"cart4.svg",

		"cash-coin.svg",

		"cash-stack.svg",

		"cash.svg",

		"cassette-fill.svg",

		"cassette.svg",

		"cast.svg",

		"cc-circle-fill.svg",

		"cc-circle.svg",

		"cc-square-fill.svg",

		"cc-square.svg",

		"chat-dots-fill.svg",

		"chat-dots.svg",

		"chat-fill.svg",

		"chat-heart-fill.svg",

		"chat-heart.svg",

		"chat-left-dots-fill.svg",

		"chat-left-dots.svg",

		"chat-left-fill.svg",

		"chat-left-heart-fill.svg",

		"chat-left-heart.svg",

		"chat-left-quote-fill.svg",

		"chat-left-quote.svg",

		"chat-left-text-fill.svg",

		"chat-left-text.svg",

		"chat-left.svg",

		"chat-quote-fill.svg",

		"chat-quote.svg",

		"chat-right-dots-fill.svg",

		"chat-right-dots.svg",

		"chat-right-fill.svg",

		"chat-right-heart-fill.svg",

		"chat-right-heart.svg",

		"chat-right-quote-fill.svg",

		"chat-right-quote.svg",

		"chat-right-text-fill.svg",

		"chat-right-text.svg",

		"chat-right.svg",

		"chat-square-dots-fill.svg",

		"chat-square-dots.svg",

		"chat-square-fill.svg",

		"chat-square-heart-fill.svg",

		"chat-square-heart.svg",

		"chat-square-quote-fill.svg",

		"chat-square-quote.svg",

		"chat-square-text-fill.svg",

		"chat-square-text.svg",

		"chat-square.svg",

		"chat-text-fill.svg",

		"chat-text.svg",

		"chat.svg",

		"check-all.svg",

		"check-circle-fill.svg",

		"check-circle.svg",

		"check-lg.svg",

		"check-square-fill.svg",

		"check-square.svg",

		"check.svg",

		"check2-all.svg",

		"check2-circle.svg",

		"check2-square.svg",

		"check2.svg",

		"chevron-bar-contract.svg",

		"chevron-bar-down.svg",

		"chevron-bar-expand.svg",

		"chevron-bar-left.svg",

		"chevron-bar-right.svg",

		"chevron-bar-up.svg",

		"chevron-compact-down.svg",

		"chevron-compact-left.svg",

		"chevron-compact-right.svg",

		"chevron-compact-up.svg",

		"chevron-contract.svg",

		"chevron-double-down.svg",

		"chevron-double-left.svg",

		"chevron-double-right.svg",

		"chevron-double-up.svg",

		"chevron-down.svg",

		"chevron-expand.svg",

		"chevron-left.svg",

		"chevron-right.svg",

		"chevron-up.svg",

		"circle-fill.svg",
		"circle-half.svg",

		"circle-square.svg",

		"circle.svg",

		"claude.svg",
		"clipboard-check-fill.svg",

		"clipboard-check.svg",

		"clipboard-data-fill.svg",

		"clipboard-data.svg",

		"clipboard-fill.svg",

		"clipboard-heart-fill.svg",

		"clipboard-heart.svg",

		"clipboard-minus-fill.svg",

		"clipboard-minus.svg",

		"clipboard-plus-fill.svg",

		"clipboard-plus.svg",

		"clipboard-pulse.svg",

		"clipboard-x-fill.svg",

		"clipboard-x.svg",

		"clipboard.svg",

		"clipboard2-check-fill.svg",

		"clipboard2-check.svg",

		"clipboard2-data-fill.svg",

		"clipboard2-data.svg",

		"clipboard2-fill.svg",

		"clipboard2-heart-fill.svg",

		"clipboard2-heart.svg",

		"clipboard2-minus-fill.svg",

		"clipboard2-minus.svg",

		"clipboard2-plus-fill.svg",

		"clipboard2-plus.svg",

		"clipboard2-pulse-fill.svg",

		"clipboard2-pulse.svg",

		"clipboard2-x-fill.svg",

		"clipboard2-x.svg",

		"clipboard2.svg",

		"clock-fill.svg",

		"clock-history.svg",

		"clock.svg",

		"cloud-arrow-down-fill.svg",

		"cloud-arrow-down.svg",

		"cloud-arrow-up-fill.svg",

		"cloud-arrow-up.svg",

		"cloud-check-fill.svg",

		"cloud-check.svg",

		"cloud-download-fill.svg",

		"cloud-download.svg",

		"cloud-drizzle-fill.svg",

		"cloud-drizzle.svg",

		"cloud-fill.svg",

		"cloud-fog-fill.svg",

		"cloud-fog.svg",

		"cloud-fog2-fill.svg",

		"cloud-fog2.svg",

		"cloud-hail-fill.svg",

		"cloud-hail.svg",

		"cloud-haze-fill.svg",

		"cloud-haze.svg",

		"cloud-haze2-fill.svg",

		"cloud-haze2.svg",

		"cloud-lightning-fill.svg",

		"cloud-lightning-rain-fill.svg",

		"cloud-lightning-rain.svg",

		"cloud-lightning.svg",

		"cloud-minus-fill.svg",

		"cloud-minus.svg",

		"cloud-moon-fill.svg",

		"cloud-moon.svg",

		"cloud-plus-fill.svg",

		"cloud-plus.svg",

		"cloud-rain-fill.svg",

		"cloud-rain-heavy-fill.svg",

		"cloud-rain-heavy.svg",

		"cloud-rain.svg",

		"cloud-slash-fill.svg",

		"cloud-slash.svg",

		"cloud-sleet-fill.svg",

		"cloud-sleet.svg",

		"cloud-snow-fill.svg",

		"cloud-snow.svg",

		"cloud-sun-fill.svg",

		"cloud-sun.svg",

		"cloud-upload-fill.svg",

		"cloud-upload.svg",

		"cloud.svg",

		"clouds-fill.svg",

		"clouds.svg",

		"cloudy-fill.svg",

		"cloudy.svg",

		"code-slash.svg",

		"code-square.svg",

		"code.svg",

		"coin.svg",

		"collection-fill.svg",

		"collection-play-fill.svg",

		"collection-play.svg",

		"collection.svg",

		"columns-gap.svg",

		"columns.svg",

		"command.svg",

		"compass-fill.svg",

		"compass.svg",

		"cone-striped.svg",

		"cone.svg",

		"controller.svg",

		"cookie.svg",

		"copy.svg",

		"cpu-fill.svg",

		"cpu.svg",

		"credit-card-2-back-fill.svg",

		"credit-card-2-back.svg",

		"credit-card-2-front-fill.svg",

		"credit-card-2-front.svg",

		"credit-card-fill.svg",

		"credit-card.svg",

		"crop.svg",

		"crosshair.svg",

		"crosshair2.svg",

		"css.svg",
		"cup-fill.svg",

		"cup-hot-fill.svg",

		"cup-hot.svg",

		"cup-straw.svg",

		"cup.svg",

		"currency-bitcoin.svg",

		"currency-dollar.svg",

		"currency-euro.svg",

		"currency-exchange.svg",

		"currency-pound.svg",

		"currency-rupee.svg",

		"currency-yen.svg",
		"cursor-fill.svg",
		"cursor-text.svg",

		"cursor.svg",

		"dash-circle-dotted.svg",

		"dash-circle-fill.svg",

		"dash-circle.svg",

		"dash-lg.svg",

		"dash-square-dotted.svg",

		"dash-square-fill.svg",

		"dash-square.svg",

		"dash.svg",

		"database-add.svg",

		"database-check.svg",

		"database-dash.svg",

		"database-down.svg",

		"database-exclamation.svg",

		"database-fill-add.svg",

		"database-fill-check.svg",

		"database-fill-dash.svg",

		"database-fill-down.svg",

		"database-fill-exclamation.svg",

		"database-fill-gear.svg",

		"database-fill-lock.svg",

		"database-fill-slash.svg",

		"database-fill-up.svg",

		"database-fill-x.svg",

		"database-fill.svg",

		"database-gear.svg",

		"database-lock.svg",

		"database-slash.svg",

		"database-up.svg",

		"database-x.svg",

		"database.svg",

		"device-hdd-fill.svg",

		"device-hdd.svg",

		"device-ssd-fill.svg",

		"device-ssd.svg",

		"diagram-2-fill.svg",

		"diagram-2.svg",

		"diagram-3-fill.svg",

		"diagram-3.svg",

		"diamond-fill.svg",

		"diamond-half.svg",

		"diamond.svg",

		"dice-1-fill.svg",

		"dice-1.svg",

		"dice-2-fill.svg",

		"dice-2.svg",

		"dice-3-fill.svg",

		"dice-3.svg",

		"dice-4-fill.svg",

		"dice-4.svg",

		"dice-5-fill.svg",

		"dice-5.svg",

		"dice-6-fill.svg",

		"dice-6.svg",

		"disc-fill.svg",

		"disc.svg",

		"discord.svg",

		"display-fill.svg",

		"display.svg",

		"displayport-fill.svg",

		"displayport.svg",

		"distribute-horizontal.svg",

		"distribute-vertical.svg",

		"door-closed-fill.svg",

		"door-closed.svg",

		"door-open-fill.svg",

		"door-open.svg",

		"dot.svg",

		"download.svg",

		"dpad-fill.svg",

		"dpad.svg",

		"dribbble.svg",

		"dropbox.svg",

		"droplet-fill.svg",

		"droplet-half.svg",

		"droplet.svg",

		"duffle-fill.svg",

		"duffle.svg",

		"ear-fill.svg",

		"ear.svg",

		"earbuds.svg",

		"easel-fill.svg",

		"easel.svg",

		"easel2-fill.svg",

		"easel2.svg",

		"easel3-fill.svg",

		"easel3.svg",

		"egg-fill.svg",

		"egg-fried.svg",

		"egg.svg",

		"eject-fill.svg",

		"eject.svg",

		"emoji-angry-fill.svg",

		"emoji-angry.svg",

		"emoji-astonished-fill.svg",

		"emoji-astonished.svg",

		"emoji-dizzy-fill.svg",

		"emoji-dizzy.svg",

		"emoji-expressionless-fill.svg",

		"emoji-expressionless.svg",

		"emoji-frown-fill.svg",

		"emoji-frown.svg",

		"emoji-grimace-fill.svg",

		"emoji-grimace.svg",

		"emoji-grin-fill.svg",

		"emoji-grin.svg",

		"emoji-heart-eyes-fill.svg",

		"emoji-heart-eyes.svg",

		"emoji-kiss-fill.svg",

		"emoji-kiss.svg",

		"emoji-laughing-fill.svg",

		"emoji-laughing.svg",

		"emoji-neutral-fill.svg",

		"emoji-neutral.svg",

		"emoji-smile-fill.svg",

		"emoji-smile-upside-down-fill.svg",

		"emoji-smile-upside-down.svg",

		"emoji-smile.svg",

		"emoji-sunglasses-fill.svg",

		"emoji-sunglasses.svg",

		"emoji-surprise-fill.svg",

		"emoji-surprise.svg",

		"emoji-tear-fill.svg",

		"emoji-tear.svg",

		"emoji-wink-fill.svg",

		"emoji-wink.svg",

		"envelope-arrow-down-fill.svg",

		"envelope-arrow-down.svg",

		"envelope-arrow-up-fill.svg",

		"envelope-arrow-up.svg",

		"envelope-at-fill.svg",

		"envelope-at.svg",

		"envelope-check-fill.svg",

		"envelope-check.svg",

		"envelope-dash-fill.svg",

		"envelope-dash.svg",

		"envelope-exclamation-fill.svg",

		"envelope-exclamation.svg",

		"envelope-fill.svg",

		"envelope-heart-fill.svg",

		"envelope-heart.svg",

		"envelope-open-fill.svg",

		"envelope-open-heart-fill.svg",

		"envelope-open-heart.svg",

		"envelope-open.svg",

		"envelope-paper-fill.svg",

		"envelope-paper-heart-fill.svg",

		"envelope-paper-heart.svg",

		"envelope-paper.svg",

		"envelope-plus-fill.svg",

		"envelope-plus.svg",

		"envelope-slash-fill.svg",

		"envelope-slash.svg",

		"envelope-x-fill.svg",

		"envelope-x.svg",

		"envelope.svg",

		"eraser-fill.svg",

		"eraser.svg",

		"escape.svg",

		"ethernet.svg",

		"ev-front-fill.svg",

		"ev-front.svg",

		"ev-station-fill.svg",

		"ev-station.svg",

		"exclamation-circle-fill.svg",

		"exclamation-circle.svg",

		"exclamation-diamond-fill.svg",

		"exclamation-diamond.svg",

		"exclamation-lg.svg",

		"exclamation-octagon-fill.svg",

		"exclamation-octagon.svg",

		"exclamation-square-fill.svg",

		"exclamation-square.svg",

		"exclamation-triangle-fill.svg",

		"exclamation-triangle.svg",

		"exclamation.svg",

		"exclude.svg",

		"explicit-fill.svg",

		"explicit.svg",

		"exposure.svg",

		"eye-fill.svg",

		"eye-slash-fill.svg",

		"eye-slash.svg",

		"eye.svg",

		"eyedropper.svg",

		"eyeglasses.svg",

		"facebook.svg",

		"fan.svg",

		"fast-forward-btn-fill.svg",

		"fast-forward-btn.svg",

		"fast-forward-circle-fill.svg",

		"fast-forward-circle.svg",

		"fast-forward-fill.svg",

		"fast-forward.svg",

		"feather.svg",

		"feather2.svg",

		"file-arrow-down-fill.svg",

		"file-arrow-down.svg",

		"file-arrow-up-fill.svg",

		"file-arrow-up.svg",

		"file-bar-graph-fill.svg",

		"file-bar-graph.svg",

		"file-binary-fill.svg",

		"file-binary.svg",

		"file-break-fill.svg",

		"file-break.svg",

		"file-check-fill.svg",

		"file-check.svg",

		"file-code-fill.svg",

		"file-code.svg",

		"file-diff-fill.svg",

		"file-diff.svg",

		"file-earmark-arrow-down-fill.svg",

		"file-earmark-arrow-down.svg",

		"file-earmark-arrow-up-fill.svg",

		"file-earmark-arrow-up.svg",

		"file-earmark-bar-graph-fill.svg",

		"file-earmark-bar-graph.svg",

		"file-earmark-binary-fill.svg",

		"file-earmark-binary.svg",

		"file-earmark-break-fill.svg",

		"file-earmark-break.svg",

		"file-earmark-check-fill.svg",

		"file-earmark-check.svg",

		"file-earmark-code-fill.svg",

		"file-earmark-code.svg",

		"file-earmark-diff-fill.svg",

		"file-earmark-diff.svg",

		"file-earmark-easel-fill.svg",

		"file-earmark-easel.svg",

		"file-earmark-excel-fill.svg",

		"file-earmark-excel.svg",

		"file-earmark-fill.svg",

		"file-earmark-font-fill.svg",

		"file-earmark-font.svg",

		"file-earmark-image-fill.svg",

		"file-earmark-image.svg",

		"file-earmark-lock-fill.svg",

		"file-earmark-lock.svg",

		"file-earmark-lock2-fill.svg",

		"file-earmark-lock2.svg",

		"file-earmark-medical-fill.svg",

		"file-earmark-medical.svg",

		"file-earmark-minus-fill.svg",

		"file-earmark-minus.svg",

		"file-earmark-music-fill.svg",

		"file-earmark-music.svg",

		"file-earmark-pdf-fill.svg",

		"file-earmark-pdf.svg",

		"file-earmark-person-fill.svg",

		"file-earmark-person.svg",

		"file-earmark-play-fill.svg",

		"file-earmark-play.svg",

		"file-earmark-plus-fill.svg",

		"file-earmark-plus.svg",

		"file-earmark-post-fill.svg",

		"file-earmark-post.svg",

		"file-earmark-ppt-fill.svg",

		"file-earmark-ppt.svg",

		"file-earmark-richtext-fill.svg",

		"file-earmark-richtext.svg",

		"file-earmark-ruled-fill.svg",

		"file-earmark-ruled.svg",

		"file-earmark-slides-fill.svg",

		"file-earmark-slides.svg",

		"file-earmark-spreadsheet-fill.svg",

		"file-earmark-spreadsheet.svg",

		"file-earmark-text-fill.svg",

		"file-earmark-text.svg",

		"file-earmark-word-fill.svg",

		"file-earmark-word.svg",

		"file-earmark-x-fill.svg",

		"file-earmark-x.svg",

		"file-earmark-zip-fill.svg",

		"file-earmark-zip.svg",

		"file-earmark.svg",

		"file-easel-fill.svg",

		"file-easel.svg",

		"file-excel-fill.svg",

		"file-excel.svg",

		"file-fill.svg",

		"file-font-fill.svg",

		"file-font.svg",

		"file-image-fill.svg",

		"file-image.svg",

		"file-lock-fill.svg",

		"file-lock.svg",

		"file-lock2-fill.svg",

		"file-lock2.svg",

		"file-medical-fill.svg",

		"file-medical.svg",

		"file-minus-fill.svg",

		"file-minus.svg",

		"file-music-fill.svg",

		"file-music.svg",

		"file-pdf-fill.svg",

		"file-pdf.svg",

		"file-person-fill.svg",

		"file-person.svg",

		"file-play-fill.svg",

		"file-play.svg",

		"file-plus-fill.svg",

		"file-plus.svg",

		"file-post-fill.svg",

		"file-post.svg",

		"file-ppt-fill.svg",

		"file-ppt.svg",

		"file-richtext-fill.svg",

		"file-richtext.svg",

		"file-ruled-fill.svg",
	}
)
