package uuid

import (
	"errors"

	googleUUID "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type UUID struct {
	googleUUID.UUID
}

func (id UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	uuid := id.String()
	return bsontype.String, bsoncore.AppendString(nil, uuid), nil
}

func (id *UUID) UnmarshalBSONValue(bsonType bsontype.Type, bytes []byte) error {
	if bsonType != bsontype.String {
		return errors.New("UnmarshalBSONValue: uuid is not string")
	}

	s, _, ok := bsoncore.ReadString(bytes)
	if !ok {
		return errors.New("invalid bson string value")
	}

	uid, err := googleUUID.Parse(s)
	if err != nil {
		return err
	}

	id.UUID = uid

	return nil
}

func Parse(s string) (UUID, error) {
	uuid, err := googleUUID.Parse(s)
	return UUID{
		UUID: uuid,
	}, err
}

func NewString() string {
	return googleUUID.NewString()
}

func New() UUID {
	return UUID{
		UUID: googleUUID.New(),
	}
}
