package dictionnary

import (
	bolt "github.com/boltdb/bolt"
	logrus "github.com/sirupsen/logrus"
)

type Data struct{
	strings string
}

func (d Data) New(*bolt.DB, error) {

	logrus.Info("DICTIONNARY")
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// defer db.Close()
	// err = db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("MyBucket"))
	// 	err := b.Put([]byte("answer"), []byte("42"))
	// 	return err
	// })
	// if err != nil {
    //     return nil, fmt.Errorf("could not set up buckets, %v", err)
    // }
	// return db, nil
}

// func (e Entry) String() string{
// 	created := e.CreatedAt.Format(time.Stamp)
// 	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
// }