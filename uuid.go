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
