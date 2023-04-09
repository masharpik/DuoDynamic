package repository

import (
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepo(db *gorm.DB) *MatchRepository {
	return &MatchRepository{
		db,
	}
}

func (r *MatchRepository) GetMatches(userId uint) (users []dataStruct.Match, err error) {
	err = r.db.Table("matches").Find(&users, "user_first_id=?", userId).Limit(30).Error
	return
}

func (r *MatchRepository) GetUser(userId uint) (user match.UserRes, err error) {
	err = r.db.Table("users u").Select("u.email, ui.name, p.photo").
		Where("u.id=?", userId).
		Joins("Join user_infos ui on u.id = ui.user_id").
		Joins("join user_photos p on u.id=p.user_id").
		Find(&user).Error
	return
}

func (r *MatchRepository) GetIdByEmail(email string) (uint, error) {
	var user dataStruct.User
	err := r.db.First(&user, "email = ?", email).Error
	return user.Id, err
}

func (r *MatchRepository) GetIdReaction(reaction string) (uint, error) {
	var react dataStruct.Reaction
	err := r.db.First(&react, "reaction = ?", reaction).Error
	return react.Id, err
}

func (r *MatchRepository) AddHistoryRow(row dataStruct.UserHistory) error {
	err := r.db.Create(&row).Error
	return err
}

func (r *MatchRepository) GetUserReaction(userId, userToId uint) (dataStruct.UserReaction, error) {
	var react dataStruct.UserReaction
	err := r.db.First(&react, "user_id = ? AND user_from_id=?", userId, userToId).Error
	return react, err
}

func (r *MatchRepository) AddUserReaction(row dataStruct.UserReaction) error {
	err := r.db.Create(&row).Error
	return err
}

func (r *MatchRepository) DeleteUserReaction(rowId uint) error {
	err := r.db.First(&dataStruct.UserReaction{}, "id =?", rowId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserReaction{}, "id=?", rowId).Error
	return err
}

func (r *MatchRepository) AddMatchRow(row dataStruct.Match) error {
	err := r.db.Create(&row).Error
	return err
}
