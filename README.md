# ant.dev

My approach to a quirky old-school tech blog. I'll be using this to host my articles on in the future :')

Published at: https://antdev-webapp-elba5bz5xq-nw.a.run.app/ (TODO: Get a proper domain)

Note: Public unauthenticated traffic is disabled on the URL above in order to minimise costs and risks while the blog is still under development.

# Buzzwords

* Go
* Templ
* The usual frontend stack: HTML, TailwindCSS, JS (+ a sprinkle of HTMX!)

# Useful Commands

- Generate Go code from `.templ`s: `~/go/bin/templ generate`
- Run the webapp: `docker compose up -d --remove-orphans --build`
- Bring down the webapp: `docker compose down -v`