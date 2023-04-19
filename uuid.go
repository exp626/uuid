package uuid

import (
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
	uid, err := googleUUID.FromBytes(bytes)
	if err != nil {
		return err
	}

	*id = UUID{
		UUID: uid,
	}
	return nil
}

func Parse(s string) (UUID, error) {
	googleUUID.NewString()
	uuid, err := googleUUID.Parse(s)
	return UUID{
		UUID: uuid,
	}, err
}

func NewString() string {
	return googleUUID.New().String()
}

func New() UUID {
	return UUID{
		UUID: googleUUID.New(),
	}
}
