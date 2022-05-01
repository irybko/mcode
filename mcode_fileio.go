package mcode
//
import (
    "os"
    "fmt"
    "net/http"
    "log"
    "io"
    "io/ioutil"
    "crypto/md5"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "crypto/rand"
)
//
func GetFileContentType(out *os.File) (string, error) {
    // Only the first 512 bytes are used to sniff the content type.
    buffer := make([]byte, 512)

    _, err := out.Read(buffer)
    if err != nil {
        return "", err
    }
    // Use the net/http package's handy DectectContentType function. Always returns a valid 
    // content-type by returning "application/octet-stream" if no others seemed to match.
    contentType := http.DetectContentType(buffer)

    return contentType, nil
}
//
func GetFileMode(path string) os.FileMode {
    fi, err := os.Lstat(path)
    if err != nil {
        log.Fatal(err)
    }
    return fi.Mode().Perm()
}
//
func CurSize(path string) int64 {
    fi, err := os.Lstat(path)
    if err != nil {
        log.Fatal(err)
    }
    return fi.Size()
}
// Random Access Binary File
type RandomAccess struct {
	Path    string
	Data    []byte
    Offset  int64
	Datasz  int
    //
    Crypto  bool
    Passwd  string
}
//
func (inst RandomAccess) CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
//
func (inst RandomAccess) Read() {
	plen := len(inst.Path)
    var fmode    os.FileMode

    switch {
		case plen == 0: log.Fatal("Path to data file has no found")
		case plen > 0:
			fmode = GetFileMode(inst.Path)
            fd, err := os.OpenFile(inst.Path, os.O_RDONLY, fmode)
            defer fd.Close()

            if err != nil {
				log.Fatal(err)
			}

			switch {
				case inst.Offset == 0:
					content := make([]byte, CurSize(inst.Path))
                    content, err =  ioutil.ReadFile(inst.Path)

					if err == nil || err == io.EOF {
						log.Fatal(err)
					}
                    fmt.Printf("%s\n", inst.Data)
					inst.Data = content
				case inst.Offset > 0:
					// read
					bufdata := make([]byte, inst.Datasz)
					_, err := fd.ReadAt(bufdata, inst.Offset)
			                if err != nil {
						log.Fatal(err)
					}
					inst.Data = bufdata
			}
    }
    //
    switch inst.Crypto {
    case true:
        inst.Data = inst.DecryptContent(inst.Data, inst.Passwd)
    case false:
        return
    }
}
//
func (inst RandomAccess) DecryptContent(data []byte, passphrase string) []byte {
	key := []byte(inst.CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
//
func (inst RandomAccess) Write() bool {
	plen := len(inst.Path)
	var result bool
    var fmode  os.FileMode
    //
    if inst.Crypto {
        inst.Data = inst.EncryptContent(inst.Data, inst.Passwd)
    }
    //
	switch {
		case plen == 0: log.Fatal("Path to data file has no founf")
		case plen > 0:
			fmode = GetFileMode(inst.Path)
            fd, err := os.OpenFile(inst.Path, os.O_RDONLY, fmode)
            defer fd.Close()

            if err != nil {
				log.Fatal(err)
			}

			switch {
				case inst.Offset == 0:
					// write data to file
					mutex.Lock()
					_, err = fd.Write(inst.Data)
					if err != nil {
						result = false
					}
					mutex.Unlock()

				case inst.Offset > 0:
					maxlen := len(inst.Data)
					// write
					mutex.Lock()

					nbytes, err := fd.WriteAt(inst.Data, inst.Offset)
					if err != nil {
						log.Fatal(err)
		                        }
					if nbytes == maxlen {
						result = true
                    } else {
                        result = false
                    }
                    mutex.Unlock()
				}
	}
    //
	return result
}
//
func (inst RandomAccess) EncryptContent(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(inst.CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
//


