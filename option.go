package shapefile

import (
	"bytes"
	"io"

	"github.com/rockwell-uk/shapefile/dbf"
)

// Option allows for options to be specified when calling NewReader.
type Option func(r *Reader) error

// WithDBF allows the specification of a dbf file that
// contains the associated attributes for each of the shapefile records.
func WithDBF(r io.Reader) Option {
	return func(rr *Reader) error {
		dr, err := dbf.NewReader(r)
		if err != nil {
			return err
		}
		rr.dbf = dr
		return nil
	}
}

func WithMemoryDBF(r []byte) Option {
	return func(rr *Reader) error {
		dr, err := dbf.NewReader(bytes.NewReader(r))
		if err != nil {
			return err
		}
		rr.dbf = dr
		return nil
	}
}
