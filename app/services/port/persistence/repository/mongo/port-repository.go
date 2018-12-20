package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/model"
	"github.com/ic2hrmk/ship_ports/app/services/port/persistence/repository"
)

const portsCollection = "ports"

type PortRepository struct {
	db *mgo.Database
}

func NewPortRepository(db *mgo.Database) repository.PortRepository {
	return &PortRepository{db: db}
}

func (r *PortRepository) collection() *mgo.Collection {
	return r.db.C(portsCollection)
}

func (r *PortRepository) Save(record *model.Port) (*model.Port, error) {
	_, err := r.collection().Upsert(bson.M{"_id": record.PortID}, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *PortRepository) SaveBulk(records []*model.Port) error {
	bulkOperator := r.collection().Bulk()

	for i := range records {
		bulkOperator.Upsert(bson.M{"_id": records[i].PortID}, records[i])
	}

	if _, err := bulkOperator.Run(); err != nil {
		return err
	}

	return nil
}

func (r *PortRepository) FindAll(limit, offset uint64) ([]*model.Port, error) {
	return r.prepareResultList(
		r.collection().Find(bson.M{}).Limit(int(limit)).Skip(int(offset)))
}

//nolint:megacheck
func (r *PortRepository) prepareOneResult(query *mgo.Query) (*model.Port, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	record := &model.Port{}

	if count == 0 {
		return nil, repository.ErrPortNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *PortRepository) prepareResultList(query *mgo.Query) ([]*model.Port, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.Port

	if count == 0 {
		return records, nil
	}

	err = query.All(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
