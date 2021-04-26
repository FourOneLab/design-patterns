package demo_authority

import "database/sql"

// 为了做到抽象封装具体的存储方式，
// 将 CredentialStorage 设计成接口，
// 基于接口而非具体的实现编程。

type CredentialStorage interface {
	GetPassWordByAppId(appId string) (string, error)
}

type MySQLCredentialStorage struct {
	db *sql.DB
}

func NewMySQLCredentialStorage(db *sql.DB) *MySQLCredentialStorage {
	return &MySQLCredentialStorage{db: db}
}

func (m *MySQLCredentialStorage) GetPassWordByAppId(appId string) (string, error) {
	var password string

	rows, err := m.db.Query("select password from credential where appid =?", appId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&password); err != nil {
			return "", err
		}
	}

	return password, nil
}
