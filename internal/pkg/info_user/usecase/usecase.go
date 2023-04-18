package usecase

import (
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/info_user"
)

type InfoUseCase struct {
	userRepo info_user.IRepositoryInfo
}

func NewInfoUseCase(userRepo info_user.IRepositoryInfo) *InfoUseCase {
	return &InfoUseCase{
		userRepo: userRepo,
	}
}

func (iu *InfoUseCase) AddInfo(userId uint, info info_user.InfoStruct) error {
	userInfo := dataStruct.UserInfo{
		UserId:      userId,
		Name:        info.Name,
		Description: info.Description,
		Sex:         info.Sex,
	}

	cityId, err := iu.userRepo.GetCityId(info.City)
	if err != nil {
		return err
	}
	zodiacId, err := iu.userRepo.GetZodiacId(info.Zodiac)
	if err != nil {
		return err
	}
	educationId, err := iu.userRepo.GetEducationId(info.Education)
	if err != nil {
		return err
	}
	jobId, err := iu.userRepo.GetJobId(info.Job)
	if err != nil {
		return err
	}

	userInfo.Zodiac = zodiacId
	userInfo.CityId = cityId
	userInfo.Education = educationId
	userInfo.Job = jobId

	err = iu.userRepo.AddInfoUser(&userInfo)
	if err != nil {
		return err
	}
	return nil
} //ok

func (iu *InfoUseCase) GetInfo(userId uint) (userInfo info_user.InfoStructAnswer, err error) {
	userInfoTemp, err := iu.userRepo.GetUserInfo(userId)
	if err != nil {
		return
	}
	userInfo = info_user.InfoStructAnswer{
		Email:       userInfoTemp.Email,
		Name:        userInfoTemp.Name,
		Sex:         userInfoTemp.Sex,
		Job:         userInfoTemp.Job,
		Education:   userInfoTemp.Education,
		Zodiac:      userInfoTemp.Zodiac,
		City:        userInfoTemp.City,
		Description: userInfoTemp.Description,
	}
	age, err := iu.userRepo.GetAge(userId)
	if err != nil {
		return
	}
	userInfo.Age = age
	avatar, err := iu.userRepo.GetAvatar(userId) //fix
	if err != nil {
		return
	}
	userInfo.Photos = append(userInfo.Photos, avatar) //fix
	photos, err := iu.userRepo.GetUserPhoto(userId)
	if err != nil {
		return
	}
	userInfo.Photos = append(userInfo.Photos, photos...)
	return
} //fix

func (iu *InfoUseCase) ChangeInfo(userId uint, infoInp info_user.InfoChange) (info_user.InfoStructAnswer, error) {
	userInfo := dataStruct.UserInfo{
		Sex:         infoInp.Sex,
		Description: infoInp.Description,
		Name:        infoInp.Name,
		UserId:      userId,
	}

	cityId, err := iu.userRepo.GetCityId(infoInp.City)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.CityId = cityId

	zodiacId, err := iu.userRepo.GetZodiacId(infoInp.Zodiac)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Zodiac = zodiacId

	educationId, err := iu.userRepo.GetEducationId(infoInp.Education)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Education = educationId

	jobId, err := iu.userRepo.GetJobId(infoInp.Job)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	userInfo.Job = jobId

	err = iu.userRepo.ChangeInfo(&userInfo)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	result, err := iu.GetInfo(userId)
	if err != nil {
		return info_user.InfoStructAnswer{}, err
	}
	return result, nil
} //ok

func (iu *InfoUseCase) AddHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	for _, hashtag := range hashtagInp.Hashtag {
		hashtagId, err := iu.userRepo.GetHashtagId(hashtag)
		if err != nil {
			return err
		}
		var userHashtag dataStruct.UserHashtag
		userHashtag.HashtagId = hashtagId
		userHashtag.UserId = userId
		err = iu.userRepo.AddUserHashtag(userHashtag)
		if err != nil {
			return err
		}
	}

	return nil
}

func (iu *InfoUseCase) GetUserHashtags(userId uint) (info_user.HashtagInp, error) {
	hashtags, err := iu.userRepo.GetUserHashtags(userId)
	if err != nil {
		return info_user.HashtagInp{}, err
	}

	var result info_user.HashtagInp
	for _, hashtagId := range hashtags {
		hashtag, err := iu.userRepo.GetHashtagById(hashtagId.HashtagId)
		if err != nil {
			return info_user.HashtagInp{}, err
		}

		result.Hashtag = append(result.Hashtag, hashtag)
	}
	return result, nil
}

func (iu *InfoUseCase) ChangeUserHashtags(userId uint, hashtagInp info_user.HashtagInp) error {
	hashtagsBD, err := iu.userRepo.GetUserHashtags(userId)
	if err != nil {
		return err
	}
	var hashtagSlice []string
	for _, hashtagId := range hashtagsBD {
		reason, err := iu.userRepo.GetHashtagById(hashtagId.HashtagId)
		if err != nil {
			return err
		}
		hashtagSlice = append(hashtagSlice, reason)
	}

	for _, hashtag := range hashtagInp.Hashtag {
		if !Contains(hashtagSlice, hashtag) {
			var hashtagAdd dataStruct.UserHashtag
			hashtagId, err := iu.userRepo.GetHashtagId(hashtag)
			if err != nil {
				return err
			}
			hashtagAdd.UserId = userId
			hashtagAdd.HashtagId = hashtagId
			err = iu.userRepo.AddUserHashtag(hashtagAdd)
		}
	}

	for _, hashtag := range hashtagSlice {
		if !Contains(hashtagInp.Hashtag, hashtag) {
			hashtagId, err := iu.userRepo.GetHashtagId(hashtag)
			if err != nil {
				return err
			}
			err = iu.userRepo.DeleteUserHashtag(userId, hashtagId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Contains(s []string, elem string) bool {
	for _, elemS := range s {
		if elem == elemS {
			return true
		}
	}
	return false
}
