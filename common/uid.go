package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

// 32 bits for Local ID, max (2^32) - 1
// 10 bits for Object Type
// 18 bits for Shard ID
const (
	objectTypeBits = 10
	shardIDBits    = 18

	objectTypeMask = (1 << objectTypeBits) - 1
	shardIDMask    = (1 << shardIDBits) - 1
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) *UID {
	return &UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (uid *UID) String() string {
	val := uint64(uid.localID)<<28 + uint64(uid.objectType)<<18 + uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid *UID) LocalID() uint32 {
	return uid.localID
}

func (uid *UID) ObjectType() int {
	return uid.objectType
}

func (uid *UID) ShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}

	if (1 << 10) > uid {
		return UID{}, errors.New("uid too short")
	}

	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & objectTypeMask),
		shardID:    uint32(uid & shardIDMask),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(b []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(b), "\"", "", -1))
	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.objectType = decodeUID.objectType
	uid.shardID = decodeUID.shardID

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}

	return int64(uid.localID), nil
}

func (uid *UID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var val uint64

	switch t := src.(type) {
	case int:
		val = uint64(t)
	case int8:
		val = uint64(t)
	case int16:
		val = uint64(t)
	case int32:
		val = uint64(t)
	case int64:
		if t < 0 {
			return errors.New("UID cannot be negative")
		}
		val = uint64(t)
	case uint:
		val = uint64(t)
	case uint8:
		val = uint64(t)
	case uint16:
		val = uint64(t)
	case uint32:
		val = uint64(t)
	case uint64:
		val = t
	default:
		return fmt.Errorf("unsupported UID source type: %T", src)
	}

	uid.localID = uint32(val >> 28)
	uid.objectType = int((val >> 18) & objectTypeMask)
	uid.shardID = uint32(val & shardIDMask)

	return nil
}
