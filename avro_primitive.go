// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     data.avsc
 */

package gosercomp

import (
	"io"

	"github.com/actgardner/gogen-avro/vm/types"
)

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func writeArrayString(r []string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeString(e, w)
		if err != nil {
			return err
		}
	}
	return writeLong(0, w)
}

func writeAvroColorGroup(r *AvroColorGroup, w io.Writer) error {
	var err error
	err = writeInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeArrayString(r.Colors, w)
	if err != nil {
		return err
	}

	return nil
}

func writeInt(r int32, w io.Writer) error {
	downShift := uint32(31)
	encoded := uint64((uint32(r) << 1) ^ uint32(r>>downShift))
	const maxByteSize = 5
	return encodeInt(w, maxByteSize, encoded)
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

type ArrayStringWrapper []string

func (_ *ArrayStringWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayStringWrapper) Finalize()                        {}
func (_ *ArrayStringWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r *ArrayStringWrapper) AppendArray() types.Field {
	var v string

	*r = append(*r, v)
	return (*types.String)(&(*r)[len(*r)-1])
}