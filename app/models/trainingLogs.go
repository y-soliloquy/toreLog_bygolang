package models

import (
	"log"
	"time"
)

type TrainingLog struct {
	ID           int
	Content      string
	Satisfaction string
	Weather      string
	UserID       int
	CreatedAt    time.Time
}

// trainingLogsテーブルにトレーニングログを作成
func (u *User) CreateTrainingLog(content string, satisfaction string, weather string) (err error) {
	cmd := `insert into trainingLogs (
		content,
		satisfaction,
		weather,
		user_id,
		created_at) values (?, ?, ?, ?, ?)`

	// 書式指定
	const format = "2006-01-02 15:04:05"

	_, err = Db.Exec(cmd, content, satisfaction, weather, u.ID, time.Now().Format(format))
	if err != nil {
		log.Fatalln(err)
	}

	return err

}

// トレーニングログ情報を取得（単数）
func GetTrainingLog(id int) (trainingLog TrainingLog, err error) {
	cmd := `select id, content, satisfaction, weather, user_id, created_at from trainingLogs where id = ?`
	trainingLog = TrainingLog{}
	err = Db.QueryRow(cmd, id).Scan(
		&trainingLog.ID,
		&trainingLog.Content,
		&trainingLog.Satisfaction,
		&trainingLog.Weather,
		&trainingLog.UserID,
		&trainingLog.CreatedAt,
	)

	return trainingLog, err
}

// トレーニングログ情報を取得（複数）
func GetTrainingLogs() (trainingLogs []TrainingLog, err error) {
	cmd := `select id, content, satisfaction, weather, user_id, created_at from trainingLogs`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var trainingLog TrainingLog
		err = rows.Scan(
			&trainingLog.ID,
			&trainingLog.Content,
			&trainingLog.Satisfaction,
			&trainingLog.Weather,
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

// ユーザーIDで絞り込んでトレーニングログ情報を取得
func (u *User) GetTrainingLogsByUser() (trainingLogs []TrainingLog, err error) {
	cmd := `select id, content, satisfaction, weather, user_id, created_at from trainingLogs where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var trainingLog TrainingLog
		err = rows.Scan(
			&trainingLog.ID,
			&trainingLog.Content,
			&trainingLog.Satisfaction,
			&trainingLog.Weather,
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

// トレーニングログ情報を更新
func (t *TrainingLog) UpdateTrainingLog() (err error) {
	cmd := `update trainingLogs set content = ?, satisfaction = ?, weather = ?, user_id =? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.Satisfaction, t.Weather, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

// トレーニングログ情報を削除
func (t *TrainingLog) DeleteTrainingLog() (err error) {
	cmd := `delete from trainingLogs where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
