// Copyright 2014 The go-ethereum Authors && Copyright 2015 shift Authors
// This file is part of the shift library.
//
// The shift library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The shift library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the shift library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

func NewDefaultGenesisReader() (io.Reader, error) {
	return gzip.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(defaultGenesisBlock)))
}

const defaultGenesisBlock = "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"
