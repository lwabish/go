package exp

import (
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/lwabish/go/pkg/util"
)

const (
	hugePageFile = "/dev/hugepages/go-test"
)

type Mapper struct {
	file *os.File
	data *[]byte
	addr []byte
	// size in bytes
	size int
}

func NewMapper(file *os.File, size int) *Mapper {
	m := &Mapper{
		file: file,
		size: size * (1 << 20),
	}
	return m
}

func (m *Mapper) mmap() {
	// os.File.Fd 在osx不存在，需要在idea里把目标平台改为linux
	b, err := syscall.Mmap(int(m.file.Fd()), 0, m.size, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	util.Pe(err == nil, "failed to mmap", err)
	m.addr = b
	m.data = (*[]byte)(unsafe.Pointer(
		&struct {
			addr *byte
			len  int
			cap  int
		}{
			addr: &b[0],
			len:  m.size,
			cap:  m.size,
		},
	))
}

func (m *Mapper) munmap() {
	// fixme: m.data改成slice后，如果size是奇数，munmap会失败
	// 查看/dev/hugepages/go-test可以看到文件是偶数大小
	err := syscall.Munmap(m.addr)
	if err != nil {
		log.Println("munmap", err)
	}
	m.data = nil
	m.addr = nil
}

func (m *Mapper) writeData(d string) {
	target := m.data
	for i, v := range d {
		(*target)[i] = byte(v)
	}
}

// Run open a file in huge page mount point, mmap it into memory, write data into it.
func Run(s bool, size int) {
	var e error
	e = os.Remove(hugePageFile)
	f, e := os.OpenFile(hugePageFile, os.O_CREATE|os.O_RDWR, 0644)
	util.Pe(e == nil, "open file error: ", e)
	defer func(f *os.File) {
		util.Pe(f.Close() == nil, "close file error: ")
	}(f)

	mapper := NewMapper(f, size)
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
