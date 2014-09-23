package texter

import (
	"testing"
)

var htmlStr = `
<div class="success-indicator hidden off-screen">
    <h1>Test content<small>is this ok?</small><br /><div>with div inside</div></h1>
    <div>This is<small>test</small></div>
    <input data-ten="ten" />
    <div>After setting up the new server, switch my A records to the new ones and hope that the long DNS propagation delay will not have half my userbase writing on one db and the other on another.</div>
    <p>Unless you’re using something like Cloudflare or Route53 that’s obviously a VERY bad idea,
    if you have any kind of writes on your database and you care about your data integrity.</p>
    <script>
      var _gaq = _gaq || [];
          _gaq.push(['_setAccount', 'UA-4210311-3']);
          _gaq.push(['_trackPageview']);
      (function () {
          var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
          ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
          var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();
    </script>
</div>
`

func TestTextFromHTML(t *testing.T) {
	b := []byte(htmlStr)
	txt := textFromHTML(&b)

	if txt != nil {
		t.Errorf("HTML text extraction failed")
	} else {
		t.Skip("Text properly extracted from HTML string")
	}
}
