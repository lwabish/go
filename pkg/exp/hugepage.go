package exp

import (
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/lwabish/go/pkg/util"
)

const (
	hugePageFile       = "/hugepages/test-go"
	defaultMaxFileSize = 1 << 30
	defaultMemMapSize  = 128 * (1 << 20)
)

type Mapper struct {
	file *os.File
	data *[defaultMaxFileSize]byte
	addr []byte
}

func (m *Mapper) mmap() {
	// os.File.Fd 在osx不存在，需要在idea里把目标平台改为linux
	b, err := syscall.Mmap(int(m.file.Fd()), 0, defaultMemMapSize, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	util.Pe(err == nil, "failed to mmap", err)
	m.addr = b
	m.data = (*[defaultMaxFileSize]byte)(unsafe.Pointer(&b[0]))
}

func (m *Mapper) munmap() {
	util.Pe(syscall.Munmap(m.addr) == nil, "failed to munmap")
	m.data = nil
	m.addr = nil
}

func (m *Mapper) writeData(d string) {
	for i, v := range d {
		m.data[i] = byte(v)
	}
}

// Run open a file in huge page mount point, mmap it into memory, write data into it.
func Run(s bool) {
	var e error
	e = os.Remove(hugePageFile)
	f, e := os.OpenFile(hugePageFile, os.O_CREATE|os.O_RDWR, 0644)
	util.Pe(e == nil, "open file error: ", e)
	defer func(f *os.File) {
		util.Pe(f.Close() == nil, "close file error: ")
	}(f)

	mapper := Mapper{file: f}
	mapper.mmap()
	defer mapper.munmap()
	log.Printf("Returned address is %p\n", mapper.addr)

	log.Printf("Writing data into hugepage memory")
	mapper.writeData("lwabish go huge page test")

	if s {
		log.Printf("Writing done, sleep forever...")
		select {}
	}
}
