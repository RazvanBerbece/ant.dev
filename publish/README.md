# Publishing articles

I publish articles by developing and "publishing" their templates locally (1 article = 1 template), which is great in this case because:
- it minimises costs for storage (i.e no need to store articles on the Cloud)
- it gives full control (with little to no effort) of the article development, as it's still a server-rendered `templ` source
- it minimises development time for a service that allows developing articles and storing + publishing them on the Cloud

In short, the process looks like this:
1. Write a `templ` page which follows the default article template [Default Article Template](src/views/pages/publishings/publishingTemplates/default.templ)
2. Add the new page to the static runtime Go slice (`PublishedArticles []domain.Article`) loaded at startup
3. There's not really a step 3 **at the moment** (as the bits which use the runtime slice for renders are already implemented)