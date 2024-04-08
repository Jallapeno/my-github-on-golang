# My Github On Golang

## 1.0 version

<h3>With this project you can print the avatar of your github followers on your terminal</h3>

## running

1 - Login to your GitHub account.

2 - Click on your profile avatar and go to "Settings"

3 - In left sidebar menu, scroll to bottom and select "Developer settings" and "Personal access tokens"

4 - Click on "Generate new token"

5 - Give it a name, select "read:follow" permission to read your followers

6 - Copy the new token generated and change "YOUR_GITHUB_TOKEN" in ``` .env ``` file.

7 - Rename ``` .env-example ``` to ``` .env ```.

8 - Run ``` $ go run main.go mygitlis ``` or ``` $ go build ``` and run ``` $ ./my-github-on-golang  mygitlist ```
