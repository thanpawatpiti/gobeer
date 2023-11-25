package repositories

import "context"

func (r *Repository) DeleteBeer(ctx context.Context, id int) error {
	_, err := r.mariaDB.Exec("DELETE FROM beer WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
