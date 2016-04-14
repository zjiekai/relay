package main

import (
	"crypto/cipher"
	"crypto/rc4"
	"flag"
	"io"
	"net"
	"sync"
)

var (
	laddr  = flag.String("l", "0.0.0.0:8081", "local address")
	raddr  = flag.String("r", "127.0.0.1:8080", "remote address")
	inwall = flag.Bool("w", true, "inside GFW")
)

func Start() {
	ln, err := net.Listen("tcp", *laddr)
	if err != nil {
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func Pipe(src io.ReadWriteCloser, dst io.ReadWriteCloser) (int64, int64) {
	var sent, received int64
	var c = make(chan struct{})
	var o sync.Once

	close := func() {
		src.Close()
		dst.Close()
		close(c)
	}

	go func() {
		received, _ = io.Copy(src, dst)
		o.Do(close)
	}()

	go func() {
		sent, _ = io.Copy(dst, src)
		o.Do(close)
	}()

	<-c
	return sent, received
}

type RC4Cipher struct {
	*cipher.StreamReader
	*cipher.StreamWriter
}

func NewRC4Cipher(rwc io.ReadWriteCloser) (*RC4Cipher, error) {
	decryptCipher, _ := rc4.NewCipher([]byte("gfw"))
	encryptCipher, _ := rc4.NewCipher([]byte("gfw"))
	return &RC4Cipher{
		StreamReader: &cipher.StreamReader{
			S: decryptCipher,
			R: rwc,
		},
		StreamWriter: &cipher.StreamWriter{
			S: encryptCipher,
			W: rwc,
		},
	}, nil
}

func handleConn(c net.Conn) {
	var local, remote io.ReadWriteCloser
	rc, err := net.Dial("tcp", *raddr)
	if err != nil {
		return
	}
	if *inwall {
		local = c
		remote, _ = NewRC4Cipher(rc)
	} else {
		local, _ = NewRC4Cipher(c)
		remote = rc
	}
	Pipe(local, remote)
}

func main() {
	flag.Parse()
	Start()
}
