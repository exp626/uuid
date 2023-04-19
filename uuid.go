package uuid

import (
	googleUUID "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type UUID struct {
	googleUUID.UUID
}

func (r *UUID) MarshalBSON() ([]byte, error) {
	uuid := r.String()
	return bson.Marshal(uuid)
}

func (r *UUID) UnmarshalBSON(data []byte) error {
	var strUUID string
	err := bson.Unmarshal(data, strUUID)
	if err != nil {
		return err
	}

	uuid, err := googleUUID.Parse(strUUID)
	if err != nil {
		return err
	}

	r.UUID = uuid
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
