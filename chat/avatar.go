package main

import (
	"errors"
	"io/ioutil"
	"path"
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

// UseGravatar is a handy GravatarAvatar type but remains of nil value
var UseGravatar GravatarAvatar

// GetAvatarURL is method of struct GravatarAvatar to implement interface Avatar
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar is a empty struct
type FileSystemAvatar struct{}

// UseFileSystemAvatar is a handy FileSystemAvatar type but remains of nil value
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is method of struct FileSystemAvatar to implement interface Avatar
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
