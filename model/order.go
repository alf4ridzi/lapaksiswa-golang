package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/alf4ridzi/lapaksiswa-golang/database"
)

type Order struct {
	OrderID  string
	ProdukID string
	Domain   string
	Username string
	Alamat   string
	Qty      int32
	Status   string
	Harga    int64
	Metode   string
	Note     string
	NoHP     int64
}

type OrderModel struct {
	DB      *sql.DB
	table   string
	columns string
}

func NewOrderModel() *OrderModel {
	// allowed columns
	var columnsAllowed = []string{
		"order_id",
		"produk_id",
		"domain",
		"username",
		"alamat",
		"qty",
		"status",
		"harga",
		"metode",
		"note",
		"no_hp",
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(fmt.Sprintf("Kesalahan database : %v", err))
	}

	columns := strings.Join(columnsAllowed, ", ")

	return &OrderModel{
		DB:      db,
		table:   "transaksi",
		columns: columns,
	}
}

func (o *OrderModel) GetOrderToko(domain string) ([]Order, error) {
	return nil, nil
}

func (o *OrderModel) GetTransaksiHariIni(domain string) (map[string]any, error) {
	var (
		Total   int64 = 0
		Sukses  int32 = 0
		Proses  int32 = 0
		Pending int32 = 0
		Gagal   int32 = 0
	)

	query := fmt.Sprintf("SELECT harga, status FROM %s WHERE domain = ? AND DATE(created_at) = CURDATE()", o.table)
	rows, err := o.DB.Query(query, domain)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			harga  int64 = 0
			status string
		)

		if err = rows.Scan(
			&harga,
			&status,
		); err != nil {
			return nil, err
		}

		switch status {
		case "waiting":
			Pending += 1
		case "proses":
			Proses += 1
		case "sukses":
			Sukses += 1
			Total += harga
		default:
			Gagal += 1
		}

	}

	data := map[string]any{
		"total":   FormatToIDR(Total),
		"proses":  Proses,
		"sukses":  Sukses,
		"pending": Pending,
		"gagal":   Gagal,
	}

	return data, nil
}

func (o *OrderModel) GetTransaksiKemarin(domain string) (map[string]any, error) {
	var (
		Total   int64 = 0
		Sukses  int32 = 0
		Proses  int32 = 0
		Pending int32 = 0
		Gagal   int32 = 0
	)

	query := fmt.Sprintf("SELECT harga, status FROM %s WHERE domain = ? AND DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)", o.table)
	rows, err := o.DB.Query(query, domain)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			harga  int64 = 0
			status string
		)

		if err = rows.Scan(
			&harga,
			&status,
		); err != nil {
			return nil, err
		}

		switch status {
		case "waiting":
			Pending += 1
		case "proses":
			Proses += 1
		case "sukses":
			Sukses += 1
			Total += harga
		default:
			Gagal += 1
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	data := map[string]any{
		"total":   FormatToIDR(Total),
		"proses":  Proses,
		"sukses":  Sukses,
		"pending": Pending,
		"gagal":   Gagal,
	}

	return data, nil
}

func (o *OrderModel) GetTotalOmset(domain string) (string, error) {
	query := fmt.Sprintf("SELECT SUM(harga) FROM %s WHERE domain = ? AND status = 'sukses'", o.table)
	var omset int64
	var harga sql.NullString

	if err := o.DB.QueryRow(query, domain).Scan(&harga); err != nil {
		if err == sql.ErrNoRows {
			return "0", nil
		}

		return "", err
	}

	if harga.Valid {
		parseOmset, err := strconv.ParseInt(harga.String, 10, 64)
		if err != nil {
			return "", err
		}

		omset = parseOmset
	} else {
		omset = 0
	}

	TotalOmset := FormatToIDR(omset)

	return TotalOmset, nil
}

func (o *OrderModel) GetOmsetBulanan(domain string) (string, error) {
	query := fmt.Sprintf("SELECT SUM(harga) FROM %s WHERE domain = ? AND status = 'sukses' AND created_at BETWEEN DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND NOW()", o.table)
	var omset int64
	var tmp sql.NullString

	if err := o.DB.QueryRow(query, domain).Scan(&tmp); err != nil {
		if err == sql.ErrNoRows {
			return "0", nil
		}

		return "", err
	}

	if tmp.Valid {
		parseOmset, err := strconv.ParseInt(tmp.String, 10, 64)
		if err != nil {
			return "", err
		}

		omset = parseOmset
	} else {
		omset = 0
	}

	TotalOmset := FormatToIDR(omset)

	return TotalOmset, nil
}

func (o OrderModel) GetSaldo(domain string) (int, error) {
	saldo := 0

	return saldo, nil
}

func (o *OrderModel) GetTotalTransaksi(domain string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(id) AS num FROM %s WHERE domain = ?", o.table)
	row := o.DB.QueryRow(query, domain)

	var total int64

	if err := row.Scan(&total); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	return total, nil
}
