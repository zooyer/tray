package tray

import (
	"github.com/lxn/walk"
)

type Tray struct {
	win     *walk.MainWindow
	not     *walk.NotifyIcon
	icon    *walk.Icon
	options *Options
}

type Options struct {
	Tip   string
	Icon  string
	Click func(x, y int)
	Menus Menus
}

type Option func(opts *Options)

type Menu struct {
	Name   string
	Action func(tray *Tray)
	//Action func()
}

type Menus []Menu

func (t *Tray) init() (err error) {
	defer func() {
		if err != nil {
			_ = t.dispose()
		}
	}()

	win, err := walk.NewMainWindow()
	if err != nil {
		return
	}
	t.win = win

	icon, err := walk.Resources.Icon(t.options.Icon)
	if err != nil {
		return
	}
	t.icon = icon

	not, err := walk.NewNotifyIcon(win)
	if err != nil {
		return
	}
	t.not = not

	if err = not.SetIcon(icon); err != nil {
		return
	}

	if t.options.Tip != "" {
		if err = not.SetToolTip(t.options.Tip); err != nil {
			return
		}
	}

	if t.options.Click != nil {
		not.MouseUp().Attach(func(x, y int, button walk.MouseButton) {
			if button == walk.LeftButton {
				t.options.Click(x, y)
			}
		})
	}

	for _, menu := range t.options.Menus {
		act := walk.NewAction()
		if err = act.SetText(menu.Name); err != nil {
			return
		}
		if menu.Action != nil {
			menu := menu
			act.Triggered().Attach(func() {
				menu.Action(t)
			})
		}
		if err = not.ContextMenu().Actions().Add(act); err != nil {
			return
		}
	}

	if err = not.SetVisible(true); err != nil {
		return
	}

	return
}

func (t *Tray) dispose() (err error) {
	if t.win != nil {
		t.win.Dispose()
	}
	if t.icon != nil {
		t.icon.Dispose()
	}
	if t.not != nil {
		if err = t.not.Dispose(); err != nil {
			return
		}
	}

	return
}

func (t *Tray) Run() {
	t.win.Run()
}

func (t *Tray) Stop() (err error) {
	if err = t.not.SetVisible(false); err != nil {
		return err
	}

	if err = t.dispose(); err != nil {
		return
	}

	return t.win.Close()
}

func (t *Tray) Message(title, message string) {
	_ = t.not.ShowMessage(title, message)
}

func WithTip(tip string) Option {
	return func(opts *Options) {
		opts.Tip = tip
	}
}

func WithIcon(icon string) Option {
	return func(opts *Options) {
		opts.Icon = icon
	}
}

func WithClick(click func(x, y int)) Option {
	return func(opts *Options) {
		opts.Click = click
	}
}

func WithMenus(menus Menus) Option {
	return func(opts *Options) {
		opts.Menus = menus
	}
}

func New(options *Options, option ...Option) (*Tray, error) {
	if options == nil {
		options = new(Options)
	}
	for _, opt := range option {
		opt(options)
	}

	tray := &Tray{
		win:     nil,
		not:     nil,
		options: options,
	}

	return tray, tray.init()
}
