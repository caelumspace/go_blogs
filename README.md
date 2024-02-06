# go_blogs-api
This is a basic blog api designed using golang.

# Getting started
1. Install Golang on your machine.
  1.1. Make sure you have GOPATH set in your environment variables.
  1.2. Ensure it using `echo %GOPATH%`
2. Get this project by this command: `go get -u https://github.com/caelumspace/go_blogs.git`
3. This will take some time because it downloads this project and downloads all the imported dependencies.
4. Now, `cd go_blogs`
5. Now, run a mogodb server on your local machine, which by default runs on port :27017.
6. Run `go build` to build the go project in a executable file.
7. Run the executable by just typing `go_blogs-api`
8. If the above command prompt Server running @5000, then you are good to go.

# Usage

Testing can be done using POSTMAN

## Endpoints
### Blogs
1. Create a blog > POST "/api/blogs/create" > Enter three key values in Body (x-www-form-urlencoded) `title, content, author, timestamp`
2. Get all blogs > GET "/api/blogs" > This get an array of all blogs.
3. Get a specific blog > GET "/api/blog/{_id}" > This gets the object of the specified blog id.
4. Delete a specific blog > DELETE "/api/blog/{_id}" > This deletes the blog identified by the specified id. 
5. Update a specific blog > PUT "/api/blog/{_id}" > This updates the specified blog. This also requires three key values in Body (x-www-form-urlencoded) `title, content, author, timestamp`

### Comments
1. Create a blog > POST "/api/comments/create" > Enter three key values in Body (x-www-form-urlencoded) `title, content, author, timestamp, blog_id, replyto`
2. Get all blogs > GET "/api/comments" > This get an array of all comments.
3. Get a specific blog > GET "/api/comment/{_id}" > This gets the object of the specified comment id.
4. Delete a specific blog > DELETE "/api/comment/{blog_id}" > This deletes the comment identified by the specified blog id. 
5. Update a specific blog > PUT "/api/comment/{_id}" > This updates the specified comment. This also requires key values in Body (x-www-form-urlencoded) `title, content, author, timestamp, blog_id, replyto`
