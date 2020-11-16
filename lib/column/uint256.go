package column

import (
	"github.com/ClickHouse/clickhouse-go/lib/binary"
)

type UInt256 struct{ base }

func (UInt256) Read(decoder *binary.Decoder, isNull bool) (interface{}, error) {
	v, err := decoder.Fixed(32)
	if err != nil {
		return []byte{}, err
	}
	return v, nil
}

func (u *UInt256) Write(encoder *binary.Encoder, v interface{}) error {
	switch v := v.(type) {
	case []byte:
		if _, err := encoder.Write(v); err != nil {
			return err
		}
		return nil
	}

	return &ErrUnexpectedType{
		T:      v,
		Column: u,
	}
}
