package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Error string
	Ok    bool
}

type MessageResponse struct {
	Response	
	Channel string
}

type User struct {
	Id      string
	Name    string
	Deleted bool
	Profile struct {
		Email string
	}
}

type UserResponse struct {
	Response
	User  User
}

type UsersResponse struct {
	Response
	Members []User
}

// Alias for usersResponse.Members
func (response UsersResponse) Users() []User {
	return response.Members
}

const (
	POST_MESSAGE_URL = "https://slack.com/api/chat.postMessage"
	USER_INFO_URL    = "https://slack.com/api/users.info"
	USER_LIST_URL    = "https://slack.com/api/users.list"
)

type UserRequest struct {
	Token string
	User  string
}

type Client struct {
	Token string
}

func NewClient(token string) Client {
	client := Client{
		token,
	}

	return client
}

func (client Client) SendMessage(channel, text, botName string) MessageResponse {
	params := url.Values{}
	params.Add("token", client.Token)
	params.Add("channel", channel)
	params.Add("text", text)
	params.Add("username", botName)

	endpoint := fmt.Sprintf("%s?%s", POST_MESSAGE_URL, params.Encode())

	response, _ := http.Get(endpoint)
	body, _ := ioutil.ReadAll(response.Body)

	var messageResponse MessageResponse
	json.Unmarshal(body, &messageResponse)

	return messageResponse
}

func (client Client) GetUsers() UsersResponse {
	params := url.Values{}
	params.Add("token", client.Token)

	endpoint := fmt.Sprintf("%s?%s", USER_LIST_URL, params.Encode())

	response, _ := http.Get(endpoint)
	body, _ := ioutil.ReadAll(response.Body)

	var usersResponse UsersResponse
	json.Unmarshal(body, &usersResponse)

	return usersResponse
}

func (client Client) GetUser(userId string) UserResponse {
	v := url.Values{}
	v.Add("token", client.Token)
	v.Add("user", userId)

	uri := fmt.Sprintf("%s?%s", USER_INFO_URL, v.Encode())

	// TODO: handle errors
	resp, _ := http.Get(uri)
	data, _ := ioutil.ReadAll(resp.Body)

	var user UserResponse
	json.Unmarshal(data, &user)

	return user
}
