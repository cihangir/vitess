// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package topo

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/youtube/vitess/go/bson"
	"github.com/youtube/vitess/go/bytes2"
)

// MarshalBson bson-encodes TabletType.
func (tabletType TabletType) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	if key == "" {
		lenWriter := bson.NewLenWriter(buf)
		defer lenWriter.Close()
		key = bson.MAGICTAG
	}
	bson.EncodeString(buf, key, string(tabletType))
}

// UnmarshalBson bson-decodes into TabletType.
func (tabletType *TabletType) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	if kind == bson.EOO {
		bson.Next(buf, 4)
		kind = bson.NextByte(buf)
		bson.ReadCString(buf)
	}
	*tabletType = TabletType(bson.DecodeString(buf, kind))
}
