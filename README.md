# Chirpy

<!-- ABOUT THE PROJECT -->
## 📌 About this Project

The goal of this project was to learn how the back-end track on Boot.dev comes together as building HTTP servers is the bread-and-butter of a backend developer's day-to-day work.

Goals of this project:
* Understand what web servers are and how they power real-world web applications
* Build a production-style HTTP server in Go, without the use of a framework
* Use JSON, headers, and status codes to communicate with clients via a RESTful API
* Learn what makes Go a great language for building fast web servers
* Use type safe SQL to store and retrieve data from a Postgres database
* Implement a secure authentication/authorization system with well-tested cryptography libraries
* Build and understand webhooks and API keys
* Document the REST API with markdown

### 🛠️ Built With

</h3>
<p align="left">
 <a href="https://go.dev/"><img src="https://skillicons.dev/icons?i=go"></a>&nbsp;
<img src="https://skillicons.dev/icons?i=html">&nbsp;
 </p>

### ✅ Prerequisites

* Install using your terminal to the latest version 
  ```sh
  curl -sS https://webi.sh/golang | sh
  ```

* Verify installation worked
  ```sh
  go version
  ```

## 📦 Installation

1. Clone the repo
   ```sh
   git clone https://github.com/justinmnge/Chirpy.git
   ```
2. Download all dependencies
   ```sh
   go mod download
   ```
3. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ```

## ⚙️ Usage

* Run the server
  ```sh
  go run .
  ```

<!-- LICENSE -->
## 🔐 License

Distributed under the Unlicense License. See `LICENSE.txt` for more information.

<!-- CONTACT -->
## ℹ️ Contact

Justin Monge - hello@justin-monge.dev

Project Link: [https://github.com/justinmnge/Chirpy](https://github.com/justinmnge/Chirpy)
<p align="right">(<a href="#readme-top">back to top</a>)</p>