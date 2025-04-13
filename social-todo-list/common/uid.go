package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/akamensky/base58"
	"strconv"
)

// UID is method to generate a virtual unique identifier for whole system
// its structure contains 62 bits: LocalID - ObjectType - ShardID
// 32 bits for Local ID, max (2^32) - 1
// 10 bits for Object Type (max 1024)
// 18 bits for Shard ID
type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{localID, objectType, shardID}
}

// Shard: 1, Object: 1, ID: 1 => 0001 0001 0001
// 1 << 8 = 0001 0000 0000
// 1 << 4 = 		1 0000
// 1 << 0 = 			 1
// => 0001 0001 0001

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%d", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetObjectType() int {
	return uid.objectType
}

func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 10) > uid {
		return UID{}, errors.New("wrong uid")
	}

	// x = 1110 1110 0101 => x >> 4 = 1110 1110 & 0000 1111 = 0000 1110
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

// Value implements the sql/driver.Valuer interface for database storage.
func (uid UID) Value() (driver.Value, error) {
	return uid.String(), nil
}

// Scan implements the sql.Scanner interface for database retrieval.
func (uid *UID) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("Scan: unable to convert value to string")
	}

	decodedUID, err := DecomposeUID(str)
	if err != nil {
		return err
	}

	*uid = decodedUID
	return nil
}

func DecodeUID(encoded string) (UID, error) {
	var uid UID

	decodedBytes, err := base58.Decode(encoded)
	if err != nil {
		return uid, fmt.Errorf("failed to decode base58: %w", err)
	}

	// Convert decoded bytes (decimal string) to uint64
	valStr := string(decodedBytes)
	val, err := strconv.ParseUint(valStr, 10, 64)
	if err != nil {
		return uid, fmt.Errorf("invalid decoded value: %w", err)
	}

	uid.localID = uint32(val >> 28)
	uid.objectType = int((val >> 18) & 0x3FF) // 10 bits for objectType
	uid.shardID = uint32(val & 0x3FFFF)       // 18 bits for shardID

	return uid, nil
}
