# ant.dev

My approach to a quirky old-school tech blog. I'll be using this to host my articles on in the future :')

Published at: https://antdev-webapp-elba5bz5xq-nw.a.run.app/ (TODO: Get a proper domain)

*Note: Public unauthenticated traffic is disabled on the URL above in order to minimise costs and security risks while the blog is still under development.*

# Buzzwords

* Go
* Templ
* The usual frontend stack: HTML, TailwindCSS, JS (+ a sprinkle of HTMX!)

# How to Run

1. Generate Go code from `.templ` files: `~/go/bin/templ generate`
2. (Optional) Build the available static styles: `npx tailwindcss build -i static/css/style.css -o static/css/tailwind.css -m`
3. Run the containerised webapp: `docker compose up -d --remove-orphans --build`
4. Bring down the containerised webapp: `docker compose down -v`