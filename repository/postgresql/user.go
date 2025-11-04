package postgresql

import (
	"fmt"

	//"log"

	"github.com/javadifa/socialmedia/entity"
)

func (d *PostgresDB) IsEmailUnique(email string) (bool, error) {
	panic("implement me")
}

func (d *PostgresDB) IsHandleUnique(email string) (bool, error) {
	panic("implement me")
}

func (d *PostgresDB) Register(user entity.User) (entity.User, error) {

	// TODO : complete database design with postgresql and use queryRow
	iErr := d.db.QueryRow(
		`INSERT INTO users(full_name, handle, email, password, birth_date, avatar_url) VALUES
      ($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.FullName, user.Handle, user.Email, user.Password, user.BirthDate, user.AvatarURL).Scan(&user.ID)

	if iErr != nil {
		return entity.User{}, fmt.Errorf("can't execut command : %w", iErr)
	}
	return user, nil
}

// TODO: error handeling and fix the fields
func (d *PostgresDB) GetUserByHandle(handle string) (entity.User, bool, error) {
	user := entity.User{}

	err := d.db.QueryRow(`SELECT * FROM users WHERE handle = $1 `, handle).
		Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.BirthDate, &user.AvatarURL)

	if err == nil {
		return entity.User{}, false, fmt.Errorf("userservice not found")
	}
	return user, true, nil
}

func (d *PostgresDB) GetUserFeedByID(userID uint) ([]entity.Post, error) {

	rows, err := d.db.Query(`
        SELECT p.id, p.user_id, p.image_url, p.caption, p.created_at
        FROM posts p
        JOIN following_relationships fr ON p.user_id = fr.following_id
        WHERE fr.follower_id = $1
        ORDER BY p.created_at DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var p entity.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.ImageURL, &p.Caption, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
