package model

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/alf4ridzi/lapaksiswa-golang/config/cookie"
	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type User struct {
	Username     string
	Nama         string
	TanggalLahir string
	Email        string
	NoHP         int64
	Role         string
	Foto         string
	Alamat       string
	Kelas        string
	Password     string
	Token_Reset  string
	JenisKelamin string
	Created_at   string
	Last_online  string
	Updated_at   string
}

type UserModel struct {
	DB      *sql.DB
	table   string
	columns string
}

// NewUserModel initializes the UserModel with necessary properties.
func NewUserModel() *UserModel {
	columnsAllowed := []string{
		"username",
		"nama",
		"tanggal_lahir",
		"email",
		"no_hp",
		"role",
		"foto",
		"alamat",
		"kelas",
		"password",
		"token_reset_password",
		"jenis_kelamin",
		"created_at",
		"last_online",
		"updated_at",
	}

	columns := columnsAllowed[0]
	for _, col := range columnsAllowed[1:] {
		columns = fmt.Sprintf("%s, %s", columns, col)
	}

	db := database.InitDatabase()
	if db == nil {
		panic("Failed to initialize database")
	}

	return &UserModel{
		DB:      db,
		table:   "user",
		columns: columns,
	}
}

func Md5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func (p *UserModel) CheckDataExist(key, value string) (bool, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", key, p.table, key)
	row := p.DB.QueryRow(query, value)

	var result string
	if err := row.Scan(&result); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *UserModel) ValidasiLogin(w http.ResponseWriter, email, password string) (bool, error) {
	password = Md5(password)
	query := fmt.Sprintf("SELECT username, role FROM %s WHERE email = ? AND password = ?", p.table)
	row := p.DB.QueryRow(query, email, password)

	var username, role string
	if err := row.Scan(&username, &role); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cookie.SetFlashCookie(w, "error", "Invalid email or password.")
			return false, nil
		}
		cookie.SetFlashCookie(w, "error", "Login validation failed.")
		return false, err
	}

	cookieData := map[string]string{
		"isLogin":  "true",
		"username": username,
		"role":     role,
	}
	cookie.SetCookie(w, cookieData)
	return true, nil
}

func (p *UserModel) ValidasiRegister(username, nama, email, nohp, password string) (bool, string) {
	if username == "" || nama == "" || email == "" || nohp == "" || password == "" {
		return false, "All fields are required."
	}

	if isUser, err := p.CheckDataExist("username", username); err != nil {
		return false, err.Error()
	} else if isUser {
		return false, "Username is already registered."
	}

	if isEmail, err := p.CheckDataExist("email", email); err != nil {
		return false, err.Error()
	} else if isEmail {
		return false, "Email is already registered."
	}

	if isNoHp, err := p.CheckDataExist("no_hp", nohp); err != nil {
		return false, err.Error()
	} else if isNoHp {
		return false, "Phone number is already registered."
	}

	password = Md5(password)
	query := "INSERT INTO user (username, nama, email, no_hp, password) VALUES (?, ?, ?, ?, ?)"
	if _, err := p.DB.Exec(query, username, nama, email, nohp, password); err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (p *UserModel) GetUser(username string) (*User, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ?", p.columns, p.table)
	row := p.DB.QueryRow(query, username)

	var user User
	var noHP sql.NullInt64
	var foto, alamat, TanggalLahir, kelas, tokenReset, jenisKelamin, createdAt, lastOnline, updatedAt sql.NullString

	err := row.Scan(
		&user.Username,
		&user.Nama,
		&TanggalLahir,
		&user.Email,
		&noHP,
		&user.Role,
		&foto,
		&alamat,
		&kelas,
		&user.Password,
		&tokenReset,
		&jenisKelamin,
		&createdAt,
		&lastOnline,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user.TanggalLahir = TanggalLahir.String
	user.NoHP = noHP.Int64
	user.Foto = foto.String
	user.Alamat = alamat.String
	user.Kelas = kelas.String
	user.Token_Reset = tokenReset.String
	user.JenisKelamin = jenisKelamin.String
	user.Created_at = createdAt.String
	user.Last_online = lastOnline.String
	user.Updated_at = updatedAt.String

	return &user, nil
}

func (p *UserModel) UpdateProfileUser(username string, nama string, tanggalLahir string, jenisKelamin string, email string, noHP int64) error {

	query := fmt.Sprintf("UPDATE %s SET nama = ?, tanggal_lahir = ?, jenis_kelamin = ?, email = ?, no_hp ? WHERE username = ?", p.table)
	if _, err := p.DB.Exec(query, nama, tanggalLahir, jenisKelamin, email, noHP, username); err != nil {
		return err
	}

	return nil
}
