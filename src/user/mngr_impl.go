package user

import (
	"github.com/perisie/mouse"
)

type Mngr_impl struct {
	mouse_fs *mouse.Mouse_fs
}

func Mngr_impl_new() *Mngr_impl {
	return &Mngr_impl{
		mouse_fs: mouse.Mouse_fs_new("data"),
	}
}
