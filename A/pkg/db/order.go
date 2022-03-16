package db

import "github.com/pkg/errors"

func (d *DB) SaveOrder(params map[string]interface{}) (string, error) {
	query := "INSERT INTO orders(id, price, title) VALUES(?,?,?)"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return "", errors.Wrap(err, "SaveOrder")
	}
	defer stmt.Close()

	id := params["order_id"]
	price := params["price"]
	title := params["title"]

	_, err = stmt.Exec(id, price, title)
	if err != nil {
		return "", errors.Wrap(err, "SaveOrder")
	}

	return "", nil
}
