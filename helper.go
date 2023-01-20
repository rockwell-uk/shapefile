package shapefile

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/rockwell-uk/shapefile/dbf"
)

func ReadShapeFile(shapeFile string) (uint32, *Reader) {

	// Open the shapefile for reading
	sr, err := os.Open(shapeFile)
	if err != nil {
		panic(err)
	}

	// Open the dbf for reading
	dbr, err := os.Open(strings.Replace(shapeFile, ".shp", ".dbf", 1))
	if err != nil {
		panic(err)
	}

	r, err := NewReader(sr, WithDBF(dbr))
	if err != nil {
		panic(err)
	}

	return GetRecordCount(shapeFile), r
}

func ReadShapeFileToMemory(shapeFile string) (uint32, *Reader) {

	// Open the shapefile for reading
	sr, err := ioutil.ReadFile(shapeFile)
	if err != nil {
		panic(err)
	}

	// Open the dbf for reading
	dr, err := ioutil.ReadFile(strings.Replace(shapeFile, ".shp", ".dbf", 1))
	if err != nil {
		panic(err)
	}

	r, err := NewMemoryReader(sr, WithMemoryDBF(dr))
	if err != nil {
		panic(err)
	}

	return GetRecordCount(shapeFile), r
}

func GetRecordCount(shapeFile string) uint32 {

	// Open the dbf for reading the header
	dbr, err := os.Open(strings.Replace(shapeFile, ".shp", ".dbf", 1))
	if err != nil {
		panic(err)
	}
	defer dbr.Close()

	dbf, err := dbf.NewReader(dbr)
	if err != nil {
		panic(err)
	}

	return dbf.NumberOfRecords
}
