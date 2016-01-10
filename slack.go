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
	POST_MESS_URL = "https://slack.com/api/chat.postMessage"
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
	return Client{token}
}

func (c Client) Post(channel, text, bot string) MessageResponse {
	params := url.Values{}
	params.Add("token", c.Token)
	params.Add("channel", channel)
	params.Add("text", text)
	params.Add("username", bot)

	endpoint := fmt.Sprintf("%s?%s", POST_MESS_URL, params.Encode())

	resp, _ := http.Get(endpoint)
	body, _ := ioutil.ReadAll(resp.Body)

	var reply MessageResponse
	json.Unmarshal(body, &reply)

	return reply
}

func (c Client) GetUsers() UsersResponse {
	params := url.Values{}
	params.Add("token", c.Token)
	endpoint := fmt.Sprintf("%s?%s", USER_LIST_URL, params.Encode())

	response, _ := http.Get(endpoint)
	body, _ := ioutil.ReadAll(response.Body)

	var users UsersResponse
	json.Unmarshal(body, &users)

	return users
}

func (c Client) GetUser(id string) UserResponse {
	v := url.Values{}
	v.Add("token", c.Token)
	v.Add("user", id)
	uri := fmt.Sprintf("%s?%s", USER_INFO_URL, v.Encode())

	// TODO: handle errors
	resp, _ := http.Get(uri)
	data, _ := ioutil.ReadAll(resp.Body)

	var user UserResponse
	json.Unmarshal(data, &user)

	return user
}
