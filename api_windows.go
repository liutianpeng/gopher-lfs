package lfs

import (
	"os"
	"syscall"
	"time"

	"github.com/yuin/gopher-lua"
)

func attributesFill(tbl *lua.LTable, stat os.FileInfo) {
	sys, ok := stat.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return
	}

	var mode string
	if stat.Mode().IsDir(){
		mode = "directory"
	} else if stat.Mode().IsRegular(){
		mode = "file"
	}else{
		mode = "file"
	}
	tbl.RawSetH(lua.LString("mode"), lua.LString(mode))

	tbl.RawSetH(lua.LString("dev"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("ino"), lua.LNumber(0))

	ss := stat.Sys()
	if ss!=nil{
		tbl.RawSetH(lua.LString("size"), lua.LNumber(stat.Size()))
		tbl.RawSetH(lua.LString("change"), lua.LNumber(sys.LastWriteTime.Nanoseconds()/int64(time.Second)))
		tbl.RawSetH(lua.LString("access"), lua.LNumber(sys.LastAccessTime.Nanoseconds()/int64(time.Second)))

		tbl.RawSetH(lua.LString("nlink"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("uid"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("gid"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("rdev"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("modification"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("blocks"), lua.LNumber(0))
		tbl.RawSetH(lua.LString("blksize"), lua.LNumber(0))
	}

/*
	sys, ok := stat.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return
	}
	tbl.RawSetH(lua.LString("dev"), lua.LNumber(sys.Dev))
	tbl.RawSetH(lua.LString("ino"), lua.LNumber(sys.Ino))
	{
		var mode string
		switch sys.Mode & syscall.S_IFMT {
		case syscall.S_IFREG:
			mode = "file"
		case syscall.S_IFDIR:
			mode = "directory"
		case syscall.S_IFLNK:
			mode = "link"
		case syscall.S_IFSOCK:
			mode = "socket"
		case syscall.S_IFIFO:
			mode = "named pipe"
		case syscall.S_IFCHR:
			mode = "char device"
		case syscall.S_IFBLK:
			mode = "block device"
		default:
			mode = "other"
		}
		tbl.RawSetH(lua.LString("mode"), lua.LString(mode))
	}
	tbl.RawSetH(lua.LString("nlink"), lua.LNumber(sys.Nlink))
	tbl.RawSetH(lua.LString("uid"), lua.LNumber(sys.Uid))
	tbl.RawSetH(lua.LString("gid"), lua.LNumber(sys.Gid))
	tbl.RawSetH(lua.LString("rdev"), lua.LNumber(sys.Rdev))
	tbl.RawSetH(lua.LString("access"), lua.LNumber(sys.Atim.Sec))
	tbl.RawSetH(lua.LString("modification"), lua.LNumber(sys.Mtim.Sec))
	tbl.RawSetH(lua.LString("change"), lua.LNumber(sys.Ctim.Sec))
	tbl.RawSetH(lua.LString("size"), lua.LNumber(sys.Size))
	tbl.RawSetH(lua.LString("blocks"), lua.LNumber(sys.Blocks))
	tbl.RawSetH(lua.LString("blksize"), lua.LNumber(sys.Blksize))
*/
}
