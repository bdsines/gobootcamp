package repos

type DBSequence struct {
	Name     string `json:"name" bson:"_id"`
	Sequence int    `json:"sequence" bson:"sequence"`
}

const (
	COUNTERS = "todocounters"
)

func getNextSequence(name string) (int, error) {
	var seq DBSequence

	err := db.C(COUNTERS).FindId(name).One(&seq)

	if err != nil {
		seq.Name = name
		seq.Sequence = 1
		err := db.C(COUNTERS).Insert(seq)
		return seq.Sequence, err
	}
	seq.Sequence++
	err = db.C(COUNTERS).UpdateId(name, &seq)

	return seq.Sequence, err
}
