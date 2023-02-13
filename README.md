## forum

### Objectives

This project consists in creating a web forum that allows :

- communication between users.
- associating categories to posts.
- liking and disliking posts and comments.
- filtering posts.

#### SQLite

In order to store the data in your forum (like users, posts, comments, etc.) you will use the database library SQLite.

#### Authentication

In this segment the client must be able to `register` as a new user on the forum, by inputting their credentials. You also have to create a `login session` to access the forum and be able to add posts and comments.

You should use cookies to allow each user to have only one opened session. Each of this sessions must contain an expiration date. It is up to you to decide how long the cookie stays "alive". The use of UUID is a Bonus task.

Instructions for user registration:

- Must ask for email
  - When the email is already taken return an error response.
- Must ask for username
- Must ask for password
  - The password must be encrypted when stored (this is a Bonus task)

The forum must be able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and, if the password is not the same, it will return an error response.

#### Communication

In order for users to communicate between each other, they will have to be able to create posts and comments.

- Only registered users will be able to create posts and comments.
- When registered users are creating a post they can associate one or more categories to it.
- The implementation and choice of the categories is up to you.
- The posts and comments should be visible to all users (registered or not).
- Non-registered users will only be able to see posts and comments.

#### Likes and Dislikes

Only registered users will be able to like or dislike posts and comments.

The number of likes and dislikes should be visible by all users (registered or not).

#### Filter

You need to implement a filter mechanism, that will allow users to filter the displayed posts by :

- categories
- created posts
- liked posts

You can look at filtering by categories as subforums. A subforum is a section of an online forum dedicated to a specific topic.

#### Forum-image-upload
In `forum image upload`, registered users have the possibility to create a post containing an image as well as text.

- When viewing the post, users and guests should see the image associated to it.

There are several extensions for images like: JPEG, SVG, PNG, GIF, etc. In this project you have to handle at least JPEG, PNG and GIF types.

#### Forum-advenced-features
Note that the last two are only available for registered users and must refer to the logged in user.
- You will have to create a way to notify users when their posts are :

  - liked/disliked
  - commented

- You have to create an activity page that tracks the user own activity. In other words, a page that :

  - Shows the user created posts
  - Shows where the user left a like or a dislike
  - Shows where and what the user has been commenting. For this, the comment will have to be shown, as well as the post commented

- You have to create a section where you will be able to Edit/Remove posts and comments.

We encourage you to add any other additional features that you find relevant.

Type in terminal
```
$ go run cmd/main.go
```
or
```
$ make server
```

## Authors

[Certina01](https://01.alem.school/git/Certina01) && [AidanaErbolatova](https://01.alem.school/git/AidanaErbolatova)