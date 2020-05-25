package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

// ErrNoAvatarURL is the error that is returned when the Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing user profile pictures
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client, or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar is a empty struct
type AuthAvatar struct{}

// UseAuthAvatar is a handy AuthAvatar type but remains of nil value
var UseAuthAvatar AuthAvatar

// GetAvatarURL is method of struct AuthAvatar to implement interface Avatar
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar is a empty struct
type GravatarAvatar struct{}

// UseGravatar is a handy AuthAvatar type but remains of nil value
var UseGravatar GravatarAvatar

// GetAvatarURL is method of struct GravatarAvatar to implement interface Avatar
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}
