package model

import "books-store/utils"

type Session struct {
	Session_id string
	User_name  string
	User_id    int
}

func AddSession(sess *Session) error {
	sqlStr := "insert into sessions values($1,$2,$3)"
	_, err := utils.Db.Exec(sqlStr, sess.Session_id, sess.User_name, sess.User_id)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionByID(sessID string) (*Session, error) {
	sqlStr := "select session_id,user_name,user_id from sessions where session_id=$1"
	row := utils.Db.QueryRow(sqlStr, sessID)
	sess := &Session{}
	err := row.Scan(&sess.Session_id, &sess.User_name, &sess.User_id)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id=$1"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}
