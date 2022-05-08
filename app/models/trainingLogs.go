package models

import (
	"log"
	"time"
)

type TrainingLog struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// trainingLogsテーブルにトレーニングログを作成
func (u *User) CreateTrainingLog(content string) (err error) {
	cmd := `insert into trainingLogs (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err

}

// トレーニングログ情報を取得
func GetTrainingLog(id int) (trainingLog TrainingLog, err error) {
	cmd := `select id, content, user_id, created_at from trainingLogs where id = ?`
	trainingLog = TrainingLog{}
	err = Db.QueryRow(cmd, id).Scan(
		&trainingLog.ID,
		&trainingLog.Content,
		&trainingLog.UserID,
		&trainingLog.CreatedAt,
	)

	return trainingLog, err
}

func GetTrainingLogs() (trainingLogs []TrainingLog, err error) {
	cmd := `select id, content, user_id, created_at from trainingLogs`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var trainingLog TrainingLog
		err = rows.Scan(
			&trainingLog.ID,
			&trainingLog.Content,
			&trainingLog.UserID,
			&trainingLog.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		trainingLogs = append(trainingLogs, trainingLog)
	}

	rows.Close()

	return trainingLogs, err

}

func (u *User) GetTrainingLogsByUser() (trainingLogs []TrainingLog, err error) {
	cmd := `select id, content, user_id, created_at from trainingLogs where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var trainingLog TrainingLog
		err = rows.Scan(
			&trainingLog.ID,
			&trainingLog.Content,
			&trainingLog.UserID,
			&trainingLog.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		trainingLogs = append(trainingLogs, trainingLog)
	}

	rows.Close()

	return trainingLogs, err
}
