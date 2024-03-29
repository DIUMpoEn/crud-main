package test

import (
	"crud/models"
	"encoding/json"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"testing"
)

var url1 = "http://127.0.0.1:5555"

func TestGetUser(t *testing.T) {
	output, err := GetUser("3")
	if err != nil {
		t.Fatalf("failed getting user: %v", err)
	}
	log.Printf("Output: %v", output)
}

func GetUser(id string) ([]models.Account, error) {
	output := []models.Account{}

	url := fmt.Sprintf("%s/get/%s", url1, id)
	log.Printf("URL: %v", url)
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetResult(&output).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code wrong. status: %d. body: %s", resp.StatusCode(), resp.String())
	}

	return output, nil
}
func TestDeleteUser(t *testing.T) {
	///output, err := DeleteUser("3")
	err := DeleteUser("3")
	if err != nil {
		t.Fatalf("failed deleting user: %v", err)
	}
	//log.Printf("Output: %v", output)
}

// /func DeleteUser(id string) ([]models.Account, error) {
func DeleteUser(id string) error {
	///output := []models.Account{}

	url := fmt.Sprintf("%s/delete/%s", url1, id)
	log.Printf("URL: %v", url)
	resp, err := resty.New().R().
		Delete(url)
	///SetHeader("Content-Type", "application/json").
	///SetResult(&output).
	///Get(url)
	///if err != nil {
	///	return nil, err
	///}
	///if resp.StatusCode() != 200 {
	///	return nil, fmt.Errorf("status code wrong. status: %d. body: %s", resp.StatusCode(), resp.String())
	///}

	///	return output, nil
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil
	}
	if resp.StatusCode() == 200 {
		fmt.Println("Запись успешна удалена.")
	} else {
		fmt.Println("Ошибка удаления записи. Код статуса:", resp.StatusCode())
	}
	return nil
}
func TestAddUser(t *testing.T) {
	newRecordData := models.Account{
		Id:       5,
		Login:    "Vladimir",
		Password: "555",
		Email:    "qqq@mail.ru",
	}

	err := AddUser(newRecordData)
	if err != nil {
		t.Fatalf("failed adding user: %v", err)
	}
	//log.Printf("Output: %v", output)
}

func AddUser(inputDB models.Account) error {
	// Сериализация данных в формат JSON
	data, err := json.Marshal(inputDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputDB)
	fmt.Println(data)

	m := make(map[string]string)
	m["login"] = inputDB.Login
	m["password"] = inputDB.Password
	m["email"] = inputDB.Email

	url := fmt.Sprintf("%s/insert", url1)
	log.Printf("URL: %v", url)
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(m).
		Post(url)

	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil
	}
	if resp.StatusCode() == 200 {
		fmt.Println("Запись успешна добавлена.")
	} else {
		fmt.Println("Ошибка добавления записи. Код статуса:", resp.StatusCode())
	}
	return nil
}

func TestUpdateUser(t *testing.T) {
	newRecordData := models.Account{
		Id:       11,
		Login:    "Антон",
		Password: "2020",
		Email:    "yma@gmail.ru",
	}

	err := UpdateUser(newRecordData)
	if err != nil {
		t.Fatalf("failed adding user: %v", err)
	}
	//log.Printf("Output: %v", output)
}

func UpdateUser(inputDB models.Account) error {
	// Сериализация данных в формат JSON
	//data, err := json.Marshal(inputDB)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(inputDB)
	//fmt.Println(data)

	m := make(map[string]string)
	m["id"] = strconv.Itoa(inputDB.Id)
	m["login"] = inputDB.Login
	m["password"] = inputDB.Password
	m["email"] = inputDB.Email

	for key, value := range m {
		fmt.Printf("Ключ: %s, Значение: %s\n", key, value)
	}

	url := fmt.Sprintf("%s/update", url1)
	log.Printf("URL: %v", url)
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(m).
		Put(url)

	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil
	}
	if resp.StatusCode() == 200 {
		fmt.Println("Запись успешн обновлена.")
	} else {
		fmt.Println("Ошибка обновления записи. Код статуса:", resp.StatusCode())
	}
	return nil
}
