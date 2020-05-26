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
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars is a valid Avatar implementation
type TryAvatars []Avatar

// GetAvatarURL takes a turn in trying to get a URL for a user
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar is a empty struct
type AuthAvatar struct{}

// UseAuthAvatar is a handy AuthAvatar type but remains of nil value
var UseAuthAvatar AuthAvatar

// GetAvatarURL is method of struct AuthAvatar to implement interface Avatar
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

// GravatarAvatar is a empty struct
type GravatarAvatar struct{}

// UseGravatar is a handy GravatarAvatar type but remains of nil value
var UseGravatar GravatarAvatar

// GetAvatarURL is method of struct GravatarAvatar to implement interface Avatar
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar is a empty struct
type FileSystemAvatar struct{}

// UseFileSystemAvatar is a handy FileSystemAvatar type but remains of nil value
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL is method of struct FileSystemAvatar to implement interface Avatar
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
