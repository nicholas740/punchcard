# Punchcard

A fun and lean service with Go and GitHub Actions to fill your GitHub activity. This will essentially mark a green sign on your GitHub contributions for today by committing the date of when the script is run automatically.

## Requirements

You only need these two softwares to run this program:

- [Go 1.19+](https://go.dev/)
- [Git](https://git-scm.com/) and [GitHub](https://github.com/)

## How to Use

There are two ways to run this repository: through GitHub Actions, and/or manually. It is recommended that you do this with GitHub Actions.

### GitHub Actions

- Call the GitHub Action by navigating to the `Actions`, selecting the relevant Action, and then running it manually.

### Manually

- Clone this repository, switch to it, run it, and commit!

```bash
git clone git@github.com:nicholas740/punchcard.git
cd punchcard
go run .
git add .
git commit -m "<YOUR_MESSAGE_HERE>"
git push -u origin main
```

## License

MIT License. Feel free to use it as you see fit and please do not abuse the GitHub Action!
