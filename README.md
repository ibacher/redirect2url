redirect2url
============

This is really, really, really not a production-worthy project. It is
just intended as a quick stop-gap for a project and a quick experiment
with a raw Go build on a minimal Docker image.

But, hey, if you want to use it... this simply binds to a port and
redirects users to the URL specified in the `redirect-to` parameter.
If the `redirect-to` parameter isn't supplied, this will just return
a 404.
