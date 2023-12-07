package hashmap

import (
	"bytes"
	"encoding/gob"
	"hash"
)

type Hasher struct {
	hasher  hash.Hash32
	buf     *bytes.Buffer
	encoder *gob.Encoder
}

// returns a new Hasher that shares resources with old one, but has new hashFn
func (hasher *Hasher) With(hashFn hash.Hash32) Hasher {
	return Hasher{
		hasher:  hashFn,
		buf:     hasher.buf,
		encoder: hasher.encoder,
	}
}

func NewHasher(hasher hash.Hash32) Hasher {
	buf := bytes.NewBuffer(make([]byte, 0, 4))
	hash := Hasher{
		hasher:  hasher,
		buf:     buf,
		encoder: gob.NewEncoder(buf),
	}
	return hash
}

func (hasher *Hasher) reset() {
	hasher.buf.Reset()
	hasher.hasher.Reset()
}

func (hasher *Hasher) Calculate(key any) (uint32, error) {
	defer hasher.reset()

	err := hasher.encoder.Encode(key)
	if err != nil {
		return 0, err
	}

	_, err = hasher.hasher.Write(hasher.buf.Bytes())
	if err != nil {
		return 0, err
	}

	return hasher.hasher.Sum32(), nil
}
