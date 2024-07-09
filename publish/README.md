# Publishing articles

I release articles by developing and "publishing" their templates locally (1 article = 1 template), which is great in this case because:
- it minimises costs for storage (i.e no need to store articles on the Cloud)
- it gives full control (with little to no effort) over article development, as they're all still server-rendered `templ` sources
- it completely removes the need for a service that allows developing articles and storing + publishing them on the Cloud

In short, the release process looks like this:
1. Write a `templ` page which follows the default article template [Default Article Template](articles/base/default.templ)
2. Add the new page to the static runtime Go slice (`PublishedArticles []domain.Article`) loaded at startup
3. There's not really a step 3 **at the moment** (as the bits which leverage the runtime slice for renders and processing are already implemented)